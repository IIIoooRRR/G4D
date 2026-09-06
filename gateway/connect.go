package gateway

import (
	"context"

	"github.com/IIIoooRRR/G4D/gateway/internal"
	"go.uber.org/zap"
)

func (r *Receiver) connect(ParentCtx context.Context) error {
	defer r.Stop()                                  // if there is an error, we roll back the ones specified in r.Stop parts of sockets
	r.ctx, r.cancel = context.WithCancel(ParentCtx) //creating a context based on the parent
	err := r.gateway()
	if err != nil {
		return err
	}

	err = r.helloDiscord()
	if err != nil {
		return err
	}
	go func() {
		err := r.heartbeat(r.ctx) // tell the program to send heartbeat messages every n seconds.
		if err != nil {
			r.logger.Error("heartbeat", zap.Error(err))
		}
	}()
	err = r.identify() // sending identification messages to discord
	if err != nil {
		return err
	}
	err = r.listen(r.ctx, r.logger.Named("connect")) // start listening and processing events from web sockets (the main program flow)
	if err != nil {
		return err
	}
	return nil
}

func (r *Receiver) listen(ctx context.Context, logger *zap.Logger) error {
	defer func() {
		_ = r.connectWS.Close()
	}()
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
				go func(event json.Payload) {
					err := r.dispatch(event) // all basic events have opcode == 0, we transfer control to dispatch
					if err != nil {
						logger.Error("case 0 error", zap.Error(err))
					}
				}(event)
			case 1:
				r.connMutex.Lock()
				err = r.connectWS.WriteJSON(json.Payload{ // it tells you what interval heartbeat.go should work with.
					Op: 1,
					S:  int(r.lastSeq.Load()),
				})
				r.connMutex.Unlock()
				if err != nil {
					r.logger.Error("case 1 error", zap.Error(err))
					return err
				}
			case 7:
				if err := r.resume(); err != nil {
					logger.Error("Resume failed, full reconnect", zap.Error(err))
					r.sessionID = ""
					return err
				}
			case 9:
				r.connMutex.Lock()
				err := r.connectWS.Close()
				if err != nil {
					logger.Error("case 9 error", zap.Error(err))
				}
				r.connMutex.Unlock()
				r.sessionID = ""
				r.logger.Info("reconnect to discord")
				return nil // 9, 7 - reconnect, hard, or resume
			}
		}
	}
}
