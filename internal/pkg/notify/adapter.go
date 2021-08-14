package notify

type Notifyer interface {
	Initialize() error
	Trigger() error
}
