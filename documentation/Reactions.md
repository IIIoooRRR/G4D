```go
func Hello(event *connect.RawEvent) error {
	// Attention! Here are examples that can cause panic if you do not specify a message (preferably an id) in options 2-4 with the second word 
	data := parse.Event[parse.GetMessage](event)
	message := strings.Split(data.Content, " ")
	var MessageId _const.MessageId
 if len(message) > 2 {
	  MessageId = _const.MessageId(message[1])
 }
  switch message[0] {
	case "Add":
		return api.AddReaction(data.ChannelID, data.ID, "💗") //adds a heart to the current emoji
	case "Remove":
		return api.DeleteReaction(data.ChannelID, MessageId, "💗") //removes the heart from the specified message
	case "RemoveAll":
		return api.DeleteAllReactions(data.ChannelID, MessageId) // removes all emojis from the specified message
	case "RemoveAllHeart":
		return api.DeleteAllReactionsForEmoji(data.ChannelID, MessageId, "💗") // removes all hearts from the specified message.
	}

	return nil
}

```