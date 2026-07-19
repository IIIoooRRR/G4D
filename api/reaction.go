package api

import (
	"fmt"
	"net/url"

	"github.com/IIIoooRRR/G4D/model/_const"
)

func (c *DiscordClient) AddReaction(channelId _const.ChannelId, messageId _const.MessageId, reactionId string) error {
	encodedReaction := url.QueryEscape(reactionId)
	uri := fmt.Sprintf("/channels/%s/messages/%s/reactions/%s/@me", channelId, messageId, encodedReaction)
	_, err := c.DoDiscordRequest("PUT", uri, []byte{})
	return err
}
func (c *DiscordClient) DeleteReaction(channelId _const.ChannelId, messageId _const.MessageId, reactionId string) error {
	encodedReaction := url.QueryEscape(reactionId)
	uri := fmt.Sprintf("/channels/%s/messages/%s/reactions/%s/@me", channelId, messageId, encodedReaction)
	_, err := c.DoDiscordRequest("DELETE", uri, []byte{})
	return err
}
func (c *DiscordClient) DeleteAllReactions(channelId _const.ChannelId, messageId _const.MessageId) error {
	uri := fmt.Sprintf("/channels/%s/messages/%s/reactions", channelId, messageId)
	_, err := c.DoDiscordRequest("DELETE", uri, []byte{})
	return err
}
func (c *DiscordClient) DeleteAllReactionsForEmoji(channelId _const.ChannelId, messageId _const.MessageId, reactionId string) error {
	encodedReaction := url.QueryEscape(reactionId)
	uri := fmt.Sprintf("/channels/%s/messages/%s/reactions/%s", channelId, messageId, encodedReaction)
	_, err := c.DoDiscordRequest("DELETE", uri, []byte{})
	return err
}
