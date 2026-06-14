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
	r.ctx, r.cancel = context.WithCancel(ParentCtx) //creating a context based on the parent
	for {
		err := r.gateway()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
		if sleep < 60 {
			sleep = sleep * 2
		}
	}

	err := r.helloDiscord()
	if err != nil {
		return err
	}
	go r.heartbeat(r.ctx) // tell the program to send heartbeat messages every n seconds.
	err = r.identify()    // sending identification messages to discord
	if err != nil {
		return err
	}
	err = r.listen(r.ctx) // start listening and processing events from web sockets (the main program flow)
	if err != nil {
		return err
	}
	r.Stop() // if there is an error, we roll back the ones specified in r.Stop parts of sockets
	return nil
}

func (r *Receiver) listen(ctx context.Context) error {
	defer func(connectWS *websocket.Conn) {
		_ = connectWS.Close()
	}(r.connectWS)
	for {
		select {
		case <-ctx.Done():
			return nil
		default:

			var event json.Payload
			err := r.connectWS.ReadJSON(&event) // start processing messages, it works via web-socket/mozilla
			if err != nil {
				return err
			}
			op := event.Op
			switch op {
			case 0:
				/*
					you should use go func to quickly release the main thread,
					since the processes themselves can be very long or slow down the program itself (affecting 10+ events per second
				*/
				go func() {
					err = r.dispatch(event) // all basic events have opcode == 0, we transfer control to dispatch
					if err != nil {
						log.Println(err)
					}
				}()
			case 1:
				r.connMutex.Lock()
				err := r.connectWS.WriteJSON(json.Payload{ // it tells you what interval heartbeat.go should work with.
					Op: 1,
					S:  r.lastSeq,
				})
				r.connMutex.Unlock()
				if err != nil {
					log.Println("[CONNECT] ", err)
					return err
				}
			case 7:
				if err := r.resume(); err != nil {
					log.Printf("[GATEWAY] Resume failed: %v, full reconnect", err)
					r.sessionID = ""
					return err
				}
			case 9:
				r.connMutex.Lock()
				r.connectWS.Close()
				r.connMutex.Unlock()
				r.sessionID = ""
				return errors.New("[LISTEN] RECONNECT TO DISCORD") // 9, 7 - reconnect, hard, or resume
			}
		}
	}
}
