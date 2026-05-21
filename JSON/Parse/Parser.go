package Parse

import (
	"encoding/json"
	"log"

	"github.com/IIIoooRRR/G4D/connect"
)

/* Basic */
func ToMessageCreate(event connect.RawEvent) *GetMessage {
	var d *GetMessage
	err := json.Unmarshal(event.Data, &d)
	if err != nil {
		log.Println(err)
	}
	return d
}

func ToMessageDelete(event connect.RawEvent) *MessageDelete {
	var d *MessageDelete
	err := json.Unmarshal(event.Data, &d)
	if err != nil {
		log.Println(err)
	}
	return d
}
func ToMessageEdited(event connect.RawEvent) *MessageEdit {
	var d *MessageEdit
	err := json.Unmarshal(event.Data, &d)
	if err != nil {
		log.Println(err)
	}
	return d
}

func ToInteraction(event connect.RawEvent) *Interaction {
	var d *Interaction
	json.Unmarshal(event.Data, &d)
	return d
}

/* Reactions */
func ToReactionsAdd(event connect.RawEvent) *MessageReactionAdd {
	var d *MessageReactionAdd
	json.Unmarshal(event.Data, &d)
	return d
}
func ToReactionRemove(event connect.RawEvent) *MessageReactionRemove {
	var d *MessageReactionRemove
	json.Unmarshal(event.Data, &d)
	return d
}

/* Guild */
func ToGuildMemberAdd(event connect.RawEvent) *GuildMemberAdd {
	var d *GuildMemberAdd
	json.Unmarshal(event.Data, &d)
	return d
}
func ToGuildMemberRemove(event connect.RawEvent) *GuildMemberRemove {
	var d *GuildMemberRemove
	json.Unmarshal(event.Data, &d)
	return d
}

/* Channel */
func ToChannel(body []byte) (*Channel, error) {
	var channel *Channel
	err := json.Unmarshal(body, channel)
	if err != nil {
		return nil, err
	}
	return channel, nil
}
