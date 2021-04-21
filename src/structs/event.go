package structs

type Control string

const (
	Quit Control = "Quit"
	Keyboard = "Keyboard"
	Reload = "Reload"
)

type Event struct {
	Source string
	Type Control
	Value string
}
