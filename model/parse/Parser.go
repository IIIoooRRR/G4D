package parse

import (
	"github.com/IIIoooRRR/G4D/model/gateway"
	"github.com/IIIoooRRR/G4D/model/shema"
	"go.uber.org/zap"
)

func Event[T any](event *gateway.RawEvent) *T {
	var d *T
	if event == nil {
		return nil
	}
	err := Unmarshal(event.Data, &d)
	if err != nil {
		logger.Error("unmarshal raw event", zap.Error(err))
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
