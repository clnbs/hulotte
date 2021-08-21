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

	configPath, err := config.ConfigFilePath()
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
	configPath, err := config.ConfigFilePath()
	if err != nil {
		panic(err)
	}
	m := menu.SystrayMenu{}
	err = m.Initialize(configPath)
	if err != nil {
		panic(err)
	}

	m.SetDeamons(eyesHandler, walkHandler)
	m.Start()
}

func eyesHandler() {
	for {
		time.Sleep(10 * time.Second)
		wg := sync.WaitGroup{}
		go func() {
			wg.Add(1)
			defer wg.Done()
			err := notifyer.Trigger(notify.EyesMessage)
			if err != nil {
				panic(err)
			}
		}()

		go func() {
			wg.Add(1)
			defer wg.Done()
			if config.Mute() {
				return
			}
			err := soundPlayer.Trigger()
			if err != nil {
				panic(err)
			}
		}()
		wg.Wait()
	}
}

func walkHandler() {
	for {
		time.Sleep(60 * time.Minute)
		wg := sync.WaitGroup{}
		go func() {
			wg.Add(1)
			defer wg.Done()
			err := notifyer.Trigger(notify.WalkMessage)
			if err != nil {
				panic(err)
			}
		}()

		go func() {
			wg.Add(1)
			defer wg.Done()
			err := soundPlayer.Trigger()
			if err != nil {
				panic(err)
			}
		}()
		wg.Wait()
	}
}
