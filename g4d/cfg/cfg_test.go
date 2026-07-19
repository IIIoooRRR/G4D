package cfg_test

import (
	"testing"

	"github.com/IIIoooRRR/G4D/g4d"
	"github.com/IIIoooRRR/G4D/g4d/cfg"
	"github.com/IIIoooRRR/G4D/model/gateway"
	"go.uber.org/zap"
)

func TestLoadCfg(t *testing.T) {
	bot := cfg.LoadBot("config.exp.yaml", zap.Must(zap.NewProduction()), panicHandler{})
	t.Log(bot)
}

type panicHandler struct {
	g4d.PanicHandler
	Logger *zap.Logger
}

func (p panicHandler) OnPanic(event *gateway.RawEvent, cmd *g4d.CommandTemplate, r any, stack []byte) {
	p.Logger.Panic("cmd:", zap.Any("cmd", cmd), zap.String("stack", string(stack)), zap.Any("r:", r), zap.Any("event", event))
}
