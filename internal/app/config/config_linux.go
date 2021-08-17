// +build linux

package config

import (
	"errors"
	"os"
)

func GetConfigFilePath() (string, error) {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		return "", errors.New("unsupported linux distribution")
	}
	return homeDir + "/.hulotte/config.json", nil
}

func GetConfigDirPath() (string, error) {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		return "", errors.New("unsupported linux distribution")
	}
	return homeDir + "/.hulotte/", nil
}

func WriteFile(data []byte, path string) error {
	// TODO thumbs
	return nil
}

func CreateDir(path string) error {
	// TODO thumbs
	return nil
}
