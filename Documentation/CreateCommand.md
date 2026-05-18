##Create a command. old method
```go
... in main
cmd := Bot.Command{			//to create a command, create a type and use the Action interface.
		Trigger:  parseJSON.EventMessageCreate,
		Action: &FirstCommand{},
	}
	bot.CommandBuffer = append(bot.CommandBuffer, cmd) 
	... in main
}

----- outside the main function:
type FirstCommand struct {
}

func (f *FirstCommand) Execute(event *ConnectToDiscord.RawEvent) error {
log.Printf("DEBUG DATA: %+v", event)
var d Dparam.Message
if err := json.Unmarshal(event.Data, &d); err != nil {
return err
}
if d.Content == "!hello" {
msg := Bot.Message{
ToChannel: d.ChannelID,
Content:   "hello world",
}
msg.SendMessage(msg.ToChannel, b)
}
return nil
}
```
## Create a command. main method(works under the hood)
```go
... in main

bot.AddCommand(MessageTemplate{parseJSON.EventMessageCreate, Hello})
... in main
----- outside the main function:
func Hello(event *ConnectToDiscord.RawEvent) error {
log.Printf("DEBUG DATA: %+v", event)
var d Dparam.Message
if err := json.Unmarshal(event.Data, &d); err != nil {
return err
}
if d.Content == "!hello" {
msg := Bot.Message{
Content:   "hello world",
}
msg.SendMessage(msg.ToChannel, b)
}
return nil
}
```
## Create a command. The main method of creating a team of commands
```go
... in main
bot.AddCommands([]Bot.CommandTemplate{
		{
			JSON.EventMessageCreate,
			Create,
		},
		{
			JSON.EventMessageDelete,
			Execute,
		},
	})
... in main
----- outside the main function:

func Execute(event *Connect.RawEvent) error {
d := Parse.ToMessageDelete(*event)
embed := Dependencies.Embed{
Title:       "hello",
Type:        "text",
Description: "hi hi hi",
}
msg := Event.Message{
Content: "message deleted " + time.Now().GoString(),
}
msg.AddEmbed(embed)
err := msg.SendMessage(d.ChannelID, b)
if err != nil {
return err
}
return nil
}
func Create(event *Connect.RawEvent) error {
d := Parse.ToMessageCreate(*event)
if d.Content == "!hello" {
msg := Event.Message{
Content: "hello world",
}
err := msg.SendMessage(d.ChannelID, b)
if err != nil {
return err
}
}
return nil
```