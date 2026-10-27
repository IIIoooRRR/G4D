package ctx

import (
	"strings"

	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/model/_const"
	"go.uber.org/zap"
)

type Context struct {
	Prefix string
	Logger *zap.Logger
	*api.DiscordClient
}

// match maked without allocate(cpu-bounding)

func (ctx *Context) IsCommand(s1, s2 string) bool {
	// s1 == prefix + s2
	if len(s1) != len(ctx.Prefix)+len(s2) {
		return false
	}
	return s1[:len(ctx.Prefix)] == ctx.Prefix && s1[len(ctx.Prefix):] == s2
}

func (ctx *Context) MatchCommand(s1, s2 string) bool {
	if len(s1) < len(ctx.Prefix)+len(s2) {
		return false
	}
	if s1[:len(ctx.Prefix)] != ctx.Prefix {
		return false
	}
	return strings.HasPrefix(s1[len(ctx.Prefix):], s2)
}

func GetAvatarURL(userID _const.UserId, avatarHash string) string {
	return api.GetURI("https://cdn.discordapp.com/avatars/", string(userID), "/", avatarHash, ".png")
}
