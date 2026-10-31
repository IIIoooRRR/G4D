package api_test

import (
	"context"
	"time"

	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/ctx"
	"github.com/IIIoooRRR/G4D/model/schema"
)

// ExampleSendMessageWithLimit shows how to send a message through a
// ctx.Context that was created by the bot and passed into a command.
func ExampleDiscordClient_SendMessageWithLimit() {
	// ctx.Context is normally provided by the bot when it invokes a command.
	var c ctx.Context

	msg := schema.NewMessage().Content("hello")
	if err := c.SendMessageWithLimit("channel_id", msg); err != nil {
		// handle error
		return //err
	}
}

// ExampleClient_DoDiscordRequest shows the low-level request primitive.
func ExampleDiscordClient_DoDiscordRequest() {
	client := api.NewClient(ptr("TOKEN"), 10*time.Second)

	body, err := client.DoDiscordRequest(
		_const.Get,
		"/guilds/123/members/456",
		nil,
	)
	if err != nil {
		// handle error
		return //err
	}
	_ = body
}

// ExampleClient_DoDiscordRequest shows the low-level request primitive.
func ExampleDiscordClient_DoDiscordLimitRequest() {
	client := api.NewClient(ptr("TOKEN"), 10*time.Second)

	body, err := client.DoDiscordLimitRequest(
		context.Background(),
		_const.Get,
		"/guilds/123/members/456",
		nil,
	)
	if err != nil {
		// handle error
		return //err
	}
	_ = body
}

func ptr[T any](v T) *T { return &v }
