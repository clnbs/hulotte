package sound

type SoundPlayer interface {
	Initialize(path string) error
	Trigger() error
}
