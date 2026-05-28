package g4d_test

import (
	"context"
	"log"

	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/g4d"
	"github.com/IIIoooRRR/G4D/gateway"
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/_struct"
	"github.com/IIIoooRRR/G4D/model/parse"
)

// an example of initializing a bot and assigning basic debug_commands
func ExampleBot() {
	token := "your-token"
	bot := &g4d.Bot{
		Token: token,
		Gateway: &gateway.Receiver{
			QueueSize: 20,
			Intents:   33280,
		},
	}
	// creating debug_commands. Assign the type of event and the function that will be executed when initializing it.
	// for slash debug_commands, see the following example
	bot.AddCommands([]g4d.CommandTemplate{
		{
			Trigger: _const.EventMessageCreate,
			Action:  Hello,
		},
	})
	//for ease of development, implement the bot prefix
	bot.Prefix = "!"
	// Implement an event processor. It can be dynamic or static. Read in the relevant section
	bot.StaticEventProcessor() // Choose wisely: performance or adaptability.
	err := bot.Gateway.CreateBot(context.Background(), &bot.Token)
	if err != nil {
		log.Println(err)
	}
	// Output:
}

// The function that will be called at the Message Create event
// It has strict Bot and RawEvent fields, and it should also return an error for logging by the processor.
func Hello(event *gateway.RawEvent) error {
	data := parse.ToMessageCreate(*event)
	if data.Content == g4d.CurrentBot().Prefix+"hello" {
		msg := _struct.NewMessage().AddContent("Ping <@" + data.Author.Id + ">")
		api.SendMessage(data.ChannelID, msg)
	}
	return nil
}
