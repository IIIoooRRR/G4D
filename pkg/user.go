package pkg

import (
	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/model/_const"
)

func GetAvatarURL(userID _const.UserId, avatarHash string) string {
	return api.GetURI("https://cdn.discordapp.com/avatars/", string(userID), "/", avatarHash, ".png")
}
