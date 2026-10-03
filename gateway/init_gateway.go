package gateway

import (
	"context"
	"sync"

	"go.uber.org/zap"
)

func (r *Receiver) InitGateway(ctx context.Context, logger *zap.Logger, token *string) error {
	var err error
	r.initOnce.Do(func() {
		r.logger = logger                        // root for gateway
		r.dLogger = r.logger.Named("dispatcher") // dispatch.go logger
		r.ctx = ctx
		r.token = token
		if ctx == nil {
			r.logger.Info("the bot will be disabled at the first failure, as the context has not been initialized")
		}
		for {
			if r.ctx != nil {
				select {
				case <-r.ctx.Done():
					return
				default:
					err = r.connect()
					if err != nil {
						r.logger.Error("error", zap.Error(err))
					}
				}
			} else {
				err = r.connect()
				if err != nil {
					r.logger.Error("error", zap.Error(err))
				}
			}
		}
	})
	r.initOnce = sync.Once{}
	return err
}
