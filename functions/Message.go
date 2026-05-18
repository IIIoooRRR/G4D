package Functions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/IIIoooRRR/G4D/G4D"
	"github.com/IIIoooRRR/G4D/JSON/Parse"
)

func SendInteractionMessage(msg *Parse.InteractionResponse, event *Parse.Interaction) error {
	url := fmt.Sprintf("https://discord.com/api/v10/interactions/%s/%s/callback", event.ID, event.Token)
	jsonData, _ := json.Marshal(msg)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+G4D.CurrentBot().Token)
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

func SendMessage(ToChannel string, msg *Parse.Message) error {
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
	req.Header.Set("Authorization", "Bot "+G4D.CurrentBot().Token)
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
func EditMessage(ToChannel, msgId string, msg *Parse.MessageEdit) error {
	url := fmt.Sprintf("https://discord.com/api/v10/channels/%s/messages/%s", ToChannel, msgId)
	jsonBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+G4D.CurrentBot().Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
func DeleteMessage(ToChannel string, msgId string) error {
	url := fmt.Sprintf("https://discord.com/api/v10/channels/%s/messages/%s", ToChannel, msgId)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+G4D.CurrentBot().Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
