package G4D

import (
	"context"
	"sync"

	"github.com/IIIoooRRR/G4D/Connect"
)

type Bot struct {
	Token         string
	Gateway       *Connect.Receiver
	Prefix        string
	CommandBuffer []Command
	appId         string
	Context       context.Context
	Cache         Connect.Cacher
}

var bot *Bot
var botMu sync.RWMutex

func CurrentBot() *Bot {
	botMu.RLock()
	defer botMu.RUnlock()
	return bot
}
func (b *Bot) Run() error {
	b.Gateway.Cache = b.Cache
	botMu.Lock()
	bot = b
	botMu.Unlock()
	err := b.Gateway.CreateBot(b.Context, &b.Token)
	if err != nil {
		return err
	}
	return nil
}
