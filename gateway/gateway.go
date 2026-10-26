package gateway

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func (r *Receiver) gateway(isResume bool) error {
	var con *websocket.Conn
	var err error
	if isResume {
		con, _, err = websocket.DefaultDialer.Dial(r.resumeURL, nil)
	} else {
		con, _, err = websocket.DefaultDialer.Dial("wss://gateway.discord.gg/?v=10&encoding=json", nil)
	}
	if err != nil {
		r.logger.Error("not connected to w-socket", zap.Error(err))
		return err
	}
	r.connectWS = con // подключаемся
	r.logger.Info("connected to w-socket")
	return nil
}
