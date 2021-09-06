package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/clnbs/hulotte/internal/app/helper"
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
	configPath, err := ConfigFilePath()
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
	configPath, _ := ConfigFilePath()
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return
	}
	config := ConfigFile{}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return
	}
	config.AudioMute = set
	configData, err = json.Marshal(config)
	if err != nil {
		return
	}
	_ = helper.WriteFile(configData, configPath)
}

func DoesConfigExists() (bool, error) {
	configDirPath, err := ConfigDirPath()
	if err != nil {
		return false, err
	}
	return helper.DoesFileExsist(configDirPath)
}

func CreateConfig() error {
	configDirPath, err := ConfigDirPath()
	if err != nil {
		return err
	}
	fmt.Println("Creating config dir in", configDirPath)
	err = helper.CreateDir(configDirPath)
	if err != nil {
		return err
	}

	fmt.Println("Generating config file ...")
	configFileData, err := generateDefaultConfigFile()
	if err != nil {
		return err
	}
	fmt.Println("Done!")
	configFilePath, err := ConfigFilePath()
	if err != nil {
		return err
	}
	fmt.Println("Writing config file at", configFilePath)
	err = helper.WriteFile(configFileData, configFilePath)
	if err != nil {
		return err
	}
	fmt.Println("Done!")
	return nil
}

func generateDefaultConfigFile() ([]byte, error) {
	var err error
	interConfig := defaultConfigFile[locale]
	interConfig.AudioSound, err = SoundPath()
	if err != nil {
		return nil, err
	}
	interConfig.MenuLogo, err = LogoPath()
	if err != nil {
		return nil, err
	}
	interConfig.NotifyLogo = interConfig.MenuLogo
	configData, err := json.Marshal(&interConfig)
	if err != nil {
		return nil, err
	}
	return configData, nil
}

func loadAudioMuteConfig(path string) (bool, error) {
	configExists, err := DoesConfigExists()
	if err != nil {
		return false, err
	}
	if !configExists {
		return true, nil
	}
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
