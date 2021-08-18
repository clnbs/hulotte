package menu

import (
	"bytes"
	"encoding/json"
	"image"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/clnbs/hulotte/internal/app/config"
	"github.com/getlantern/systray"
)

type SystrayMenu struct {
	logo     []byte
	tooltip  string
	handlers []func()
}

type SystrayMenuConfig struct {
	MenuLogo    string `json:"menu_logo"`
	MenuTooltip string `json:"menu_tooltip"`
}

type menuItemNaming struct {
	title   string
	tooltip string
}

const (
	quitItem = "menu"
)

var internationalMenuNaming = map[string]map[string]menuItemNaming{
	config.FrenchISO639: {
		quitItem: {
			title:   "Quitter",
			tooltip: "Quitte l'application",
		},
	},
	config.EnglishIOS639: {
		quitItem: {
			title:   "Quit",
			tooltip: "Quit the application",
		},
	},
}

func (sm *SystrayMenu) Initialize(path string) error {
	config, err := loadConfig(path)
	if err != nil {
		return err
	}

	logoFile, err := os.Open(config.MenuLogo)
	if err != nil {
		return err
	}
	defer logoFile.Close()

	imData, _, err := image.Decode(logoFile)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = png.Encode(buf, imData)
	if err != nil {
		return err
	}

	sm.logo = buf.Bytes()
	sm.tooltip = config.MenuTooltip
	return nil
}

func (sm *SystrayMenu) Start() {
	systray.Run(sm.onReady, sm.onExit)
}

func (sm *SystrayMenu) SetDeamons(handlers ...func()) {
	sm.handlers = handlers
}

func loadConfig(path string) (*SystrayMenuConfig, error) {
	configData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &SystrayMenuConfig{}
	err = json.Unmarshal(configData, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (sm *SystrayMenu) onExit() {}

func (sm *SystrayMenu) onReady() {
	systray.SetIcon(sm.logo)
	systray.SetTitle("")
	systray.SetTooltip(sm.tooltip)

	localeMenuItem := internationalMenuNaming[config.Locale()]

	menuQuit := systray.AddMenuItem(localeMenuItem[quitItem].title, localeMenuItem[quitItem].tooltip)
	go func() {
		<-menuQuit.ClickedCh
		systray.Quit()
	}()

	for _, h := range sm.handlers {
		go h()
	}

	menuQuit.SetIcon(sm.logo)
}
