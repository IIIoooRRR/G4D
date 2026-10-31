package g4d

import (
	"errors"
	"fmt"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/ctx"
	"go.uber.org/zap"
)

func (b *Bot) initProcessors(pType _const.ProcessorType, quantity _const.Quantity, limitSize _const.SemaphoreLimit) error {
	processor := b.processorPool[pType]
	b.Logger.Info(string(pType), zap.Any("processor", processor))
	if processor == nil {
		return errors.New("unkown processor type")
	}
	channel := make(chan struct{}, limitSize)
	b.processorsOnce.Do(func() {
		for i := range quantity {
			go processor(b, int(i), channel)
		}
	})
	return nil
}

func (b *Bot) RegisterProcessor(pType _const.ProcessorType, fn func(*Bot, int, chan struct{})) error {
	if _, exists := b.processorPool[pType]; exists {
		return fmt.Errorf("processor %q already registered", pType)
	}
	b.processorPool[pType] = fn
	return nil
}

func (b *Bot) newCtx() ctx.Context {
	return ctx.Context{
		Prefix:        b.Prefix,
		DiscordClient: b.Client,
		Logger:        b.cmdLogger,
	}
}
