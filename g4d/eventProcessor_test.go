package g4d_test

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/IIIoooRRR/G4D/g4d"
	gw "github.com/IIIoooRRR/G4D/gateway"
	gateway "github.com/IIIoooRRR/G4D/model/gateway"
	"go.uber.org/zap"
)

type mockCommand struct {
	called atomic.Int64
	err    error
	wg     *sync.WaitGroup
}

func (m *mockCommand) Execute(event *gateway.RawEvent) error {
	defer m.wg.Done()
	m.called.Add(1)
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
	mockCmd := &mockCommand{
		wg: &sync.WaitGroup{},
	}
	bot := &g4d.Bot{
		Gateway: &gw.Receiver{
			Queue: make(chan *gateway.RawEvent, 10),
		},
		CommandBuffer: []g4d.CommandTemplate{
			{Trigger: "MESSAGE_CREATE", Name: "test", Action: mockCmd.Execute},
		},
		Logger: zap.NewNop(),
	}

	go bot.StaticEventProcessor(5)

	for x := 0; x < 100; x++ {
		mockCmd.wg.Add(1)
		go func() {
			bot.Gateway.Queue <- &gateway.RawEvent{Type: "MESSAGE_CREATE"}
		}()
	}
	mockCmd.wg.Wait()
	if mockCmd.called.Load() != 100 {
		t.Errorf("Expected %d commands, got 100", mockCmd.called.Load())
	}
}

// Тест динамического процессора
func TestBot_DynamicEventProcessor(t *testing.T) {
	mockCmd := &mockCommand{
		wg: &sync.WaitGroup{},
	}
	bot := &g4d.Bot{
		Gateway: &gw.Receiver{
			Queue: make(chan *gateway.RawEvent, 10),
		},
		CommandBuffer: []g4d.CommandTemplate{
			{Trigger: "MESSAGE_CREATE", Name: "test", Action: mockCmd.Execute},
		},
		Logger: zap.NewNop(),
	}

	go bot.DynamicEventProcessor(5)
	for x := 0; x < 100; x++ {
		mockCmd.wg.Add(1)
		go func() {
			bot.Gateway.Queue <- &gateway.RawEvent{Type: "MESSAGE_CREATE"}
		}()
	}
	mockCmd.wg.Wait()
	if mockCmd.called.Load() != 100 {
		t.Errorf("Expected %d commands, got 100", mockCmd.called.Load())
	}
}
func BenchmarkStaticEventProcessor(b *testing.B) {
	mockCmd := &mockCommand{
		wg: &sync.WaitGroup{},
	}

	bot := &g4d.Bot{
		Gateway: &gw.Receiver{
			Queue: make(chan *gateway.RawEvent, 1000),
		},
		CommandBuffer: []g4d.CommandTemplate{
			{Trigger: "MESSAGE_CREATE", Name: "test", Action: mockCmd.Execute},
		},
		Logger: zap.NewNop(),
	}

	go bot.StaticEventProcessor(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mockCmd.wg.Add(1)
		bot.Gateway.Queue <- &gateway.RawEvent{Type: "MESSAGE_CREATE"}
	}
	mockCmd.wg.Wait()
}
