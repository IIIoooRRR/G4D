# G4D — Go for Discord 🚀

**G4D** is a partially modular library written in Go for creating Discord bots. The project is designed according to a clear architecture - each package and file has its own purpose.

> **Project Status:** , An actively developing project. Moved from the status of a "fully modular library" to a "library for enterprise development"## The basic information is provided in the root .md files.

## ✨ Features
- **Bot-oriented architecture**: Centralized management of all modules through the basic bot structure
- **Isolated gateway and API**: Network logic is strictly separated from interaction methods to simplify maintenance and refactoring.
- **Enterprise approach**: Earlier I promised no magic. Now I retract my words. g4d is a library that will focus on best-practice solutions. The library supports *.yaml configs, and also works on zap, Godot, sonic
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
import (
	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/g4d"
	way "github.com/IIIoooRRR/G4D/gateway"
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/ctx"
	"github.com/IIIoooRRR/G4D/model/customize"
	"github.com/IIIoooRRR/G4D/model/parse"
	shema "github.com/IIIoooRRR/G4D/model/schema"
	"go.uber.org/zap"
)

func main() {
	var token = ""
	var prefix = ""
	logger := zap.Must(zap.NewProduction()).Named("bot")
	bot := g4d.NewBot(
		token,
		PanicHandler{Logger: logger.Named("panic-handler")},
		way.NewGateway(way.BufferSize(15),
			way.Intents(34307),
			way.Activity(customize.Activity{
				Name: "With G4D",
				Type: _const.ActivityStreaming,
			}),
			way.NetStatus(_const.NetStatusIDLE)).WithDescription("hello!"),
		logger,
		api.NewClient(&token, 10),
	)

	bot.SetPrefix(prefix).SetBotBio("test bot").AddCommands([]g4d.CommandTemplate{
		{_const.EventMessageCreate, "", Button},
		{_const.EventInteractionCreate, "", ButtonReaction},
	})

	if err := bot.Run(g4d.WithDispQuantity(2),
		g4d.WithSemaphoreLimit(12),
		_const.StaticEventProcessor); err != nil {
		logger.Panic(err.Error())
	}
}

type PanicHandler struct {
	g4d.PanicHandler
	Logger *zap.Logger
}

func (p PanicHandler) OnPanic(event *parse.RawEvent, cmd *g4d.CommandTemplate, r any, stack []byte) {
	p.Logger.Panic("cmd:", zap.Any("cmd", cmd), zap.String("stack", string(stack)), zap.Any("r:", r))
}

```

## Commands:
```go
func Button(event *parse.RawEvent, ctx *ctx.Context) error {
	data := parse.GetEvent[shema.GetMessage](event)
	if  data.Content != ctx.Prefix+"button" {
		return nil
	}
	return ctx.SendMessageWithLimit(data.ChannelID, shema.NewMessage().AddReferencedMessage(data).AddContent("pipi"))
}

func ButtonReaction(event *parse.RawEvent, ctx *ctx.Context) error {
	d := parse.GetEvent[shema.Interaction](event)
	if d.Data.ComponentType != _const.ButtonType && d.Data.CustomID == "button" {
		return nil
	}

	data := shema.NewInteractionResponseData("hihi!")
	response := shema.NewInteractionResponse(_const.ResponseChannelMessageWithSource).SetData(*data)
	return ctx.SendInteractionMessage(d, response)
}
```
## 🛠 Project Structure

The project follows a modular hierarchy inspired by structured programming:
G4D/\
├── api/           # REST API client\
├── g4d/           # Core logic, commands, processors\
├── gateway/       # WebSocket connection\
├── model/         # Data structures and parsing\
│   ├── parse/     # Event cache and parsing\
│   ├── other...\
│   └── schema/     # Discord API structures\
├── cli/           # CLI tool (separate module)\
└── test/          # Tests and examples\
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
go-doc: https://pkg.go.dev/github.com/IIIoooRRR/G4D
---
**Author:** [IIIoooRRR](https://github.com)
---
