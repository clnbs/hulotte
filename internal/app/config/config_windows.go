// +build windows

package config

func GetConfigFilePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return configDir + "\\hulotte\\config.json", nil
}

func GetConfigDirPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return configDir + "\\hulotte", nil
	return "", nil
}
