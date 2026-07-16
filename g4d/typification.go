package g4d

import "github.com/IIIoooRRR/G4D/model/gateway"

type Processor func(*Bot, uint)

type ToCommand func(*gateway.RawEvent) error //объявляет любую функцию Execute, которая подъодит по условию

var (
	StaticEventProcessor  Processor = staticEventProcessor
	DynamicEventProcessor Processor = dynamicEventProcessor
)
