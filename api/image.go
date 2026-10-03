package api

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
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

	_, err = c.DoDiscordRequest(_const.Post, url, body.Bytes())
	if err != nil {
		return err
	}

	return nil
}
