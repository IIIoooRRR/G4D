package gateway

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

func (r *Receiver) InitGateway(parentCtx context.Context, logger *zap.Logger, token *string) error {
	r.logger = logger                        // root for gateway
	r.dLogger = r.logger.Named("dispatcher") // dispatch.go logger

	if parentCtx == nil {
		return errors.New("ParentCtx must be initialized before calling")
	}
	r.token = token
	for {
		err := r.connect(parentCtx)
		if err != nil {
			logger.Error("connecting", zap.Error(err))
			r.lastSeq.Store(0)
			r.sessionID = ""
		}
	}
	return nil
}
