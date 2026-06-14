package g4d_test

import (
	"testing"

	"github.com/IIIoooRRR/G4D/g4d"
	gw "github.com/IIIoooRRR/G4D/gateway"
	gateway "github.com/IIIoooRRR/G4D/model/gateway"
)

// Мок команды
type mockCommand struct {
	called bool
	err    error
}

func (m *mockCommand) Execute(event *gateway.RawEvent) error {
	m.called = true
	return m.err
}

// Тест добавления команд
func TestBot_AddCommands(t *testing.T) {
	bot := &g4d.Bot{
		CommandBuffer: []g4d.CommandTemplate{},
	}

	cmd := &mockCommand{}
	bot.AddCommands([]g4d.CommandTemplate{
		{Trigger: "MESSAGE_CREATE", Name: "test", Action: cmd.Execute},
	})

	if len(bot.CommandBuffer) != 1 {
		t.Errorf("Expected 1 command, got %d", len(bot.CommandBuffer))
	}
}

// Тест статического процессора
func TestBot_StaticEventProcessor(t *testing.T) {
	mockCmd := &mockCommand{}

	bot := &g4d.Bot{
		Gateway: &gw.Receiver{
			Queue: make(chan *gateway.RawEvent, 10),
		},
		CommandBuffer: []g4d.CommandTemplate{
			{Trigger: "MESSAGE_CREATE", Name: "test", Action: mockCmd.Execute},
		},
	}

	go bot.StaticEventProcessor(5)

	// Отправляем тестовое событие
	bot.Gateway.Queue <- &gateway.RawEvent{Type: "MESSAGE_CREATE"}

	// Ждём выполнения
	// В реальном тесте нужно добавить ожидание
}

// Тест динамического процессора
func TestBot_DynamicEventProcessor(t *testing.T) {
	mockCmd := &mockCommand{}

	bot := &g4d.Bot{
		Gateway: &gw.Receiver{
			Queue: make(chan *gateway.RawEvent, 10),
		},
		CommandBuffer: []g4d.CommandTemplate{
			{Trigger: "MESSAGE_CREATE", Name: "test", Action: mockCmd.Execute},
		},
	}

	go bot.DynamicEventProcessor(5)

	bot.Gateway.Queue <- &gateway.RawEvent{Type: "MESSAGE_CREATE"}
}

// Бенчмарк
func BenchmarkStaticEventProcessor(b *testing.B) {
	mockCmd := &mockCommand{}

	bot := &g4d.Bot{
		Gateway: &gw.Receiver{
			Queue: make(chan *gateway.RawEvent, 1000),
		},
		CommandBuffer: []g4d.CommandTemplate{
			{Trigger: "MESSAGE_CREATE", Name: "test", Action: mockCmd.Execute},
		},
	}

	go bot.StaticEventProcessor(10)

	for i := 0; i < b.N; i++ {
		bot.Gateway.Queue <- &gateway.RawEvent{Type: "MESSAGE_CREATE"}
	}
}
