package main

import (
	"sync"
	"time"

	"github.com/clnbs/hulotte/internal/pkg/menu"
	"github.com/clnbs/hulotte/internal/pkg/notify"
	"github.com/clnbs/hulotte/internal/pkg/sound"
)

var notifyer notify.Notifyer
var soundPlayer sound.SoundPlayer

func init() {
	var err error
	notifyer = notify.NewBeeeeep()
	err = notifyer.Initialize()
	if err != nil {
		panic(err)
	}

	soundPlayer = sound.NewAudio()
	err = soundPlayer.Initialize()
	if err != nil {
		panic(err)
	}
}

func main() {
	m := menu.SystrayMenu{}
	err := m.Initialize()
	if err != nil {
		panic(err)
	}

	m.SetDeamons(handler)
	m.Start()
}

func handler() {
	for {
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
		time.Sleep(10 * time.Second)
	}
}
