package gateway

import (
	"log"

	"github.com/gorilla/websocket"
)

func (r *Receiver) gateway() error {
	con, _, err := websocket.DefaultDialer.Dial("wss://gateway.discord.gg/?v=10&encoding=json", nil)
	if err != nil {
		log.Println("[GATEWAY] NOT CONNECT TO WEBSOCKET")
		return err
	}
	r.connectWS = con // подключаемся
	log.Println("[GATEWAY] CONNECTED")
	return nil
}
