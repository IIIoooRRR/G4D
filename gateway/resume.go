package gateway

import (
	"fmt"

	"github.com/IIIoooRRR/G4D/gateway/internal"
	"go.uber.org/zap"
)

/*
It is triggered by opcode 7, which comes to listen
and helps you reconnect without losing your events
Discord will simply send events from the last sequence provided to you by discord
*/
func (r *Receiver) resume() error {
	logger := r.logger.Named("resume")
	if r.sessionID != "" {
		resumePackage := json.Resume{
			Op: 6,
			Data: json.RData{
				Token:     *r.token,
				SessionID: r.sessionID,
				Sequence:  r.lastSeq.Load(),
			},
		}

		r.connMutex.Lock()
		var err error
		if err = r.connectWS.Close(); err != nil {
			logger.Warn("close old socket", zap.Error(err))
		}
		if err = r.gateway(); err != nil {
			r.connMutex.Unlock()
			return fmt.Errorf("gateway: %w", err)
		}
		if err := r.connectWS.WriteJSON(resumePackage); err != nil {
			r.connMutex.Unlock()
			return fmt.Errorf("write resume: %w", err)
		}
	}
	return nil
}
