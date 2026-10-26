package gateway

import (
	"context"
	"sync"

	"go.uber.org/zap"
)

// it connects to Discord, performing a loop if the context is established
func (r *Receiver) InitGateway(ctx context.Context, logger *zap.Logger, token *string) error {
	var err error
	r.initOnce.Do(func() {
		r.logger = logger                        // root for gateway
		r.dLogger = r.logger.Named("dispatcher") // dispatch.go logger
		r.ctx = ctx
		r.token = token
		for {
			select {
			case <-r.ctx.Done():
				return
			default:
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
