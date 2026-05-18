package G4D_test

import (
	"github.com/IIIoooRRR/G4D/G4D"
	"github.com/IIIoooRRR/G4D/JSON/Parse"
	"github.com/IIIoooRRR/G4D/JSON/Type"
	"github.com/IIIoooRRR/G4D/connect"
	"github.com/IIIoooRRR/G4D/functions"
)

func ExampleBot_AddCommand() {

	bot := &G4D.Bot{
		Token: "token",
		Gateway: &connect.Receiver{
			Intents:   33280,
			QueueSize: 40,
		},
		Prefix: "!",
	}
	//To implement basic(text) commands, use the addCommand(s) method

	bot.AddCommands([]G4D.CommandTemplate{
		{
			Trigger: Type.EventMessageCreate,
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
				Trigger: Type.EventInteractionCreate, //specify the type of event to initialize and the function
				Action:  Slash,
			},
		},
	})
	//the rest of the bot implementation
}

func Execute(event *connect.RawEvent) error {
	d := Parse.ToMessageCreate(*event) //turn the resulting RawEvent into the structure you need.
	if d.Content == "!hello" {
		msg := Parse.Message{
			Content: "hello world",
		}
		functions.SendMessage(d.ChannelID, &msg)
	}
	return nil
}
func Slash(event *connect.RawEvent) error {
	data := Parse.ToInteraction(*event) // turn the resulting RawEvent into the structure you need.
	msg := Parse.InteractionResponse{
		Type: 0,
		Data: Parse.InteractionResponseData{
			Content: "hello world",
			Flags:   0,
			Embeds:  nil,
		},
	}
	functions.SendInteractionMessage(&msg, &data)
	return nil
}
