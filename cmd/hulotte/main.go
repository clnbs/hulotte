package main

import (
	"sync"
	"time"

	"github.com/clnbs/hulotte/internal/app/config"
	"github.com/clnbs/hulotte/internal/app/menu"
	"github.com/clnbs/hulotte/internal/app/notify"
	"github.com/clnbs/hulotte/internal/app/sound"
)

var notifyer notify.Notifyer
var soundPlayer sound.SoundPlayer

func init() {
	var err error

	configPath, err := config.GetConfigFilePath()
	if err != nil {
		panic(err)
	}

	notifyer = notify.NewBeeeeep()
	err = notifyer.Initialize(configPath)
	if err != nil {
		panic(err)
	}

	soundPlayer = sound.NewAudio()
	err = soundPlayer.Initialize(configPath)
	if err != nil {
		panic(err)
	}
}

func main() {
	configPath, err := config.GetConfigFilePath()
	if err != nil {
		panic(err)
	}
	m := menu.SystrayMenu{}
	err = m.Initialize(configPath)
	if err != nil {
		panic(err)
	}

	m.SetDeamons(eyesHandler)
	m.Start()
}

func eyesHandler() {
	for {
		time.Sleep(10 * time.Second)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			err := notifyer.Trigger()
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			err := soundPlayer.Trigger()
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()
		wg.Wait()
	}
}
