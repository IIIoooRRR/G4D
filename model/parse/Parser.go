package parse

import (
	"encoding/json"
	"errors"

	"github.com/IIIoooRRR/G4D/gateway"
	"github.com/IIIoooRRR/G4D/model/_struct"
	"github.com/IIIoooRRR/G4D/model/dependencies/ui"
)

/* Basic */
func ToMessageCreate(event gateway.RawEvent) *_struct.GetMessage {
	var d *_struct.GetMessage
	json.Unmarshal(event.Data, &d)
	return d
}

func ToMessageDelete(event gateway.RawEvent) *_struct.MessageDelete {
	var d *_struct.MessageDelete
	json.Unmarshal(event.Data, &d)
	return d
}
func ToMessageEdited(event gateway.RawEvent) *_struct.MessageEdit {
	var d *_struct.MessageEdit
	json.Unmarshal(event.Data, &d)
	return d
}

func ToInteraction(event gateway.RawEvent) *_struct.Interaction {
	var d *_struct.Interaction
	json.Unmarshal(event.Data, &d)
	return d
}

/* Reactions */
func ToReactionsAdd(event gateway.RawEvent) *_struct.MessageReactionAdd {
	var d *_struct.MessageReactionAdd
	json.Unmarshal(event.Data, &d)
	return d
}
func ToReactionRemove(event gateway.RawEvent) *_struct.MessageReactionRemove {
	var d *_struct.MessageReactionRemove
	json.Unmarshal(event.Data, &d)
	return d
}

/* Guild */
func ToGuildMemberAdd(event gateway.RawEvent) *_struct.GuildMemberAdd {
	var d *_struct.GuildMemberAdd
	json.Unmarshal(event.Data, &d)
	return d
}
func ToGuildMemberRemove(event gateway.RawEvent) *_struct.GuildMemberRemove {
	var d *_struct.GuildMemberRemove
	json.Unmarshal(event.Data, &d)
	return d
}

/* Channel */
func ToChannel(body []byte) (*_struct.Channel, error) {
	var channel *_struct.Channel
	err := json.Unmarshal(body, channel)
	if err != nil {
		return nil, err
	}
	return channel, nil
}

/* Components */
func GetModalValue(interaction *_struct.Interaction, customID string) (string, error) {
	for _, row := range interaction.Data.Components {
		for _, comp := range row.Components {
			if input, ok := comp.(*ui.TextInput); ok {
				if input.CustomID == customID {
					return input.Value, nil
				}
			}
		}
	}
	return "", errors.New("input not found")
}
