package gateway

import (
	"context"

	"go.uber.org/zap"
)

func (r *Receiver) InitGateway(logger *zap.Logger, token *string) error {
	r.logger = logger                        // root for gateway
	r.dLogger = r.logger.Named("dispatcher") // dispatch.go logger

	r.token = token
	var err error
	for {
		err = r.connect(context.Background())
		if err != nil {
			r.logger.Error("error", zap.Error(err))
		}
		r.logger.Info("reconnect to discord")
	}
	return err
}
