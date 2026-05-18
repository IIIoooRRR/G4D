package Connect

import (
	"encoding/json"
	"log"

	"time"

	"github.com/IIIoooRRR/G4D/JSON"
)

func (b *Receiver) helloDiscord() time.Duration {
	var hello JSON.Payload

	if err := b.connectWS.ReadJSON(&hello); err != nil {
		return -1
	}
	var d JSON.Hello
	err := json.Unmarshal(hello.D, &d)
	if err != nil {
		log.Println("[DISCORD] parse to hello is bad")
		return -1
	}
	b.interval = time.Duration(d.HeartbeatInterval) * time.Millisecond
	return b.interval
}
