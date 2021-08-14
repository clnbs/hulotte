package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/png"
	"os"
	"sync"
	"time"

	"github.com/bobertlo/go-mpg123/mpg123"
	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/gordonklaus/portaudio"
)

func main() {
	systray.Run(onReady, onExit)

}

// TODO naming
func timeToLookAway() {
	for {
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			err := playNotification()
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			err := playSound()
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()
		wg.Wait()
		time.Sleep(10 * time.Second)
	}
}

func playNotification() error {
	err := beeep.Notify("Notify", "Notify message body", "/home/colin/go/src/github.com/clnbs/hulotte/assets/information.png")
	if err != nil {
		return err
	}
	return nil
}

func playSound() error {
	decoder, err := mpg123.NewDecoder("")
	if err != nil {
		return err
	}

	err = decoder.Open("/home/colin/Downloads/alerte.mp3")
	if err != nil {
		return err
	}
	// get audio format information
	rate, channels, _ := decoder.GetFormat()

	// make sure output format does not change
	decoder.FormatNone()
	decoder.Format(rate, channels, mpg123.ENC_SIGNED_16)

	portaudio.Initialize()
	defer portaudio.Terminate()
	out := make([]int16, 8192)
	stream, err := portaudio.OpenDefaultStream(0, channels, float64(rate), len(out), &out)
	if err != nil {
		return err
	}
	defer stream.Close()

	err = stream.Start()
	if err != nil {
		return err
	}

	defer stream.Stop()
	for {
		audio := make([]byte, 2*len(out))
		_, err = decoder.Read(audio)
		if err == mpg123.EOF {
			break
		}
		if err != nil {
			return err
		}

		err = binary.Read(bytes.NewBuffer(audio), binary.LittleEndian, out)
		if err != nil {
			return err
		}
		err = stream.Write()
		if err != nil {
			return err
		}
	}

	return nil
}

func onExit() {
}

func onReady() {
	logoFile, err := os.Open("/home/colin/go/src/github.com/clnbs/hulotte/assets/information.png")
	if err != nil {
		panic(err)
	}
	defer logoFile.Close()

	imData, _, err := image.Decode(logoFile)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	err = png.Encode(buf, imData)
	if err != nil {
		panic(err)
	}

	systray.SetIcon(buf.Bytes())
	systray.SetTitle("")
	systray.SetTooltip("Hulotte - L'ami des yeux")
	mQuit := systray.AddMenuItem("Quitter", "Quitte l'application")
	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	go timeToLookAway()

	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(buf.Bytes())
}
