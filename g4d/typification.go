package g4d

import (
	"github.com/IIIoooRRR/G4D/model/ctx"
	"github.com/IIIoooRRR/G4D/model/gateway"
)

type Processor func(*Bot, uint)

type Execute func(event *gateway.RawEvent, ctx *ctx.Context) error //объявляет любую функцию Execute, которая подъодит по условию

var (
	StaticEventProcessor  Processor = staticEventProcessor
	DynamicEventProcessor Processor = dynamicEventProcessor
)
