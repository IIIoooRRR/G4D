# G4D — Go for Discord 🚀

**G4D** is a modular and lightweight library written in Go for building Discord bots. The project is designed with a focus on clear separation of concerns (Gateway, API, Event Processing), inspired by structured architectural patterns

> **Project Status:** 🛠️ Educational project under active development. It is an excellent resource for learning the inner workings of the Discord API and WebSocket management in Golang
## The basic information is provided in the root .md files.
CreateCommand.md
Bot.md
Structure.md
## ✨ Features

- **Bot-Centric Architecture**: Centralized management of all modules via a main `Bot` structure
- **Decoupled Gateway & API**: Networking logic is strictly separated from interaction methods for easier maintenance and refactoring
- **Command Interface (Executor)**: A flexible system for creating custom event handlers using interfaces
- **Minimalist Approach**: No hidden "magic" — you have full control over the data flow and event processing

## 📦 Installation

```bash
go get github.com/IIIoooRRR/G4D
```

## 🚀 Quick Start

Here is an example of a simple bot that responds to the `!hello` command:

```go
var token = ""
var details = "Go to codding"
var state = "Believe"
var activity = customize.Activity{
	Name:    "Codding",
	Type:    Type.ActivityStreaming,
	Details: &details,
	State:   &state,
}
var bot = G4D.Bot{
	Token:   token,
	Gateway: connect.NewGateway().WithActivity(activity).WithNetStatus(Type.NetStatusOnline).WithIntents(34307).WithQueueSize(20),
	Context: context.Background(),
	Prefix:  "!",
}

func main() {
	bot.SetBotDescription("Example Bot")
	bot.AddCommands([]G4D.CommandTemplate{
		{Trigger: Type.EventGuildCreate, Action: Commands.RolePermCache},
		{Trigger: Type.EventMessageCreate, Action: Hello},
	})
	go bot.DynamicEventProcessor()
	bot.Run()
}

func Hello(event *connect.RawEvent) error {
	data := Parse.ToMessageCreate(*event)
	message := strings.Split(data.Content, " ")
	var MessageId Type.MessageId
	if len(message) > 2 {
		MessageId = Type.MessageId(message[1])
	}
	switch message[0] {
	case "Add":
		functions.AddReaction(data.ChannelID, data.ID, "💗") //adds a heart to the current emoji
	case "Remove":
		functions.DeleteReaction(data.ChannelID, MessageId, "💗") //removes the heart from the specified message
	case "RemoveAll":
		functions.DeleteAllReactions(data.ChannelID, MessageId) // removes all emojis from the specified message
	case "RemoveAllHeart":
		functions.DeleteAllReactionsForEmoji(data.ChannelID, MessageId, "💗") // removes all hearts from the specified message.
	default:
		return nil
	}
	msg := Parse.NewMessage().AddReferencedMessage(data).AddContent("hihi")
	err := functions.SendMessage(data.ChannelID, msg)
	if err != nil {
		return err
	}
	return nil
}

```


## 🛠 Project Structure

The project follows a modular hierarchy inspired by structured programming:
- `G4D/` — Core logic, command management, and lifecycle.
- `ConnectToDiscord/` — Low-level WebSocket (Gateway) and network handling.
- `parseJSON/` — Specialized packages for parsing and typifying Discord data structures.

## ⚡ Philosophy

- **Full control** over every event and decision.
- **No magic** — you see what the library does.
- **Backward compatibility** — your old code keeps working.

> **Any problem is the user's problem :)**
>
> G4D gives you tools, not guarantees. If you cause a race condition, a panic, or a memory leak — that's your call. You're an engineer. Act like one.

## 🤝 Contributing

Contributions are welcome! Feel free to:
1. Fix bugs or report issues.
2. Implement new REST API endpoints.
3. Improve documentation and examples.

---
godoc: https://pkg.go.dev/github.com/IIIoooRRR/G4D
---
**Author:** [IIIoooRRR](https://github.com)
---
