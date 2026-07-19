package g4d

import (
	"context"
	"time"

	"github.com/IIIoooRRR/G4D/model/ctx"
)

func (b *Bot) InitProcessors(p Processor, quantity uint, limitSize uint) {
	for i := 0; i < int(quantity); i++ {
		go p(b, limitSize)
	}
}
func (b *Bot) newCtx() ctx.Context {
	context, cancel := context.WithTimeout(context.Background(), time.Second*b.CtxTimeout)
	return ctx.Context{
		ContextTimeout: context,
		Cancel:         cancel,
		Token:          &b.Token,
		DiscordClient:  b.Client,
		Logger:         b.cmdLogger,
	}
}
