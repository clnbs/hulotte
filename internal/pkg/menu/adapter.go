package menu

type MenuPrinter interface {
	Initialize() error
	Start()
	SetDeamons(...func())
}
