package Connect

import (
	"encoding/json"
	"log"

	"github.com/IIIoooRRR/G4D/JSON"
)

func (b *Receiver) identify() error {
	Data := JSON.Identify{
		Token:   b.token,
		Intents: b.Intents,
		Properties: JSON.IdentifyProperties{
			OS:      "linux",
			Browser: "G4D",
			Device:  "G4D",
		},
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
	b.connMutex.Lock()
	err = b.connectWS.WriteJSON(&identify)
	b.connMutex.Unlock()
	if err != nil {
		b.Stop()
		return err

	}
	return nil
}
