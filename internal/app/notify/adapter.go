package notify

type MessageTag string
type Notifyer interface {
	Initialize(path string) error
	Trigger(message MessageTag) error
}
