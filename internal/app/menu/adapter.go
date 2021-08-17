package menu

type MenuPrinter interface {
	Initialize(string) error
	Start()
	SetDeamons(...func())
}
