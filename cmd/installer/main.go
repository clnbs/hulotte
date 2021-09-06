package main

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io/fs"

	"github.com/clnbs/hulotte/internal/app/config"
	"github.com/clnbs/hulotte/internal/app/helper"
	"github.com/clnbs/hulotte/internal/app/installer"
)

//go:embed content/hulotte.zip
var hulotteContent []byte

func main() {
	fmt.Println("extracting zip ...")
	hulotteData, soundData, logoData, err := extractHulotteZip()
	if err != nil {
		panic(err)
	}

	fmt.Println("Done! \nChecking if Hulotte is already installed ..")

	isHulotteInstalled, err := installer.DoesHulotteExists()
	if err != nil {
		panic(err)
	}
	if !isHulotteInstalled {
		fmt.Println("Hulotte is not installed! Install ongoing ...")
		err = installHulotte(hulotteData)
		if err != nil {
			panic(err)
		}

		fmt.Println("Done!")
	} else {
		fmt.Println("Hulotte is already installed.")
	}

	fmt.Println("Checking if config exists ...")
	isConfigExists, err := config.DoesConfigExists()
	if err != nil {
		panic(err)
	}
	if !isConfigExists {
		fmt.Println("Config does not exists! Install ongoing ...")
		err = config.CreateConfig()
		if err != nil {
			panic(err)
		}
		fmt.Println("Done!\nInstalling assets ...")
		err = installAssets(soundData, logoData)
		if err != nil {
			panic(err)
		}
		fmt.Println("Done!")
	} else {
		fmt.Println("Config is already installed, leaving ...")
	}
}

func installHulotte(hulotte []byte) error {
	// TODO get the binary
	fmt.Println("Writing Hulotte binary ...")
	err := installer.WriteHulotte(hulotte)
	if err != nil {
		return err
	}
	fmt.Println("Done!\nRegistering Hulotte ...")
	err = installer.RegisterHulotte()
	if err != nil {
		return err
	}
	fmt.Println("Done!\nSetting up autostart ...")
	err = installer.SetAutoStart()
	if err != nil {
		return err
	}
	fmt.Println("Done!")
	return nil
}

func installAssets(sound []byte, logo []byte) error {
	logoPath, err := config.LogoPath()
	if err != nil {
		return err
	}
	fmt.Println("Writing logo at", logoPath)
	err = helper.WriteFile(logo, logoPath)
	if err != nil {
		return err
	}

	soundPath, err := config.SoundPath()
	if err != nil {
		return err
	}
	fmt.Println("Writting sound file at", soundPath)
	return helper.WriteFile(sound, soundPath)
}

func extractHulotteZip() ([]byte, []byte, []byte, error) {
	var hulotte, sound, logo []byte
	reader := bytes.NewReader(hulotteContent)
	zipReader, err := zip.NewReader(reader, reader.Size())
	if err != nil {
		return nil, nil, nil, err
	}

	hulotte, err = fs.ReadFile(zipReader, "assets/hulotte")
	if err != nil {
		return nil, nil, nil, err
	}

	sound, err = fs.ReadFile(zipReader, "assets/sound.mp3")
	if err != nil {
		return nil, nil, nil, err
	}

	logo, err = fs.ReadFile(zipReader, "assets/logo.png")
	if err != nil {
		return nil, nil, nil, err
	}

	if len(hulotte) == 0 {
		return nil, nil, nil, errors.New("Hulotte data is empty")
	}

	if len(sound) == 0 {
		return nil, nil, nil, errors.New("sound data is empty")
	}
	if len(logo) == 0 {
		return nil, nil, nil, errors.New("logo data is empty")
	}
	return hulotte, sound, logo, nil
}
