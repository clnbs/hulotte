package main

import (
	"os"

	"github.com/clnbs/hulotte/internal/app/config"
)

func main() {
	err := createConfig()
	if err != nil {
		panic(err)
	}
}

func checkConfigPresence() (bool, error) {
	configDirPath, err := config.GetConfigDirPath()
	if err != nil {
		return false, err
	}
	_, err = os.Stat(configDirPath)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func createConfig() error {
	configDirPath, err := config.GetConfigDirPath()
	if err != nil {
		return err
	}

	err = config.CreateDir(configDirPath)
	if err != nil {
		return err
	}

	configFileData, err := config.GenerateDefaultConfigFile()
	if err != nil {
		return err
	}
	configFilePath, err := config.GetConfigFilePath()
	if err != nil {
		return err
	}
	err = config.WriteFile(configFileData, configFilePath)
	if err != nil {
		return err
	}
	return nil
}
