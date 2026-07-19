package ctx

import (
	"context"

	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/model/gateway"
	"go.uber.org/zap"
)

type Context struct {
	Event          *gateway.RawEvent
	ContextTimeout context.Context
	Cancel         context.CancelFunc
	Token          *string
	*zap.Logger
	*api.DiscordClient
}
