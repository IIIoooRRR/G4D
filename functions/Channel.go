package functions

import (
	"encoding/json"
	"fmt"

	"github.com/IIIoooRRR/G4D/JSON/Parse"
	"github.com/IIIoooRRR/G4D/JSON/Type"
)

func CreateChannel(guildId Type.GuildId, channel *Parse.Channel) error {
	jsonBody, err := json.Marshal(channel)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/channels/%s", guildId)
	_, err = doDiscordRequest("POST", endpoint, jsonBody)
	return err
}

func DeleteChannel(channelId Type.ChannelId, channel *Parse.Channel) error {
	jsonBody, err := json.Marshal(channel)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/channels/%s", channelId)
	_, err = doDiscordRequest("DELETE", endpoint, jsonBody)
	return err
}

func ChangeChannels(channelId Type.ChannelId, channel *Parse.Channel) error {
	jsonBody, err := json.Marshal(channel)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("/channels/%s", channelId)
	_, err = doDiscordRequest("PATCH", endpoint, jsonBody)
	return err
}

func GetChannel(channelId Type.ChannelId) (*Parse.Channel, error) {
	endpoint := fmt.Sprintf("/channels/%s", channelId)
	abstract, err := doDiscordRequest("GET", endpoint, []byte{})
	if err != nil {
		return nil, err
	}
	return Parse.ToChannel(abstract)
}
