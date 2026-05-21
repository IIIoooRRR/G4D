package functions

import (
	"encoding/json"
	"fmt"

	"github.com/IIIoooRRR/G4D/JSON/Parse"
	"github.com/IIIoooRRR/G4D/JSON/Type"
)

func SendInteractionMessage(msg *Parse.InteractionResponse, event *Parse.Interaction) error {
	uri := fmt.Sprintf("/interactions/%s/%s/callback", event.ID, event.Token)
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = doDiscordRequest("POST", uri, jsonData)
	return err
}

func SendMessage(ToChannel Type.ChannelId, msg *Parse.SendMessage) error {
	uri := fmt.Sprintf("/channels/%s/messages", ToChannel)
	body := msg
	jsonBody, err := json.Marshal(body) //delaem жсон из message
	if err != nil {
		return err
	}
	_, err = doDiscordRequest("POST", uri, jsonBody)
	return err
}
func EditMessage(ToChannel Type.ChannelId, msgId Type.MessageId, msg *Parse.MessageEdit) error {
	uri := fmt.Sprintf("/channels/%s/messages/%s", ToChannel, msgId)
	jsonBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = doDiscordRequest("PATCH", uri, jsonBody)
	return err
}
func DeleteMessage(ToChannel Type.ChannelId, msgId Type.MessageId) error {
	uri := fmt.Sprintf("/channels/%s/messages/%s", ToChannel, msgId)
	_, err := doDiscordRequest("DELETE", uri, []byte{})
	return err
}
