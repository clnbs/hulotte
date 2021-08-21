package main

import (
	"archive/zip"
	"bytes"

	"github.com/clnbs/hulotte/internal/app/config"
	"github.com/clnbs/hulotte/internal/app/helper"
	"github.com/clnbs/hulotte/internal/app/installer"
)

//go:embed content/hulotte.zip
var hulotteContent []byte

func main() {
	hulotteData, soundData, logoData, err := extractHulloteZip()
	if err != nil {
		panic(err)
	}

	isHulotteInstalled, err := installer.DoesHulotteExists()
	if err != nil {
		panic(err)
	}
	if !isHulotteInstalled {

		if err != nil {
			panic(err)
		}

		err = installHulotte(hulotteData)
		if err != nil {
			panic(err)
		}
	}

	isConfigExists, err := config.DoesConfigExists()
	if err != nil {
		panic(err)
	}
	if !isConfigExists {
		err = config.CreateConfig()
		if err != nil {
			panic(err)
		}
		err = installAssets(soundData, logoData)
		if err != nil {
			panic(err)
		}
	}
}

func installHulotte(hulotte []byte) error {
	// TODO get the binary
	err := installer.WriteHulotte(hulotte)
	if err != nil {
		return err
	}
	err = installer.RegisterHulotte()
	if err != nil {
		return err
	}
	err = installer.SetAutoStart()
	if err != nil {
		return err
	}
	return nil
}

func installAssets(sound []byte, logo []byte) error {
	logoPath, err := config.LogoPath()
	if err != nil {
		return err
	}
	err = helper.WriteFile(logo, logoPath)
	if err != nil {
		return err
	}

	soundPath, err := config.SoundPath()
	if err != nil {
		return err
	}
	return helper.WriteFile(sound, soundPath)
}

func extractHulloteZip() (hulotte []byte, sound []byte, logo []byte, err error) {
	reader := bytes.NewReader(hulotteContent)
	zipReader, err := zip.NewReader(reader, reader.Size())
	if err != nil {
		return
	}
	hulloteFile, err := zipReader.Open("hullote")
	if err != nil {
		return
	}

	soundFile, err := zipReader.Open("sound.mp3")
	if err != nil {
		return
	}

	logoFile, err := zipReader.Open("logo.png")
	if err != nil {
		return
	}

	_, err = hulloteFile.Read(hulotte)
	if err != nil {
		return
	}

	_, err = soundFile.Read(sound)
	if err != nil {
		return
	}

	_, err = logoFile.Read(logo)
	return
}
