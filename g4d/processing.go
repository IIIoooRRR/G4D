package g4d

import (
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/ctx"
)

func (b *Bot) initProcessors(pType _const.ProcessorType, quantity _const.Quantity, limitSize _const.SemaphoreLimit) {
	b.processorsOnce.Do(func() {
		var processor func(int, *chan struct{})
		if pType == _const.DynamicEventProcessor {
			processor = b.dynamicEventProcessor
		} else if pType == _const.StaticEventProcessor {
			processor = b.staticEventProcessor
		} else {
			panic("Unknown processor type")
		}
		channel := make(chan struct{}, limitSize)
		for i := range quantity {
			go processor((int)(i), &channel)
		}
	})
}

func (b *Bot) newCtx() ctx.Context {
	return ctx.Context{
		Prefix:        b.Prefix,
		DiscordClient: b.Client,
		Logger:        b.cmdLogger,
	}
}
