package G4D_test

import (
	"context"
	"log"

	"github.com/IIIoooRRR/G4D/Connect"
	"github.com/IIIoooRRR/G4D/G4D"
	"github.com/IIIoooRRR/G4D/G4D/Event"
	"github.com/IIIoooRRR/G4D/JSON/Parse"
)

// an example of initializing a bot and assigning basic commands
func ExampleBot() {
	token := "your-token"
	bot := &G4D.Bot{
		Token: token,
		Gateway: &Connect.Receiver{
			QueueSize: 20,
			Intents:   33280,
		},
	}
	// creating commands. Assign the type of event and the function that will be executed when initializing it.
	// for slash commands, see the following example
	bot.AddCommands([]G4D.CommandTemplate{
		{
			Trigger: Event.EventMessageCreate,
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
func Hello(event *Connect.RawEvent) error {
	data := Parse.ToMessageCreate(*event)
	if data.Content == G4D.CurrentBot().Prefix+"hello" {
		msg := Event.Message{
			ChannelID: data.ChannelID,
			GuildID:   data.GuildID,
			Content:   "Ping <@" + data.Author.Id + ">",
			Flags:     0,
			Type:      0,
		}
		Event.SendMessage(data.ChannelID, &msg)
	}
	return nil
}
