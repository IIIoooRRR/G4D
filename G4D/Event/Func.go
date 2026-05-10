package Event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IIIoooRRR/G4D/G4D"
)

func SendInteractionMessage(msg *InteractionResponse, event *Interaction) error {
	url := fmt.Sprintf("https://discord.com/api/v10/interactions/%s/%s/callback", event.ID, event.Token)
	jsonData, _ := json.Marshal(msg)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "G4D "+G4D.CurrentBot().Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Println("[SEND-MESSAGE] Discord API ERROR " + resp.Status)
	}
	return nil
}

func SendMessage(ToChannel string, msg *Message) error {
	var url = "https://discord.com/api/v10/channels/"
	url = url + ToChannel + "/messages"
	body := msg
	jsonBody, err := json.Marshal(body) //delaem жсон из message
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	// добавляем поля
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "G4D "+G4D.CurrentBot().Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // ждем пока закроется
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Println("[SEND-MESSAGE] Discord API ERROR " + resp.Status)
	}

	return nil
}
func BanUser(guildId string, userId string) error {
	var url = fmt.Sprintf("https://discord.com/api/v10/guilds/%s/bans/%s", guildId, userId)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "G4D "+G4D.CurrentBot().Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
func MuteUser(guildId string, userId string, dur time.Duration) error {
	until := time.Now().Add(dur).Format(time.RFC3339)
	var url = fmt.Sprintf("https://discord.com/api/v10/guilds/%s/members/%s", guildId, userId)
	body := Mute{
		Time: &until,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "G4D "+G4D.CurrentBot().Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
