package gateway

import (
	"encoding/json"
	"log"

	json2 "github.com/IIIoooRRR/G4D/model/codec"
)

func (r *Receiver) identify() error {
	Data := json2.Identify{
		Token:   r.token,
		Intents: r.Intents,
		Properties: json2.IdentifyProperties{
			OS:      "linux",
			Browser: "g4d",
			Device:  "g4d",
		},
		Presence: r.Presence,
	}
	DataBytes, err := json.Marshal(&Data)
	if err != nil {
		log.Println("[IDENTIFY] marshalling data error")
		return err
	}
	identify := json2.Payload{
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
