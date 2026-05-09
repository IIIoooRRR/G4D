package Parse

import (
	"encoding/json"

	"github.com/IIIoooRRR/G4D/Connect"
	"github.com/IIIoooRRR/G4D/G4D/Event"
)

func ToInteraction(event Connect.RawEvent) Event.Interaction {
	var d Event.Interaction
	json.Unmarshal(event.Data, &d)
	return d
}
