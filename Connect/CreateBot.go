package Connect

import (
	"context"
	"errors"
	"log"
)

func (b *Receiver) CreateBot(ParentCtx context.Context, token *string) error {
	if ParentCtx == nil {
		return errors.New("[BOT CREATE] ParentCtx must be initialized before calling")
	}
	if b.QueueSize == 0 {
		log.Println("[BOT CREATE] Queue size is zero. An unlimited queue has been created. To fix it, define the QueueSize parameter.")
		b.Queue = make(chan *RawEvent)
	} else {
		b.Queue = make(chan *RawEvent, b.QueueSize) // создание очереди
	}
	b.token = *token
	for {
		err := b.connect(ParentCtx)
		if err != nil {
			log.Println("[BOT CREATE] Error connecting: ", err)
			return err

		}
	}
	return nil
}
