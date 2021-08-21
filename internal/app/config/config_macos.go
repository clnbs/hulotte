// +build macos

package config

import (
	"os"
)

func ConfigFilePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return configDir + "/hulotte/config.json", nil
}

func LogoPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return configDir + "/hulotte/logo.png", nil
}

func SoundPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return configDir + "/hulotte/sound.mp3", nil
}

func ConfigDirPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return configDir + "/hulotte", nil
}
