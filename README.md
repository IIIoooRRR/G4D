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
func main() {
token := "token"
gateway := Connect.Receiver{
QueueSize: 20,
Intents:   34307,
}
bot := G4D.Bot{
Token:   token,
Gateway: &gateway,
Cache:   nil,
Context: context.Background(),
Prefix:  "!",
}

bot.AddCommands([]G4D.CommandTemplate{
{Trigger: Type.EventGuildCreate, Action: Commands.LALA},
{Trigger: Type.EventMessageCreate, Action: Hello},
})

go bot.DynamicEventProcessor()
bot.Run()
}
func Hello(event *Connect.RawEvent) error {
log.Println(time.Now())
data := Parse.ToMessageCreate(*event)
if data.Content == "hello" {
err := Functions.DeleteMessage(data.ChannelID, data.ID)
if err != nil {
log.Println(err)
}
}
log.Println(time.Now())
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
The project documentation is available on godoc
---
**Author:** [IIIoooRRR](https://github.com)
---