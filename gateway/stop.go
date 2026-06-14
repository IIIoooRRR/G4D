package gateway

import "log"

func (r *Receiver) Stop() {
	if r.ctx != nil || r.cancel != nil {
		r.cancel()
	}
	if r.connectWS != nil {
		err := r.connectWS.Close()
		if err != nil {
			log.Printf("failed to close websocket connection: %v", err)
		}
	}
	log.Println("[BOT] Disconnected from Discord")
}
