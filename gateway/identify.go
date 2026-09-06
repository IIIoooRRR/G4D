package gateway

import (
	"runtime"

	"github.com/IIIoooRRR/G4D/gateway/internal"
	"github.com/IIIoooRRR/G4D/model/parse"
)

/*
the very first event in the entire web socket
sends data about what kind of bot, where and how it connects
*/
func (r *Receiver) identify() error {
	Data := json.Identify{
		Token:   *r.token,
		Intents: int(r.Intents),
		Properties: json.IdentifyProperties{
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
	identify := json.Payload{
		Op: 2,
		D:  DataBytes,
	}
	r.connMutex.Lock()
	err = r.connectWS.WriteJSON(&identify)
	r.connMutex.Unlock()
	if err != nil {
		return err

	}
	return nil
}
