package Parse

import (
	"encoding/json"

	"github.com/IIIoooRRR/G4D/G4D/Event"
)

func ToChannel(body []byte) (Event.Channel, error) {
	var channel Event.Channel
	err := json.Unmarshal(body, &channel)
	if err != nil {
		return Event.Channel{}, err
	}
	return channel, nil
}
