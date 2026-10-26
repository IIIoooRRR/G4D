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

func (ctx *Context) WithPrefix(s1, s2 string) bool {
	return s1 == ctx.Prefix+s2
}
