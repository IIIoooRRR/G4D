package G4D

import (
	"github.com/IIIoooRRR/G4D/Connect"
)

type Command struct {
	Trigger string
	Action  Action
}
type CommandTemplate struct {
	Trigger string
	Action  ToCommand
}
type SlashCommandTemplate struct {
	Form            SlashCreateCommand
	CommandTemplate CommandTemplate
}

type ToCommand func(*Connect.RawEvent) error //объявляет любую функцию Execute, которая подъодит по условию
func (f ToCommand) Execute(event *Connect.RawEvent) error {
	return f(event)
}

type Action interface {
	Execute(event *Connect.RawEvent) error
}
