package notify

type Notifyer interface {
	Initialize(path string) error
	Trigger() error
}
