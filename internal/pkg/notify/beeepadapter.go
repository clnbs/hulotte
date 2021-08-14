// +build linux

package notify

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/gen2brain/beeep"
)

type Beeeeep struct {
	title   string
	message string
	logo    string
}

type BeeeeepConfig struct {
	NotifyTitle   string `json:"notify_title"`
	NotifyMessage string `json:"notify_message"`
	NotifyLogo    string `json:"notify_logo"`
}

func NewBeeeeep() *Beeeeep {
	return &Beeeeep{}
}

func (b *Beeeeep) Initialize() error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	b.logo = config.NotifyLogo
	b.message = config.NotifyMessage
	b.title = config.NotifyTitle
	return nil
}

func (b *Beeeeep) Trigger() error {
	// logo path : "/home/colin/go/src/github.com/clnbs/hulotte/assets/information.png"
	err := beeep.Notify(b.title, b.message, b.logo)
	if err != nil {
		return err
	}
	return nil
}

func loadConfig() (*BeeeeepConfig, error) {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		return nil, errors.New("HOME env variable is empty, not supported")
	}
	configFilePath := homeDir + "/.hulotte/config.json"
	configData, err := ioutil.ReadFile(configFilePath)
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