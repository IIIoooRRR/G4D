package Functions

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/IIIoooRRR/G4D/G4D"
)

func AddReaction(channelId, messageId, reactionId string) error {
	encodedReaction := url.QueryEscape(reactionId)
	url := fmt.Sprintf("https://discord.com/api/v10/channels/%s/messages/%s/reactions/%s/@me", channelId, messageId, encodedReaction)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+G4D.CurrentBot().Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
func DeleteReaction(channelId, messageId, reactionId string) error {
	encodedReaction := url.QueryEscape(reactionId)
	url := fmt.Sprintf("https://discord.com/api/v10/channels/%s/messages/%s/reactions/%s/@me", channelId, messageId, encodedReaction)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+G4D.CurrentBot().Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
func DeleteAllReactions(channelId, messageId string) error {
	url := fmt.Sprintf("https://discord.com/api/v10/channels/%s/messages/%s/reactions", channelId, messageId)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+G4D.CurrentBot().Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
func DeleteAllReactionsForEmoji(channelId, messageId, reactionId string) error {
	encodedReaction := url.QueryEscape(reactionId)
	url := fmt.Sprintf("https://discord.com/api/v10/channels/%s/messages/%s/reactions/%s", channelId, messageId, encodedReaction)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+G4D.CurrentBot().Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
