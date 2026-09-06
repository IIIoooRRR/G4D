package g4d

import (
	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/gateway"
	"github.com/IIIoooRRR/G4D/model/_const"
	"go.uber.org/zap"
)

func NewBot(token string, handler PanicHandler, gateway *gateway.Receiver, logger *zap.Logger, client *api.DiscordClient) *Bot {
	return &Bot{
		Token:        token,
		Logger:       logger,
		Gateway:      gateway,
		PanicHandler: handler,
		Client:       client,
	}
}
func (b *Bot) SetPrefix(pref string) *Bot {
	b.Prefix = pref
	return b
}

func WithDispQuantity(qnt int) _const.Quantity {
	return _const.Quantity(qnt)
}
func WithSemaphoreLimit(limit int) _const.SemaphoreLimit {
	return _const.SemaphoreLimit(limit)
}
