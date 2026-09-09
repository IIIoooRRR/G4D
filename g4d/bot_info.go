package g4d

import (
	"go.uber.org/zap"
)

func (b *Bot) getBotInfo() error {
	info, err := b.Client.GetBotInfo()
	if err != nil {
		return err
	}

	b.appId = string(info.Id)
	b.Logger.Info("bot info retrieved", zap.String("bot_id:", b.appId))
	return nil
}
