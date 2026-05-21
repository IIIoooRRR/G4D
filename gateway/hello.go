package gateway

import (
	"encoding/json"
	"log"

	"time"

	json2 "github.com/IIIoooRRR/G4D/model/codec"
)

func (r *Receiver) helloDiscord() time.Duration {
	var hello json2.Payload

	if err := r.connectWS.ReadJSON(&hello); err != nil {
		log.Println("[HELLO DISCORD]", err)
		return -1

	}
	var d json2.Hello
	err := json.Unmarshal(hello.D, &d)
	if err != nil {
		log.Println("[DISCORD] parse to hello is bad")
		return -1
	}
	r.interval = time.Duration(d.HeartbeatInterval) * time.Millisecond
	return r.interval
}
