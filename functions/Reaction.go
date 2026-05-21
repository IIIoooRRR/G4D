package functions

import (
	"fmt"
	"net/url"

	"github.com/IIIoooRRR/G4D/JSON/Type"
)

func AddReaction(channelId Type.ChannelId, messageId Type.MessageId, reactionId string) error {
	encodedReaction := url.QueryEscape(reactionId)
	uri := fmt.Sprintf("/channels/%s/messages/%s/reactions/%s/@me", channelId, messageId, encodedReaction)
	_, err := doDiscordRequest("PATCH", uri, []byte{})
	return err
}
func DeleteReaction(channelId Type.ChannelId, messageId Type.MessageId, reactionId string) error {
	encodedReaction := url.QueryEscape(reactionId)
	uri := fmt.Sprintf("/channels/%s/messages/%s/reactions/%s/@me", channelId, messageId, encodedReaction)
	_, err := doDiscordRequest("PATCH", uri, []byte{})
	return err
}
func DeleteAllReactions(channelId Type.ChannelId, messageId Type.MessageId) error {
	uri := fmt.Sprintf("/channels/%s/messages/%s/reactions", channelId, messageId)
	_, err := doDiscordRequest("DELETE", uri, []byte{})
	return err
}
func DeleteAllReactionsForEmoji(channelId Type.ChannelId, messageId Type.MessageId, reactionId string) error {
	encodedReaction := url.QueryEscape(reactionId)
	uri := fmt.Sprintf("/channels/%s/messages/%s/reactions/%s", channelId, messageId, encodedReaction)
	_, err := doDiscordRequest("DELETE", uri, []byte{})
	return err
}
