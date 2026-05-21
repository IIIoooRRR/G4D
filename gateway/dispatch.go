package gateway

import (
	"encoding/json"
	"log"

	json2 "github.com/IIIoooRRR/G4D/model/codec"
)

func (r *Receiver) dispatch(event json2.Payload) error {
	t := event.T
	log.Println(t)
	switch t {
	case "READY":
		var d json2.ReadyEvent
		if err := json.Unmarshal(event.D, &d); err != nil {
			log.Println("[DISPATCH] ", err)
		}
		log.Println("[DISPATCH] g4d is ready")
		r.sessionID = d.SessionID

		r.resumeURL = d.ResumeGatewayURL

	case "GUILD_CREATE":
		if r.Cache != nil {
			r.Cache.CacheGuildCreate(&RawEvent{t, event.D})
			break
		}
		r.Queue <- &RawEvent{
			t,
			event.D,
		}
		return nil
	default:
		r.Queue <- &RawEvent{
			t,
			event.D,
		}
	}
	r.lastSeq = event.S
	return nil
}
