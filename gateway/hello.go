package gateway

import (
	"time"

	"github.com/IIIoooRRR/G4D/gateway/internal"
	"github.com/IIIoooRRR/G4D/model/parse"
)

/*
it is used when a socket connection is created
It is needed to get the interval at which heartbeat will send its messages
*/
func (r *Receiver) helloDiscord() error {
	var hello json.Payload
	if err := r.connectWS.ReadJSON(&hello); err != nil {
		return err
	}
	var d json.Hello
	err := parse.Unmarshal(hello.D, &d)
	if err != nil {
		return err
	}
	r.interval = time.Duration(d.HeartbeatInterval) * time.Millisecond
	return nil
}
