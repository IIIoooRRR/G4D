package gateway

import (
	"log"

	json "github.com/IIIoooRRR/G4D/model/codec"
	gw "github.com/IIIoooRRR/G4D/model/gateway"
	"github.com/IIIoooRRR/G4D/model/parse"
)

/*
the dispatcher of all events that come to our program via the socket.
It serves as a consumer for EventQueue,
which ensures the operation of the event processor
restrict access to the READY event because it is unsafe for the user to work with it
There may be a data leak
All other events are available to the user, he can manage them as he wants
*/
func (r *Receiver) dispatch(event json.Payload) error {

	t := event.T
	log.Println(t)
	switch t {
	case "READY":
		var d json.ReadyEvent
		if err := parse.Unmarshal(event.D, &d); err != nil {
			log.Println("[DISPATCH] ", err)
		}
		log.Println("[DISPATCH] g4d is ready")
		r.sessionID = d.SessionID

		r.resumeURL = d.ResumeGatewayURL
	default:
		r.Queue <- &gw.RawEvent{
			Type: t,
			Data: event.D,
		}
	}
	r.lastSeq = event.S
	return nil
}
