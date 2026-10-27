package parse

import (
	"sync"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/entity"
	"github.com/IIIoooRRR/G4D/model/schema"
	"go.uber.org/zap"
)

type Cache struct {
	Entry  []EventEntry
	logger *zap.Logger
}

func InitCache(quantity _const.Quantity, logger *zap.Logger) *Cache {
	return &Cache{
		Entry:  make([]EventEntry, quantity),
		logger: logger,
	}
}

func (c *Cache) AddEvent(event *RawEvent, wg *sync.WaitGroup, seq, qnt int, factory func() any) {
	wg.Add(qnt)
	c.Entry[seq] = EventEntry{
		Data: c.reflectParsing(event, factory()),
	}
	event.cache, event.idx = c.Entry, seq
}

func ToChannel(body []byte) (*schema.Channel, error) {
	var channel schema.Channel
	err := Unmarshal(body, &channel)
	if err != nil {
		return nil, err
	}
	return &channel, nil
}
func ToUser(body []byte) (*entity.User, error) {
	var user entity.User
	err := Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
