package api

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/parse"
	"github.com/IIIoooRRR/G4D/model/schema"
	"go.uber.org/zap"
)

func (c *DiscordClient) SendImage(toChannel _const.ChannelId, msg schema.SendMessage, path string) error {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			c.logger.Warn("failed to close file", zap.Error(err))
		}
	}()
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	if stat.Size() == 0 {
		return errors.New("file is empty")
	}
	url := GetURI("https://discord.com/api/v10/channels/", string(toChannel), "/messages")
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	payloadJSON, err := parse.Marshal(msg)
	if err != nil {
		return err
	}

	if err := writer.WriteField("payload_json", string(payloadJSON)); err != nil {
		return err
	}

	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return err
	}

	if _, err := io.Copy(part, file); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bot "+*c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			c.logger.Warn("failed to close response body", zap.Error(err))
		}
	}()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		if len(respBody) > 1024 {
			respBody = respBody[:1024]
		}
		return fmt.Errorf("discord returned %d: %s", resp.StatusCode, string(respBody))
	}

	return nil
}
