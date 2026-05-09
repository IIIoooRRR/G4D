package Connect

import (
	"encoding/json"
	"log"

	"github.com/IIIoooRRR/G4D/JSON"
)

func (b *Receiver) dispatch(event JSON.Payload) error {
	t := event.T

	if t == "READY" {
		var d JSON.ReadyEvent
		if err := json.Unmarshal(event.D, &d); err != nil {
			log.Println("[DISPATCH] ", err)
		}
		log.Println("[DISPATCH] G4D is ready")
		b.sessionID = d.SessionID

		b.resumeURL = d.ResumeGatewayURL
	} else {

		b.Queue <- &RawEvent{
			t,
			event.D,
		}
	}
	b.lastSeq = event.S
	return nil
}
