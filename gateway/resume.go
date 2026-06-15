package gateway

import (
	"encoding/json"

	json2 "github.com/IIIoooRRR/G4D/model/codec"
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
		Data := json2.Resume{
			Token:     r.token,
			SessionID: r.sessionID,
			Sequence:  r.lastSeq,
		}
		DataBytes, err := json.Marshal(&Data)
		if err != nil {
			logger.Error("marshal error:", zap.Error(err))
			return err
		}
		answerToDiscord :=
			json2.Payload{
				Op: 6,
				D:  DataBytes,
			}
		r.connMutex.Lock()
		err = r.connectWS.WriteJSON(&answerToDiscord)
		r.connMutex.Unlock()
		if err != nil {
			return err
		}
	}
	return nil
}
