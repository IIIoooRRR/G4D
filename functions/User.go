package functions

import (
	"fmt"

	"github.com/IIIoooRRR/G4D/JSON/Type"
)

func GetAvatarURL(userID Type.UserId, avatarHash string) string {
	if avatarHash == "" {
		return ""
	}
	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", userID, avatarHash)
}
