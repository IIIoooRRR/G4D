package api

import (
	"fmt"
	"net/url"

	"github.com/IIIoooRRR/G4D/model/_const"
)

func AddReaction(channelId _const.ChannelId, messageId _const.MessageId, reactionId string) error {
	encodedReaction := url.QueryEscape(reactionId)
	uri := fmt.Sprintf("/channels/%s/messages/%s/reactions/%s/@me", channelId, messageId, encodedReaction)
	_, err := DoDiscordRequest("PUT", uri, []byte{})
	return err
}
func DeleteReaction(channelId _const.ChannelId, messageId _const.MessageId, reactionId string) error {
	encodedReaction := url.QueryEscape(reactionId)
	uri := fmt.Sprintf("/channels/%s/messages/%s/reactions/%s/@me", channelId, messageId, encodedReaction)
	_, err := DoDiscordRequest("DELETE", uri, []byte{})
	return err
}
func DeleteAllReactions(channelId _const.ChannelId, messageId _const.MessageId) error {
	uri := fmt.Sprintf("/channels/%s/messages/%s/reactions", channelId, messageId)
	_, err := DoDiscordRequest("DELETE", uri, []byte{})
	return err
}
func DeleteAllReactionsForEmoji(channelId _const.ChannelId, messageId _const.MessageId, reactionId string) error {
	encodedReaction := url.QueryEscape(reactionId)
	uri := fmt.Sprintf("/channels/%s/messages/%s/reactions/%s", channelId, messageId, encodedReaction)
	_, err := DoDiscordRequest("DELETE", uri, []byte{})
	return err
}
