# G4D — Go for Discord 🚀

**G4D** is a partially modular library written in Go for creating Discord bots. The project is designed according to a clear architecture - each package and file has its own purpose.

> **Project Status:** , An actively developing project. Moved from the status of a "fully modular library" to a "library for enterprise development"## The basic information is provided in the root .md files.

## ✨ Features
- **Bot-oriented architecture**: Centralized management of all modules through the basic bot structure
- **Isolated gateway and API**: Network logic is strictly separated from interaction methods to simplify maintenance and refactoring.
- **Enterprise approach**: Earlier I promised no magic. Now I retract my words. g4d is a library that will focus on best-practice solutions. The library supports *.yaml configs, and also works on zap, godot, sonic
- **CLI-utils**: Since recent patches, I have started developing CLI utilities for working with the library.
## 📦 Installation

```bash
# Install the library
go get github.com/IIIoooRRR/G4D
# Install the CLI tool
go install github.com/IIIoooRRR/G4D/cli@latest
# Add ~/go/bin to PATH (if not already)
export PATH=$PATH:$(go env GOPATH)/
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
    Gateway: gateway.NewGateway(100 /*limit queue size*/).WithActivity(activity).WithNetStatus(_const.NetStatusOnline).WithIntents(34307),
    Context: context.Background(),
    Prefix:  "!",
}
func main() {
	go bot.SetBotDescription("Example Bot").
		AddCommands([]g4d.CommandTemplate{
			{Trigger: _const.EventMessageCreate, Name: "Hello", Action: cmd.Hello},
			{Trigger: _const.EventMessageDelete, Name: "Bye", Action: cmd.Bye},
			{Trigger: _const.EventMessageCreate, Name: "Menu", Action: Menu},
			{Trigger: _const.EventMessageCreate, Name: "Button", Action: Button},
			{Trigger: _const.EventGuildCreate, Name: "Input", Action: Input},
			{Trigger: _const.EventInteractionCreate, Name: "ButtonInteraction", Action: ButtonReaction},
		}).InitProcessors(g4d.StaticEventProcessor, 3, 35)
	bot.Run()
}
```
## OR
```go
func main() {
	cfg := g4d.MustLoadCfg("config.yaml")    // see documentation/config.yaml
	logger, err := zap.NewDevelopment()
	if err != nil {
		logger.Error("Error initializing logger", zap.Error(err))
	}
	bot, err := cfg.NewBot(logger, context.Background())
	go func() {
		bot.AddCommands([]g4d.CommandTemplate{
			{_const.EventMessageCreate, "", Menu},
			{_const.EventMessageCreate, "", Button},
			{_const.EventMessageUpdate, "", ButtonReaction},
		})
	}()
	if err != nil {
		logger.Error("Error initializing bot", zap.Error(err))
	}
	if err := bot.Run(); err != nil {
		logger.Error("Error bot run", zap.Error(err))
	}
}
```
## Commands: 
```go

func Menu(event *gateway.RawEvent) error {
	d := parse.GetEvent[shema.GetMessage](event)
	if d.Content != "menu" {
		return nil
	}
	menu := ui.NewMenu("test")
	menu.Options = append(menu.Options, ui.NewSelectOption("test", "12").SetDescription("test1"))
	row := ui.NewActionRow().AddComponents(menu)
	msg := shema.NewMessage().AddActionRow(row).AddContent("gello")
	return api.SendMessage(d.ChannelID, msg)
}
func Button(event *gateway.RawEvent) error {
	d := parse.GetEvent[shema.GetMessage](event)
	if d.Content != "button" {
		return nil
	}
	button := ui.NewButton("button").SetLabel("button").SetStyle(1)
	row := ui.NewActionRow().AddComponents(button)
	msg := shema.NewMessage().AddActionRow(row).AddContent("buttooonn!!!!")
	return api.SendMessage(d.ChannelID, msg)
}

func ButtonReaction(event *gateway.RawEvent) error {
	d := parse.GetEvent[shema.Interaction](event)
	if d.Data.ComponentType != _const.ButtonType && d.Data.ComponentCustomID == "button" {
		return nil
	}
	data := shema.NewInteractionResponseData("hihi!")
	response := shema.NewInteractionResponse(_const.InteractionApplicationCommandAutocomplete).SetData(*data)
	return api.SendInteractionMessage(d, response)
}

```
## 🛠 Project Structure

The project follows a modular hierarchy inspired by structured programming:
G4D/
├── api/           # REST API client
├── g4d/           # Core logic, commands, processors
├── gateway/       # WebSocket connection
├── model/         # Data structures and parsing
│   ├── parse/     # Event cache and parsing
│   ├── other...
│   └── shema/     # Discord API structures
├── cli/           # CLI tool (separate module)
└── test/          # Tests and examples
## Philosophy
- **You are in control of the situation** - G4D does not hide the complexity. You manage events, caching, and errors yourself.
- **Strict typing** - helps to avoid mistakes, but does not limit you.
- **No magic** - what you write happens. Every HTTP request, every WebSocket message is visible.
- **A framework with obvious limitations** - g4d has abandoned full modularity. some functions can still be rebuilt, but the most important is hidden. Read the source code
- **Functionality** - creation of an optimized (fast) library that will be safe for working with the discord api.
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
