package gateway

import (
	"time"

	json2 "github.com/IIIoooRRR/G4D/model/codec"
	"github.com/IIIoooRRR/G4D/model/parse"
)

/*
it is used when a socket connection is created
It is needed to get the interval at which heartbeat will send its messages
*/
func (r *Receiver) helloDiscord() error {
	var hello json2.Payload
	if err := r.connectWS.ReadJSON(&hello); err != nil {
		return err
	}
	var d json2.Hello
	err := parse.Unmarshal(hello.D, &d)
	if err != nil {
		return err
	}
	r.interval = time.Duration(d.HeartbeatInterval) * time.Millisecond
	return nil
}
