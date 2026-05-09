package Connect

import (
	"log"

	"github.com/gorilla/websocket"
)

func (b *Receiver) gateway() error {
	con, _, err := websocket.DefaultDialer.Dial("wss://gateway.discord.gg/?v=10&encoding=json", nil)
	if err != nil {
		log.Println("[GATEWAY] NOT CONNECT TO WEBSOCKET")
		return err
	}
	b.connectWS = con // подключаемся
	log.Println("[GATEWAY] CONNECTED")
	return nil
}
