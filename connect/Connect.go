package Connect

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/IIIoooRRR/G4D/JSON"
	"github.com/gorilla/websocket"
)

func (b *Receiver) connect(ParentCtx context.Context) error {
	sleep := 1
	b.ctx, b.cancel = context.WithCancel(ParentCtx) //создаем контекст на основе родительского
	defer b.cancel()
	for {
		err := b.gateway()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
		if sleep < 60 {
			sleep = sleep * 2
		} // создаем цикл для гатеваев
	}

	interval := b.helloDiscord()
	if interval == -1 {
		return errors.New("[DISCORD] hello json is bad :[")
	}

	go b.heartbeat(b.ctx)
	err := b.identify()
	if err != nil {
		return err
	}
	err = b.listen(b.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (b *Receiver) listen(ctx context.Context) error {
	defer func(connectWS *websocket.Conn) {
		err := connectWS.Close()
		if err != nil {

		}
	}(b.connectWS)
	for {
		select {
		case <-ctx.Done():
			return nil
		default:

			var event JSON.Payload
			err := b.connectWS.ReadJSON(&event)
			if err != nil {
				return err
			}
			op := event.Op
			switch op {
			case 0:
				go func() {
					err = b.dispatch(event)
					if err != nil {
						log.Println(err)
					}
				}()
			case 1:
				b.connMutex.Lock()
				err := b.connectWS.WriteJSON(JSON.Payload{
					Op: 1,
					S:  b.lastSeq,
				})
				b.connMutex.Unlock()
				log.Println("[CONNECT] last Seq: ", b.lastSeq)
				if err != nil {
					log.Println("[CONNECT] ", err)
					return err
				}
			case 7:
				b.sessionID = ""
				return errors.New("[LISTEN] HARD RECONNECT TO DISCORD")
			case 9:
				b.sessionID = ""
				return errors.New("[LISTEN] RECONNECT TO DISCORD")
			}
		}
	}
}
