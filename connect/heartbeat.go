package Connect

import (
	"context"
	"log"
	"time"

	"github.com/IIIoooRRR/G4D/JSON"
)

func (b *Receiver) heartbeat(ctx context.Context) error {
	ticker := time.NewTicker(b.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			b.connMutex.Lock()
			err := b.connectWS.WriteJSON(
				JSON.Payload{
					Op: 1,
					S:  b.lastSeq,
				})
			b.connMutex.Unlock()
			if err != nil {
				return err
			}
		case <-ctx.Done():
			log.Println("[HEARTBEAT] HEARTBEAT DONE")
			return nil
		}
	}
}
