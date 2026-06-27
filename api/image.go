package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/IIIoooRRR/G4D/g4d"
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/shema"
)

func SendImage(toChannel _const.ChannelId, msg shema.SendMessage, path string) error {
	var url = fmt.Sprintf("https://discord.com/api/v10/channels/%s/messages", toChannel)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)

	payload := map[string]interface{}{
		"content": msg.Content,
		"embeds":  msg.Embeds,
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = writer.WriteField("payload_json", string(payloadJSON))
	if err != nil {
		return err
	}
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bot "+g4d.CurrentBot().Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("discord error %d: %s", resp.StatusCode, string(respBody))
	}

	return nil
}
