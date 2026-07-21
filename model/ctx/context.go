package ctx

import (
	"github.com/IIIoooRRR/G4D/api"
	"go.uber.org/zap"
)

type Context struct {
	Prefix string
	Logger *zap.Logger
	*api.DiscordClient
}
