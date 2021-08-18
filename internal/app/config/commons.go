package config

import (
	"encoding/json"
	"os"

	"github.com/cloudfoundry/jibber_jabber"
)

type ConfigFile struct {
	NotifyTitle   string `json:"notify_title"`
	NotifyMessage string `json:"notify_message"`
	NotifyLogo    string `json:"notify_logo"`
	AudioSound    string `json:"audio_sound"`
	MenuLogo      string `json:"menu_logo"`
	MenuTooltip   string `json:"menu_tooltip"`
}

const (
	FrenchISO639  = "fr"
	EnglishIOS639 = "en"
)

var locale string

var defaultConfigFile = map[string]ConfigFile{
	FrenchISO639: {
		NotifyTitle:   "Hulotte",
		NotifyMessage: "Regardez un point à plus de 20 mètres pendant 20 secondes",
		NotifyLogo:    "logo.png",
		AudioSound:    "sound.mp3",
		MenuLogo:      "logo.png",
		MenuTooltip:   "Hulotte - L'ami des yeux",
	},
	EnglishIOS639: {
		NotifyTitle:   "Hulotte",
		NotifyMessage: "Look at a point at least 20 meters far for 20 seconds",
		NotifyLogo:    "logo.png",
		AudioSound:    "sound.mp3",
		MenuLogo:      "logo.png",
		MenuTooltip:   "Hulotte - The friend of the eyes",
	},
}

func init() {
	userLocale, err := jibber_jabber.DetectLanguage()
	if err != nil {
		panic(err)
	}
	locale = userLocale
}

func Locale() string {
	return locale
}

func GenerateDefaultConfigFile() ([]byte, error) {
	interConfig := defaultConfigFile[locale]
	configData, err := json.Marshal(&interConfig)
	if err != nil {
		return nil, err
	}
	return configData, nil
}

func WriteFile(data []byte, path string) error {
	err := os.WriteFile(path, data, 0755)
	if err != nil {
		return err
	}
	return nil
}

func CreateDir(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}
	return nil
}
