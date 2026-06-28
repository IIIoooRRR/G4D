package g4d

import "github.com/IIIoooRRR/G4D/model/gateway"

type Definition func(*Bot, uint)

type ToCommand func(*gateway.RawEvent) error //объявляет любую функцию Execute, которая подъодит по условию
func (f ToCommand) Execute(event *gateway.RawEvent) error {
	return f(event)
}

var (
	StaticEventProcessor  Definition = staticEventProcessor
	DynamicEventProcessor Definition = dynamicEventProcessor
)
