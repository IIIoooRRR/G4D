package parse

import (
	"log"

	"github.com/IIIoooRRR/G4D/model/gateway"
	"github.com/IIIoooRRR/G4D/model/shema"
)

func Event[T any](event *gateway.RawEvent) *T {
	var d *T
	err := Unmarshal(event.Data, &d)
	if err != nil {
		log.Printf("Error unmarshal raw event: %v", err)
		return nil
	}
	return d
}

/* Channel */
func ToChannel(body []byte) (*shema.Channel, error) {
	var channel *shema.Channel
	err := Unmarshal(body, channel)
	if err != nil {
		return nil, err
	}
	return channel, nil
}
