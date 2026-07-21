package g4d

import (
	"github.com/IIIoooRRR/G4D/model/ctx"
)

func (b *Bot) InitProcessors(p Processor, quantity uint, limitSize uint) {
	for i := 0; i < int(quantity); i++ {
		go p(b, limitSize)
	}
}
func (b *Bot) newCtx() ctx.Context {
	return ctx.Context{
		Prefix:        b.Prefix,
		DiscordClient: b.Client,
		Logger:        b.cmdLogger,
	}
}
