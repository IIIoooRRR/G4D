package g4d

import (
	"github.com/IIIoooRRR/G4D/model/gateway"
)

type CommandTemplate struct {
	Trigger string
	Name    string
	Action  ToCommand
}

type SlashCommandTemplate struct {
	Form            SlashCreateCommand
	CommandTemplate CommandTemplate
}
type ToCommand func(*gateway.RawEvent) error //объявляет любую функцию Execute, которая подъодит по условию
func (f ToCommand) Execute(event *gateway.RawEvent) error {
	return f(event)
}
