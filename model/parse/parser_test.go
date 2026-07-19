package parse

import (
	"reflect"
	"sync"
	"testing"

	"github.com/IIIoooRRR/G4D/model/gateway"
	"github.com/IIIoooRRR/G4D/model/parse/types"
	"github.com/IIIoooRRR/G4D/model/schema"
)

func TestEvent_GetMessage(t *testing.T) {
	jsonData := []byte(`{
		"id": "123456789",
		"channel_id": "987654321",
		"content": "!hello world",
		"author": {
			"id": "111222333",
			"username": "TestUser",
			"global_name": "tester"
		}
	}`)

	event := &gateway.RawEvent{
		Type: "MESSAGE_CREATE",
		Data: jsonData,
	}

	wg := sync.WaitGroup{}
	AddEvent(event, &wg, 1, reflect.TypeOf(schema.GetMessage{}))
	msg := GetEvent[schema.GetMessage](event)

	if msg.ID != "123456789" {
		t.Errorf("Expected ID '123456789', got '%s'", msg.ID)
	}
	if msg.ChannelID != "987654321" {
		t.Errorf("Expected ChannelID '987654321', got '%s'", msg.ChannelID)
	}
	if msg.Content != "!hello world" {
		t.Errorf("Expected Content '!hello world', got '%s'", msg.Content)
	}
	if string(msg.Author.Id) != "111222333" {
		t.Errorf("Expected Author.ID '111222333', got '%s'", msg.Author.Id)
	}
}

func TestEvent_MessageDelete(t *testing.T) {
	jsonData := []byte(`{
		"id": "123456789",
		"channel_id": "987654321"
	}`)

	event := &gateway.RawEvent{
		Type: "MESSAGE_DELETE",
		Data: jsonData,
	}

	wg := sync.WaitGroup{}
	AddEvent(event, &wg, 1, types.Get(event.Type))
	deleted := GetEvent[schema.MessageDelete](event)

	if deleted.ID != "123456789" {
		t.Errorf("Expected ID '123456789', got '%s'", deleted.ID)
	}
	if string(deleted.ChannelID) != "987654321" {
		t.Errorf("Expected ChannelID '987654321', got '%s'", deleted.ChannelID)
	}
}

func TestEvent_Interaction(t *testing.T) {
	jsonData := []byte(`{
		"id": "23",
		"type": 2,
		"data": {
			"name": "ping",
			"type": 1
		},
		"member": {
			"user": {
				"id": "111222333"
			}
		}
	}`)
	event := &gateway.RawEvent{
		Type: "INTERACTION_CREATE",
		Data: jsonData,
	}
	wg := sync.WaitGroup{}
	AddEvent(event, &wg, 1, types.Get(event.Type))
	interaction := GetEvent[schema.Interaction](event)

	if interaction.ID != "23" {
		t.Errorf("Expected ID '23', got '%v'", interaction.ID)
	}
	if interaction.Type != 2 {
		t.Errorf("Expected Type 2, got %d", interaction.Type)
	}
}

func BenchmarkEvent_GetMessage(b *testing.B) {
	jsonData := []byte(`{
		"id": "123456789",
		"channel_id": "987654321",
		"content": "!hello world",
		"author": {"id": "111222333", "username": "TestUser"}
	}`)

	event := &gateway.RawEvent{
		Type: "MESSAGE_CREATE",
		Data: jsonData,
	}

	wg := sync.WaitGroup{}
	AddEvent(event, &wg, 1, types.Get(event.Type))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetEvent[schema.GetMessage](event)
	}
}

func TestEvent_TableDriven(t *testing.T) {
	tests := []struct {
		name      string
		eventType string
		jsonData  string
		wantID    string
		wantType  interface{} // ← добавляем ожидаемый тип
	}{
		{
			name:      "message create",
			eventType: "MESSAGE_CREATE",
			jsonData:  `{"id":"111","channel_id":"222","content":"test"}`,
			wantID:    "111",
			wantType:  schema.GetMessage{},
		},
		{
			name:      "message delete",
			eventType: "MESSAGE_DELETE",
			jsonData:  `{"id":"333","channel_id":"444"}`,
			wantID:    "333",
			wantType:  schema.MessageDelete{},
		},
		// ...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := &gateway.RawEvent{
				Type: tt.eventType,
				Data: []byte(tt.jsonData),
			}

			wg := sync.WaitGroup{}
			AddEvent(event, &wg, 1, types.Get(tt.eventType))
			var id string
			switch tt.eventType {
			case "MESSAGE_CREATE":
				msg := GetEvent[schema.GetMessage](event)
				id = string(msg.ID)
			case "MESSAGE_DELETE":
				msg := GetEvent[schema.MessageDelete](event)
				id = string(msg.ID)
			case "INTERACTION_CREATE":
				msg := GetEvent[schema.Interaction](event)
				id = msg.ID
			}

			if id != tt.wantID {
				t.Errorf("Expected ID '%s', got '%s'", tt.wantID, id)
			}
		})
	}
}
