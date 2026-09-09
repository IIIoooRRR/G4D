package gateway

import (
	"context"

	"go.uber.org/zap"
)

func (r *Receiver) InitGateway(logger *zap.Logger, token *string) error {
	var err error
	r.initOnce.Do(func() {
		r.logger = logger                        // root for gateway
		r.dLogger = r.logger.Named("dispatcher") // dispatch.go logger

		r.token = token

		for {
			err = r.connect(context.Background())
			if err != nil {
				r.logger.Error("error", zap.Error(err))
			}
			r.logger.Info("reconnect to discord")
		}
	})
	return err
}
