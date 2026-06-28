package pkg

import (
	"fmt"

	"github.com/IIIoooRRR/G4D/model/_const"
)

func GetAvatarURL(userID _const.UserId, avatarHash string) string {
	if avatarHash == "" {
		return ""
	}
	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", userID, avatarHash)
}
