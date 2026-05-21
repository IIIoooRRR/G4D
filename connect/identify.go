package connect

import (
	"encoding/json"
	"log"

	"github.com/IIIoooRRR/G4D/JSON"
)

func (r *Receiver) identify() error {
	Data := JSON.Identify{
		Token:   r.token,
		Intents: r.Intents,
		Properties: JSON.IdentifyProperties{
			OS:      "linux",
			Browser: "G4D",
			Device:  "G4D",
		},
		Presence: r.Presence,
	}
	DataBytes, err := json.Marshal(&Data)
	if err != nil {
		log.Println("[IDENTIFY] marshalling data error")
		return err
	}
	identify := JSON.Payload{
		Op: 2,
		D:  DataBytes,
	}
	r.connMutex.Lock()
	err = r.connectWS.WriteJSON(&identify)
	r.connMutex.Unlock()
	if err != nil {
		r.Stop()
		return err

	}
	return nil
}
