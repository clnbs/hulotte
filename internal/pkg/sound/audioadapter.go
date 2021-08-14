package sound

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/bobertlo/go-mpg123/mpg123"
	"github.com/gordonklaus/portaudio"
)

type Audio struct {
	decoder   *mpg123.Decoder
	soundPath string
	rate      int64
	channels  int
}

type AudioConfig struct {
	SoundPath string `json:"audio_sound"`
}

func NewAudio() *Audio {
	return &Audio{}
}

func (a *Audio) Initialize() error {
	config, err := loadConfig()
	if err != nil {
		return err
	}
	a.soundPath = config.SoundPath

	a.decoder, err = mpg123.NewDecoder("")
	if err != nil {
		return err
	}

	return nil
}

func (a *Audio) Trigger() error {
	err := a.decoder.Open(a.soundPath)
	if err != nil {
		return err
	}

	// get audio format information
	a.rate, a.channels, _ = a.decoder.GetFormat()

	// make sure output format does not change
	a.decoder.FormatNone()
	a.decoder.Format(a.rate, a.channels, mpg123.ENC_SIGNED_16)

	portaudio.Initialize()
	defer portaudio.Terminate()
	out := make([]int16, 8192)

	stream, err := portaudio.OpenDefaultStream(0, a.channels, float64(a.rate), len(out), &out)
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
		_, err = a.decoder.Read(audio)
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

func loadConfig() (*AudioConfig, error) {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		return nil, errors.New("HOME env variable is empty, not supported")
	}
	configFilePath := homeDir + "/.hulotte/config.json"
	configData, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	config := &AudioConfig{}
	err = json.Unmarshal(configData, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
