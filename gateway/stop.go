package gateway

import (
	"go.uber.org/zap"
)

func (r *Receiver) Stop() {
	if r.ctx != nil || r.cancel != nil {
		r.cancel()
	}
	if r.connectWS != nil {
		err := r.connectWS.Close()
		if err != nil {
			r.logger.Error("failed to close websocket connection", zap.Error(err))
		}
	}
	r.logger.Info("Disconnected from Discord")
}
