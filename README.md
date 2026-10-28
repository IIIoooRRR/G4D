# G4D — Go for Discord 🚀

**G4D** is a partially modular library written in Go for creating Discord bots. The project is designed according to a clear architecture - each package and file has its own purpose.

> **Project Status:** , An actively developing project. Moved from the status of a "fully modular library" to a "library for enterprise development"## The basic information is provided in the root .md files.
> **API design orientation:** introducing dependencies within the team through context, Flutter‑like layout for components.
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
	way "github.com/IIIoooRRR/G4D/gateway"
)

func main() {
	var token = os.Getenv("TOKEN")
	const prefix = "r."
	logger := zap.Must(zap.NewProduction()).Named("bot")
	bot := g4d.NewBot(
		token,
		PanicHandler{Logger: logger.Named("panic-handler")},
		way.NewGateway(way.BufferSize(15),
			way.Intents(34307),
			way.Activity(customize.Activity{
				Name:    "With G4D",
				Type:    _const.ActivityStreaming,
				Details: "with g4d",
			}),
			way.NetStatus(_const.NetStatusIDLE)),
		logger,
		api.NewClient(&token, 10*time.Second),
	)

	bot.SetPrefix(prefix).SetBotBio("test bot").AddCommands([]g4d.CommandTemplate{
		{_const.EventMessageCreate, "", Button},
		{_const.EventInteractionCreate, "", ButtonReaction},
		{_const.EventMessageCreate, "", UiTest},
	})

	botCtx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	bot.MemReport(botCtx, time.Second*30)
	bot.StartProfiling(botCtx)

	if err := bot.Run(
		g4d.WithProcessorsQuantity(2),
		g4d.WithSemaphoreLimit(12),
		_const.StaticEventProcessor,
		botCtx,
	); err != nil {
		logger.Error("error", zap.Error(err))
	}
}

```

## Commands:
```go
func Button(event *parse.RawEvent, ctx ctx.Context) error {
	data := parse.GetEvent[schema.GetMessage](event)
	if !ctx.IsCommand(data.Content, "button") { //find hard match
		return nil
	}
	msg := schema.NewMessage().
		Reply(data).
		Content("button").
		Components(
			ui.ActionRow(
				ui.DangerButton("button", "danger"),
			),
		)
	return ctx.SendMessageWithLimit(data.ChannelID, msg)
}

func ButtonReaction(event *parse.RawEvent, ctx ctx.Context) error {
	d := parse.GetEvent[schema.Interaction](event)
	if d.Data.ComponentType != _const.ComponentButtonType || d.Data.CustomID != "button" {
		return nil
	}

	data := schema.NewInteractionResponseData("hihi!")
	response := schema.NewInteractionResponse(_const.ResponseChannelMessageWithSource).SetData(*data)
	return ctx.SendInteractionMessage(d, response)
}

type PanicHandler struct {
	g4d.PanicHandler
	Logger *zap.Logger
}

func (p PanicHandler) OnPanic(event *parse.RawEvent, cmd *g4d.CommandTemplate, r any, stack []byte) {
	p.Logger.Error("cmd:", zap.Any("cmd", cmd), zap.String("stack", string(stack)), zap.Any("r:", r))
}

func UiTest(event *parse.RawEvent, ctx ctx.Context) error {
	data := parse.GetEvent[schema.GetMessage](event)
	if !ctx.Match(data.Content, "help") { // find soft match
		return nil
	}
	msg := schema.NewMessage().Flags(_const.FlagComponentsV2)
	cnt := ui.Container(
		ui.TextDisplay("1 display"),

		ui.MediaGallery().Add(
			ui.MediaItem(
				ctx.GetAvatarURL(data.Author.Id, data.Author.Avatar)).
				Description("true"),
		),

		ui.Section().Text("sections").
			Accessory(
				ui.DangerButton("custom_button", "danger"),
			),
	)
	msg.Components(cnt)
	return ctx.SendMessageWithMultipart(data.ChannelID, msg.Reply(data))
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
