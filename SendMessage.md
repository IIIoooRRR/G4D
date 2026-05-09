## Interaction Message and TextMessage
```go
//TextMessage
func Hello(b *Bot.Bot, event *Connect.RawEvent) error {
	log.Printf("DEBUG DATA: %+v", event)
	d := Parse.ToMessageCreate(*event)
	if d.Content == "!hello" {
		msg := Event.Message{

			Content: "hello world",
		}
		msg.SendMessage(d.ChannelID, b)
	}
	return nil
}
// Interaction
func Slash(b *Bot.Bot, event *Connect.RawEvent) error {
	log.Printf("DEBUG DATA: %+v", event)
	d := Parse.ToInteraction(*event)
	response := Event.InteractionResponse{
		Type: 4,
		Data: Event.InteractionResponseData{
			Content: "Hello World!",
			Flags:   0,
			Embeds:  nil,
		},
	}
	err := response.SendInteractionMessage(d, b)
	if err != nil {
		log.Println(err)
	}
	return nil
}

```