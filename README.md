# G4D — Go for Discord 🚀

**G4D** is a modular and lightweight library written in Go for building Discord bots. The project is designed with a focus on clear separation of concerns (Gateway, API, Event Processing), inspired by structured architectural patterns

> **Project Status:** 🛠️ Educational project under active development. It is an excellent resource for learning the inner workings of the Discord API and WebSocket management in Golang
## The basic information is provided in the root .md files.

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

Here is an example of a simple bot that responds to the basic command:


```go
var details = "Go to codding"
var state = "Believe"
var activity = customize.Activity{
    Name:    "Codding",
    Type:    _const.ActivityStreaming,
    Details: &details,
    State:   &state,
}

var token = ""
var bot = g4d.Bot{
    Token:   token,
    Gateway: gateway.NewGateway().WithActivity(activity).WithNetStatus(_const.NetStatusOnline).WithIntents(34307).WithQueueSize(100),
    Context: context.Background(),
    Prefix:  "!",
}
func main() {
	bot.SetBotDescription("Example Bot").
		AddCommands([]g4d.CommandTemplate{
			{Trigger: _const.EventMessageCreate, Name: "Hello", Action: cmd.Hello},
			{Trigger: _const.EventMessageDelete, Name: "Bye", Action: cmd.Bye},
			{Trigger: _const.EventMessageCreate, Name: "Menu", Action: Menu},
			{Trigger: _const.EventMessageCreate, Name: "Button", Action: Button},
			{Trigger: _const.EventGuildCreate, Name: "Input", Action: Input},
			{Trigger: _const.EventInteractionCreate, Name: "ButtonInteraction", Action: ButtonReaction},
		})
	go bot.StaticEventProcessor()
	bot.Run()
}

func Menu(event *gateway.RawEvent) error {
	d := parse.Event[shema.GetMessage](*event)
	if d.Content != "menu" {
		return nil
	}
	menu := ui.NewMenu("test")
	menu.Options = append(menu.Options, ui.NewSelectOption("test", "12").SetDescription("test1"))
	row := ui.NewActionRow().AddComponents(menu)
	msg := _struct.NewMessage().AddActionRow(row).AddContent("gello")
	return api.SendMessage(d.ChannelID, msg)
}
func Button(event *gateway.RawEvent) error {
	d := parse.Event[schema.GetMessage](event)
	if d.Content != "button" {
		return nil
	}
	button := ui.NewButton("button").SetLabel("button").SetStyle(1)
	row := ui.NewActionRow().AddComponents(button)
	msg := _struct.NewMessage().AddActionRow(row).AddContent("buttooonn!!!!")
	return api.SendMessage(d.ChannelID, msg)
}

func ButtonReaction(event *gateway.RawEvent) error {
	d := parse.Event[schema.Interaction](event)
	if d.Data.ComponentType != _const.ButtonType {
		return nil
	}
	data := _struct.NewInteractionResponseData("hihi!")
	response := _struct.NewInteractionResponse(_const.InteractionApplicationCommandAutocomplete).SetData(*data)
    return api.SendInteractionMessage(d, response)
}

func Input(event *gateway.RawEvent) error {
	d := parse.Event[shema.GetMessage](event)
	if d.Content != "input" {
		return nil
	}
	input := ui.NewTextInput("input", "tututu", _const.InputParagraph)
	row := ui.NewActionRow().AddComponents(input)
	msg := _struct.NewMessage().AddActionRow(row).AddContent("tututu")
	return api.SendMessage(d.ChannelID, msg)
}

```



## 🛠 Project Structure

The project follows a modular hierarchy inspired by structured programming:
- `g4d/` — Core logic, command management, and lifecycle.
- `gateway/` — Low-level WebSocket (Gateway) and network handling.
- `model/` — Specialized packages for parsing and typifying Discord data structures.

## Philosophy
- **You are in control** - G4D doesn't hide complexity. You handle events, caching, and errors yourself.
- **Strict typing** - helps you avoid mistakes, but doesn't limit you.
- **No magic** - what you write is what happens. Every HTTP request, every WebSocket message is visible.
- **Your responsibility** - this is a tool, not a guarantee. If you break something, you know why.
- **Assembly, not framework** - you build your bot brick by brick.
- **FunFact** - creating an optimized (fast) library that will be safe for working with the discord api
> *"I give you the tools. The rest is up to you."*
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
