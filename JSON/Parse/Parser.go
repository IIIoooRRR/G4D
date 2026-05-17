package Parse

import (
	"encoding/json"
	"log"

	"github.com/IIIoooRRR/G4D/Connect"
)

func ToMessageCreate(event Connect.RawEvent) Message {
	var d Message
	err := json.Unmarshal(event.Data, &d)
	if err != nil {
		log.Println(err)
	}
	return d
}

func ToMessageDelete(event Connect.RawEvent) MessageDelete {
	var d MessageDelete
	err := json.Unmarshal(event.Data, &d)
	if err != nil {
		log.Println(err)
	}
	return d
}
func ToMessageEdited(event Connect.RawEvent) MessageEdit {
	var d MessageEdit
	err := json.Unmarshal(event.Data, &d)
	if err != nil {
		log.Println(err)
	}
	return d
}

func ToInteraction(event Connect.RawEvent) Interaction {
	var d Interaction
	json.Unmarshal(event.Data, &d)
	return d
}

func ToGuildMemberAdd(event Connect.RawEvent) GuildMemberAdd {
	var d GuildMemberAdd
	json.Unmarshal(event.Data, &d)
	return d
}
func ToGuildMemberRemove(event Connect.RawEvent) GuildMemberRemove {
	var d GuildMemberRemove
	json.Unmarshal(event.Data, &d)
	return d
}

func ToChannel(body []byte) (Channel, error) {
	var channel Channel
	err := json.Unmarshal(body, &channel)
	if err != nil {
		return Channel{}, err
	}
	return channel, nil
}
func ToReactionsAdd(event Connect.RawEvent) MessageReactionAdd {
	var d MessageReactionAdd
	json.Unmarshal(event.Data, &d)
	return d
}
