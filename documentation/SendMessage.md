## Interaction Message and TextMessage
```go
//TextMessage
func Hello(event *Connect.RawEvent) error {
	log.Printf("DEBUG DATA: %+v", event)
	d := parse.Event[shema.GetMessage](event)
	if d.Content != "!hello" { return nil }
	msg := Parse.NewMessage().AddReferencedMessage(&data).AddContent("hihi")
    return api.SendMessage(data.ChannelID, msg)
}
// Interaction
func Slash(event *Connect.RawEvent) error {
	log.Printf("DEBUG DATA: %+v", event)
	d := parse.Event[shema.Interaction](event)
	response := Event.InteractionResponse{
		Type: 4,
		Data: Event.InteractionResponseData{
			Content: "Hello World!",
			Flags:   0,
			Embeds:  nil,
		},
	}
	return api.SendInteractionMessage(d)
}

```