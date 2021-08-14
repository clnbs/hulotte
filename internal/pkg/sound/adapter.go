package sound

type SoundPlayer interface {
	Initialize() error
	Trigger() error
}
