package gateway

import (
	"runtime"

	json2 "github.com/IIIoooRRR/G4D/model/codec"
	"github.com/IIIoooRRR/G4D/model/parse"
)

/*
the very first event in the entire web socket
sends data about what kind of bot, where and how it connects
*/
func (r *Receiver) identify() error {
	Data := json2.Identify{
		Token:   *r.token,
		Intents: r.Intents,
		Properties: json2.IdentifyProperties{
			OS:      runtime.GOOS,
			Browser: "g4d",
			Device:  "g4d",
		},
		Presence: r.Presence, // activity, streaming...
	}
	DataBytes, err := parse.Marshal(&Data)
	if err != nil {
		r.logger.Info("marshalling data error in identify")
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
