package connect

import "log"

func (r *Receiver) Stop() {
	if r.ctx != nil || r.cancel != nil {
		r.cancel()
	}
	if r.connectWS != nil {
		r.connectWS.Close()
	}
	log.Println("[BOT] Disconnected from Discord")
}
