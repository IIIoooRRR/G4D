package G4D_test

import (
	"github.com/IIIoooRRR/G4D/Connect"
	"github.com/IIIoooRRR/G4D/G4D"
	"github.com/IIIoooRRR/G4D/G4D/Event"
	"github.com/IIIoooRRR/G4D/JSON/Parse"
)

func ExampleBot_AddCommand() {

	bot := &G4D.Bot{
		Token: "token",
		Gateway: &Connect.Receiver{
			Intents:   33280,
			QueueSize: 40,
		},
		Prefix: "!",
	}
	//To implement basic(text) commands, use the addCommand(s) method

	bot.AddCommands([]G4D.CommandTemplate{
		{
			Trigger: Event.EventMessageCreate,
			Action:  Execute, //use any function
		}, //. It should accept G4D and RawEvent, return error.
	})
	// to implement slash commands, use the AddSlashCommand(s) method
	bot.AddSlashCommands([]G4D.SlashCommandTemplate{
		{
			Form: G4D.SlashCreateCommand{ //use struct SlashCreateCommand
				Name:        "",
				Description: "", //implement the fields for the command
				Type:        0,
				Options:     nil,
			},
			CommandTemplate: G4D.CommandTemplate{
				Trigger: Event.EventInteractionCreate, //specify the type of event to initialize and the function
				Action:  Slash,
			},
		},
	})
	//the rest of the bot implementation
}

func Execute(event *Connect.RawEvent) error {
	d := Parse.ToMessageCreate(*event) //turn the resulting RawEvent into the structure you need.
	if d.Content == "!hello" {
		msg := Event.Message{
			Content: "hello world",
		}
		Event.SendMessage(d.ChannelID, &msg)
	}
	return nil
}
func Slash(event *Connect.RawEvent) error {
	data := Parse.ToInteraction(*event) // turn the resulting RawEvent into the structure you need.
	msg := Event.InteractionResponse{
		Type: 0,
		Data: Event.InteractionResponseData{
			Content: "hello world",
			Flags:   0,
			Embeds:  nil,
		},
	}
	Event.SendInteractionMessage(&msg, &data)
	return nil
}
