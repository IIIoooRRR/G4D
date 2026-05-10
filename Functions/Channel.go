package Functions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/IIIoooRRR/G4D/G4D"
	"github.com/IIIoooRRR/G4D/JSON/Parse"
)

func CreateChannel(guildId string, channel *Parse.Channel) error {
	var url = fmt.Sprintf("https://discord.com/api/v10/guilds/%s/channels", guildId)
	jsonBody, err := json.Marshal(channel)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+G4D.CurrentBot().Token)
	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return errors.New("[DISCORD] Failed to create channel")
	}
	return nil
}

func DeleteChannel(channelId string, channel *Parse.Channel) error {
	var url = fmt.Sprintf("https://discord.com/api/v10/channels/%s", channelId)
	jsonBody, err := json.Marshal(channel)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+G4D.CurrentBot().Token)
	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return errors.New("[DISCORD] Failed to delete channel")
	}
	return nil
}

func ChangeChannels(channelID string, channel *Parse.Channel) error {
	var url = fmt.Sprintf("https://discord.com/api/v10/channels/%s", channelID)
	jsonBody, err := json.Marshal(channel)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+G4D.CurrentBot().Token)
	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return errors.New("[DISCORD] Failed to update channels")
	}
	return nil
}

func GetChannel(channelId string) (*Parse.Channel, error) {
	if G4D.CurrentBot().Cache != nil {
		cache, err := G4D.CurrentBot().Cache.GetChannel(channelId)
		if err == nil {
			return &cache, err
		}

	}
	var url = fmt.Sprintf("https://discord.com/api/v10/channels/%s", channelId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &Parse.Channel{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+G4D.CurrentBot().Token)
	client := http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, errors.New("[DISCORD] Failed to get channel")
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var channel Parse.Channel
	err = json.Unmarshal(bodyBytes, &channel)
	if err != nil {
		return &Parse.Channel{}, err
	}

	return &channel, nil
}
