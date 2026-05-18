package Connect

import "log"

func (b *Receiver) Stop() {
	if b.ctx != nil || b.cancel != nil {
		b.cancel()
	}
	if b.connectWS != nil {
		b.connectWS.Close()
	}
	log.Println("[BOT] Disconnected from Discord")
}
