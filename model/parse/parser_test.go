package parse

import (
	"testing"

	"github.com/IIIoooRRR/G4D/model/gateway"
	"github.com/IIIoooRRR/G4D/model/shema"
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

	msg := Event[shema.GetMessage](event)

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

	deleted := Event[shema.MessageDelete](event)

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

	interaction := Event[shema.Interaction](event)

	if interaction.ID != "23" {
		t.Errorf("Expected ID '23', got '%v'", interaction.ID)
	}
	if interaction.Type != 2 {
		t.Errorf("Expected Type 2, got %d", interaction.Type)
	}
}

func TestEvent_NilEvent(t *testing.T) {
	result := Event[shema.GetMessage](nil)
	if result != nil {
		t.Error("Expected nil result for nil event")
	}
}

func TestEvent_InvalidJSON(t *testing.T) {
	event := &gateway.RawEvent{
		Type: "MESSAGE_CREATE",
		Data: []byte(`{invalid json`),
	}

	// Не должно паниковать
	result := Event[shema.GetMessage](event)

	// Должен вернуть nil или zero value в зависимости от реализации
	if result != nil {
		t.Log("Expected nil or zero value for invalid JSON")
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Event[shema.GetMessage](event)
	}
}

func TestEvent_TableDriven(t *testing.T) {
	tests := []struct {
		name      string
		eventType string
		jsonData  string
		wantID    string
	}{
		{
			name:      "message create",
			eventType: "MESSAGE_CREATE",
			jsonData:  `{"id":"111","channel_id":"222","content":"test"}`,
			wantID:    "111",
		},
		{
			name:      "message delete",
			eventType: "MESSAGE_DELETE",
			jsonData:  `{"id":"333","channel_id":"444"}`,
			wantID:    "333",
		},
		{
			name:      "interaction",
			eventType: "INTERACTION_CREATE",
			jsonData:  `{"id":"555","type":2}`,
			wantID:    "555",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := &gateway.RawEvent{
				Type: tt.eventType,
				Data: []byte(tt.jsonData),
			}

			msg := Event[shema.GetMessage](event)

			if string(msg.ID) != tt.wantID {
				t.Errorf("Expected ID '%s', got '%s'", tt.wantID, msg.ID)
			}
		})
	}
}
