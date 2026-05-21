package gateway

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/IIIoooRRR/G4D/model/codec"
	"github.com/gorilla/websocket"
)

func (r *Receiver) connect(ParentCtx context.Context) error {
	sleep := 1
	r.ctx, r.cancel = context.WithCancel(ParentCtx) //создаем контекст на основе родительского
	defer r.cancel()
	for {
		err := r.gateway()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
		if sleep < 60 {
			sleep = sleep * 2
		} // создаем цикл для гатеваев
	}

	interval := r.helloDiscord()
	if interval <= 0 {
		return errors.New("[DISCORD] hello json is bad :[")
	}
	go r.heartbeat(r.ctx)
	err := r.identify()
	if err != nil {
		return err
	}
	err = r.listen(r.ctx)
	if err != nil {
		return err
	}
	defer func() {

	}()
	return nil
}

func (r *Receiver) listen(ctx context.Context) error {
	defer func(connectWS *websocket.Conn) {
		err := connectWS.Close()
		if err != nil {

		}
	}(r.connectWS)
	for {
		select {
		case <-ctx.Done():
			return nil
		default:

			var event json.Payload
			err := r.connectWS.ReadJSON(&event)
			if err != nil {
				return err
			}
			op := event.Op
			switch op {
			case 0:
				go func() {
					err = r.dispatch(event)
					if err != nil {
						log.Println(err)
					}
				}()
			case 1:
				r.connMutex.Lock()
				err := r.connectWS.WriteJSON(json.Payload{
					Op: 1,
					S:  r.lastSeq,
				})
				r.connMutex.Unlock()
				log.Println("[CONNECT] last Seq: ", r.lastSeq)
				if err != nil {
					log.Println("[CONNECT] ", err)
					return err
				}
			case 7:
				r.sessionID = ""
				return errors.New("[LISTEN] HARD RECONNECT TO DISCORD")
			case 9:
				r.sessionID = ""
				return errors.New("[LISTEN] RECONNECT TO DISCORD")
			}
		}
	}
}
