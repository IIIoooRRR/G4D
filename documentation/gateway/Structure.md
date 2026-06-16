## g4d
The main package for working with data.\
It is controlled by the user and has an Event sub-package that allows you to interact with discord - to parse or send/track data.\
There will be an Event sub-package that will expand
## gateway
The main package for interacting with the Discord API.\ Responsible for listening to received events and connecting to Discord
## model
Independent storage of data models.\
It contains all the Discord API structures (Message, User, Embed) and tools for parsing them from raw JSON.

├── api
│   ├── channel.go
│   ├── doDiscordRequest.go
│   ├── message.go
│   ├── reaction.go
│   └── restriction.go
├── config.locale.yaml
├── coverage.out
├── documentation
│   ├── BaseBot.md
│   ├── config.yaml
│   ├── gateway
│   │   ├── Bot.md
│   │   ├── Connect.md
│   │   ├── EventProcessor.md
│   │   └── Structure.md
│   ├── Reactions.md
│   └── SendMessage.md
├── g4d
│   ├── addCommand.go
│   ├── bot.go
│   ├── command.go
│   ├── config_func.go
│   ├── config.go
│   ├── description.go
│   ├── eventProcessor.go
│   ├── eventProcessor_test.go
│   ├── GetBotInfo.go
│   └── SlashCommand.go
├── gateway
│   ├── connect.go
│   ├── createBot.go
│   ├── custom_test.go
│   ├── dispatch.go
│   ├── gateway.go
│   ├── heartbeat.go
│   ├── hello.go
│   ├── helpers.go
│   ├── identify.go
│   ├── receiver.go
│   ├── resume.go
│   ├── status.go
│   └── stop.go
├── global_test
│   └── bot_test.go
├── go.mod
├── go.sum
├── helpers
│   └── user.go
├── LICENSE
├── model
│   ├── codec
│   │   ├── identify.go
│   │   ├── payload.go
│   │   ├── ready.go
│   │   └── resume.go
│   ├── _const
│   │   ├── button.go
│   │   ├── channel.go
│   │   ├── command.go
│   │   ├── component.go
│   │   ├── customize.go
│   │   ├── dashboard.go
│   │   ├── interaction.go
│   │   ├── permission.go
│   │   └── types.go
│   ├── customize
│   │   └── status.go
│   ├── dependencies
│   │   ├── Attachment.go
│   │   ├── Embed.go
│   │   ├── Emoji.go
│   │   ├── GuildMember.go
│   │   ├── Role.go
│   │   ├── ui
│   │   │   ├── button.go
│   │   │   ├── component.go
│   │   │   ├── input.go
│   │   │   └── select_menu.go
│   │   ├── upload.go
│   │   ├── User.go
│   │   └── VoiceState.go
│   ├── gateway
│   │   └── raw_event.go
│   ├── opcode
│   │   └── Памятка
│   ├── parse
│   │   ├── init.go
│   │   ├── Parser.go
│   │   └── parser_test.go
│   └── shema
│       ├── Channel.go
│       ├── Guilds.go
│       ├── Interaction.go
│       ├── InteractionResponse.go
│       ├── Message.go
│       └── Reactions.go
├── README.md
└── test
└── main.go
