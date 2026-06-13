package gateway

import (
	"log"

	"time"

	json2 "github.com/IIIoooRRR/G4D/model/codec"
	"github.com/IIIoooRRR/G4D/model/parse"
)

func (r *Receiver) helloDiscord() time.Duration {
	var hello json2.Payload

	if err := r.connectWS.ReadJSON(&hello); err != nil {
		log.Println("[HELLO DISCORD]", err)
		return -1

	}
	var d json2.Hello
	err := parse.Unmarshal(hello.D, &d)
	if err != nil {
		log.Println("[DISCORD] parse to hello is bad")
		return -1
	}
	r.interval = time.Duration(d.HeartbeatInterval) * time.Millisecond
	return r.interval
}
