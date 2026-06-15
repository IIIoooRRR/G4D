package gateway

import (
	"context"
	"errors"

	"github.com/IIIoooRRR/G4D/model/gateway"
	"go.uber.org/zap"
)

func (r *Receiver) CreateBot(parentCtx context.Context, logger *zap.Logger, token *string) error {
	r.logger = logger                        // root for gateway
	r.dLogger = r.logger.Named("dispatcher") // dispatch.go logger

	if parentCtx == nil {
		return errors.New("ParentCtx must be initialized before calling")
	}

	if r.Queue == nil {
		if r.QueueSize == 0 {
			logger.Info("Queue size is zero. An unlimited queue has been created. To fix it, define the QueueSize parameter")
			r.Queue = make(chan *gateway.RawEvent)
		} else {
			r.Queue = make(chan *gateway.RawEvent, r.QueueSize)
		}
	}
	r.token = *token
	for {
		err := r.connect(parentCtx)
		if err != nil {
			logger.Error("connecting", zap.Error(err))
			r.lastSeq = 0
			r.sessionID = ""
		}
	}
	return nil
}
