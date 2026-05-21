package connect

import (
	"encoding/json"
	"log"

	"github.com/IIIoooRRR/G4D/JSON"
)

func (r *Receiver) dispatch(event JSON.Payload) error {
	t := event.T
	switch t {
	case "READY":
		var d JSON.ReadyEvent
		if err := json.Unmarshal(event.D, &d); err != nil {
			log.Println("[DISPATCH] ", err)
		}
		log.Println("[DISPATCH] G4D is ready")
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
