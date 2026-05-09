package Connect

import (
	"encoding/json"
	"log"

	"github.com/IIIoooRRR/G4D/JSON"
)

func (b *Receiver) resume() error {
	if b.sessionID != "" {
		Data := JSON.Resume{
			Token:     b.token,
			SessionID: b.sessionID,
			Sequence:  b.lastSeq,
		}
		DataBytes, err := json.Marshal(&Data)
		if err != nil {
			log.Println("[RESUME] marshal error:", err)
			return err
		}
		answerToDiscord :=
			JSON.Payload{
				Op: 6,
				D:  DataBytes,
			}
		b.connMutex.Lock()
		err = b.connectWS.WriteJSON(&answerToDiscord)
		b.connMutex.Unlock()
		if err != nil {
			return err
		}
	}
	return nil
}
