package sound

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
)

type Audio struct {
	decoder   *minimp3.Decoder
	soundPath string
	data      []byte
	ctx       *oto.Context
}

type AudioConfig struct {
	SoundPath string `json:"audio_sound"`
}

func NewAudio() *Audio {
	return &Audio{}
}

func (a *Audio) Initialize(path string) error {
	config, err := loadConfig(path)
	if err != nil {
		return err
	}
	a.soundPath = config.SoundPath

	a.decoder = &minimp3.Decoder{}

	var file []byte
	if file, err = ioutil.ReadFile(a.soundPath); err != nil {
		return err
	}
	a.decoder, a.data, err = minimp3.DecodeFull(file)
	if err != nil {
		return err
	}
	if a.ctx, err = oto.NewContext(a.decoder.SampleRate, a.decoder.Channels, 2, 1024); err != nil {
		return err
	}

	return nil
}

func (a *Audio) Trigger() error {
	var player = a.ctx.NewPlayer()
	player.Write(a.data)

	<-time.After(time.Second)

	a.decoder.Close()
	if err := player.Close(); err != nil {
		return err
	}

	return nil
}

func loadConfig(path string) (*AudioConfig, error) {
	configData, err := ioutil.ReadFile(path)
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
