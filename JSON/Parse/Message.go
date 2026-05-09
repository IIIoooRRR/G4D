package Parse

import (
	"encoding/json"
	"log"

	"github.com/IIIoooRRR/G4D/Connect"
	"github.com/IIIoooRRR/G4D/G4D/Event"
)

func ToMessageCreate(event Connect.RawEvent) Event.Message {
	var d Event.Message
	err := json.Unmarshal(event.Data, &d)
	if err != nil {
		log.Println(err)
	}
	return d
}
func ToMessageDelete(event Connect.RawEvent) Event.MessageDelete {
	var d Event.MessageDelete
	err := json.Unmarshal(event.Data, &d)
	if err != nil {
		log.Println(err)
	}
	return d
}
