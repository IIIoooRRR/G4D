## Interaction Message and TextMessage
```go
//TextMessage
func Hello(event *Connect.RawEvent) error {
	log.Printf("DEBUG DATA: %+v", event)
	d := Parse.ToMessageCreate(*event)
	if d.Content != "!hello" { return nil }
	msg := Parse.NewMessage().AddReferencedMessage(&data).AddContent("hihi")
    err := functions.SendMessage(data.ChannelID, msg)
    if err != nil { return err }
	return nil
}
// Interaction
func Slash(event *Connect.RawEvent) error {
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
	err := Functions.SendInteractionMessage(d)
	if err != nil {
		log.Println(err)
	}
	return nil
}

```