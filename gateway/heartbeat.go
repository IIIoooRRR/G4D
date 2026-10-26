package gateway

import (
	"time"

	"github.com/IIIoooRRR/G4D/gateway/internal"
)

/*
a function that takes every n seconds that were specified when connecting
It consumes almost no CPU or RAM, as it is almost always waiting.
*/
func (r *Receiver) heartbeat() error {
	logger := r.logger.Named("heartbeat")

	ticker := time.NewTicker(r.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			r.connMutex.Lock()
			err := r.connectWS.WriteJSON(
				json.Payload{
					Op: 1,
					S:  r.lastSeq.Load(),
				})
			r.connMutex.Unlock()
			if err != nil {
				return err
			}
		case <-r.ctx.Done():
			logger.Info("heartbeat stopped")
			return nil
		}
	}
}
