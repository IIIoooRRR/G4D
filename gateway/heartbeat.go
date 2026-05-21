package gateway

import (
	"context"
	"log"
	"time"

	"github.com/IIIoooRRR/G4D/model/codec"
)

func (r *Receiver) heartbeat(ctx context.Context) error {
	ticker := time.NewTicker(r.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			r.connMutex.Lock()
			err := r.connectWS.WriteJSON(
				json.Payload{
					Op: 1,
					S:  r.lastSeq,
				})
			r.connMutex.Unlock()
			if err != nil {
				return err
			}
		case <-ctx.Done():
			log.Println("[HEARTBEAT] HEARTBEAT DONE")
			return nil
		}
	}
}
