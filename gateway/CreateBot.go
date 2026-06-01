package gateway

import (
	"context"
	"errors"
	"log"
)

func (r *Receiver) CreateBot(ParentCtx context.Context, token *string) error {
	if ParentCtx == nil {
		return errors.New("[BOT CREATE] ParentCtx must be initialized before calling")
	}
	if r.Queue == nil {
		if r.QueueSize == 0 {
			log.Println("[BOT CREATE] Queue size is zero. An unlimited queue has been created. To fix it, define the QueueSize parameter.")
			r.Queue = make(chan *RawEvent)
		} else {
			r.Queue = make(chan *RawEvent, r.QueueSize) // создание очереди
		}
	}
	r.token = *token
	for {
		err := r.connect(ParentCtx)
		if err != nil {
			log.Println("[BOT CREATE] Error connecting: ", err)
			r.lastSeq = 0
			r.sessionID = ""
		}
	}
	return nil
}
