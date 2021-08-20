package config

import (
	"encoding/json"
	"os"

	"github.com/cloudfoundry/jibber_jabber"
)

type ConfigFile struct {
	NotifyTitle string `json:"notify_title"`
	NotifyLogo  string `json:"notify_logo"`
	AudioSound  string `json:"audio_sound"`
	AudioMute   bool   `json:"audio_mute"`
	MenuLogo    string `json:"menu_logo"`
	MenuTooltip string `json:"menu_tooltip"`
}

const (
	FrenchISO639  = "fr"
	EnglishIOS639 = "en"
)

var locale string
var audioMute bool

var defaultConfigFile = map[string]ConfigFile{
	FrenchISO639: {
		NotifyTitle: "Hulotte",
		NotifyLogo:  "logo.png",
		AudioSound:  "sound.mp3",
		MenuLogo:    "logo.png",
		MenuTooltip: "Hulotte - L'ami des yeux",
	},
	EnglishIOS639: {
		NotifyTitle: "Hulotte",
		NotifyLogo:  "logo.png",
		AudioSound:  "sound.mp3",
		MenuLogo:    "logo.png",
		MenuTooltip: "Hulotte - The friend of the eyes",
	},
}

func init() {
	configPath, err := GetConfigFilePath()
	if err != nil {
		panic(err)
	}
	audioMute, err = loadAudioMuteConfig(configPath)
	if err != nil {
		panic(err)
	}

	userLocale, err := jibber_jabber.DetectLanguage()
	if err != nil {
		panic(err)
	}
	if userLocale != "en" && userLocale != "fr" {
		locale = "en"
		return
	}
	locale = userLocale
}

func Locale() string {
	return locale
}

func Mute() bool {
	return audioMute
}

func SetMute(set bool) {
	audioMute = set
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

func loadAudioMuteConfig(path string) (bool, error) {
	configData, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}

	config := &ConfigFile{}
	err = json.Unmarshal(configData, config)
	if err != nil {
		return false, err
	}

	return config.AudioMute, nil
}
