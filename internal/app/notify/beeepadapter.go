// +build linux

package notify

import (
	"encoding/json"
	"io/ioutil"

	"github.com/clnbs/hulotte/internal/app/config"
	"github.com/gen2brain/beeep"
)

type Beeeeep struct {
	title string
	logo  string
}

type BeeeeepConfig struct {
	NotifyTitle string `json:"notify_title"`
	NotifyLogo  string `json:"notify_logo"`
}

const (
	EyesMessage MessageTag = "eyes"
	WalkMessage MessageTag = "walk"
)

var internationalMessageNaming = map[string]map[MessageTag]string{
	config.FrenchISO639: {
		EyesMessage: "Prenez soin de vos yeux ... Regardez un point à plus de 20 mètres pendant 20 secondes",
		WalkMessage: "Allez prendre une peu l'air ... marcher de temps en temps est important",
	},
	config.EnglishIOS639: {
		EyesMessage: "Take care of your eyes ... Look at a point at least 20 meters far for 20 seconds",
		WalkMessage: "Go get some fresh air... walking from time to time is important",
	},
}

func NewBeeeeep() *Beeeeep {
	return &Beeeeep{}
}

func (b *Beeeeep) Initialize(path string) error {
	config, err := loadConfig(path)
	if err != nil {
		return err
	}

	b.logo = config.NotifyLogo
	b.title = config.NotifyTitle
	return nil
}

func (b *Beeeeep) Trigger(tag MessageTag) error {
	locale := internationalMessageNaming[config.Locale()]
	err := beeep.Notify(b.title, locale[tag], b.logo)
	if err != nil {
		return err
	}
	return nil
}

func loadConfig(path string) (*BeeeeepConfig, error) {
	configData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &BeeeeepConfig{}
	err = json.Unmarshal(configData, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
