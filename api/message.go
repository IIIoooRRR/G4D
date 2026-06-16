package api

import (
	"fmt"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/parse"
	"github.com/IIIoooRRR/G4D/model/shema"
)

func SendInteractionMessage(event *shema.Interaction, msg shema.InteractionResponse) error {
	uri := fmt.Sprintf("/interactions/%v/%s/callback", event.ID, event.Token)
	jsonData, err := parse.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = DoDiscordRequest("POST", uri, jsonData)
	return err
}

func SendMessage(ToChannel _const.ChannelId, msg *shema.SendMessage) error {
	uri := fmt.Sprintf("/channels/%s/messages", ToChannel)
	body := msg
	jsonBody, err := parse.Marshal(body) //delaem жсон из message
	if err != nil {
		return err
	}
	_, err = DoDiscordRequest("POST", uri, jsonBody)
	return err
}
func EditMessage(ToChannel _const.ChannelId, msgId _const.MessageId, msg *shema.MessageEdit) error {
	uri := fmt.Sprintf("/channels/%s/messages/%s", ToChannel, msgId)
	jsonBody, err := parse.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = DoDiscordRequest("PATCH", uri, jsonBody)
	return err
}
func DeleteMessage(ToChannel _const.ChannelId, msgId _const.MessageId) error {
	uri := fmt.Sprintf("/channels/%s/messages/%s", ToChannel, msgId)
	_, err := DoDiscordRequest("DELETE", uri, []byte{})
	return err
}
