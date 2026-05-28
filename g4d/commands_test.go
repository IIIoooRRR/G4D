package g4d_test

import (
	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/g4d"
	"github.com/IIIoooRRR/G4D/gateway"
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/_struct"
	"github.com/IIIoooRRR/G4D/model/parse"
)

func ExampleBot_AddCommand() {

	bot := &g4d.Bot{
		Token: "token",
		Gateway: &gateway.Receiver{
			Intents:   33280,
			QueueSize: 40,
		},
		Prefix: "!",
	}
	//To implement basic(text) debug_commands, use the addCommand(s) method

	bot.AddCommands([]g4d.CommandTemplate{
		{
			Trigger: _const.EventMessageCreate,
			Action:  Execute, //use any function
		}, //. It should accept g4d and RawEvent, return error.
	})
	// to implement slash debug_commands, use the AddSlashCommand(s) method
	bot.AddSlashCommands([]g4d.SlashCommandTemplate{
		{
			Form: g4d.SlashCreateCommand{ //use struct SlashCreateCommand
				Name:        "",
				Description: "", //implement the fields for the command
				Type:        0,
				Options:     nil,
			},
			CommandTemplate: g4d.CommandTemplate{
				Trigger: _const.EventInteractionCreate, //specify the type of event to initialize and the function
				Action:  Slash,
			},
		},
	})
	//the rest of the bot implementation
}

func Execute(event *gateway.RawEvent) error {
	d := parse.ToMessageCreate(*event) //turn the resulting RawEvent into the structure you need.
	if d.Content == "!hello" {
		msg := _struct.SendMessage{
			Content: "hello world",
		}
		api.SendMessage(d.ChannelID, &msg)
	}
	return nil
}
func Slash(event *gateway.RawEvent) error {
	data := parse.ToInteraction(*event) // turn the resulting RawEvent into the structure you need.
	msg := _struct.InteractionResponse{
		Type: 0,
		Data: _struct.InteractionResponseData{
			Content: "hello world",
			Flags:   0,
			Embeds:  nil,
		},
	}
	api.SendInteractionMessage(data, msg)
	return nil
}
