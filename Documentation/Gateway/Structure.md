├── connect
│   ├── Connect.go
│   ├── CreateBot.go
│   ├── dispatch.go
│   ├── gateway.go
│   ├── heartbeat.go
│   ├── hello.go
│   ├── helpers.go
│   ├── identify.go
│   ├── Receiver.go
│   ├── resume.go
│   ├── status.go
│   └── Stop.go
├── Documentation
│   ├── Bot.md
│   ├── CreateCommand.md
│   ├── EventProcessor.md
│   ├── SendMessage.md
│   └── Structure.md
├── functions
│   ├── Channel.go
│   ├── Message.go
│   ├── Reaction.go
│   ├── Restriction.go
│   └── User.go
├── G4D
│   ├── AddCommand.go
│   ├── Bot.go
│   ├── bot_test.go
│   ├── Command.go
│   ├── commands_test.go
│   ├── decription.go
│   ├── EventProcessor.go
│   ├── eventProcessor_test.go
│   ├── GetBotInfo.go
│   └── SlashCommand.go
├── go.mod
├── go.sum
├── JSON
│   ├── customize
│   │   └── status.go
│   ├── Dependencies
│   │   ├── Attachment.go
│   │   ├── Embed.go
│   │   ├── Emoji.go
│   │   ├── GuildMember.go
│   │   ├── Role.go
│   │   ├── User.go
│   │   └── VoiceState.go
│   ├── Identify.go
│   ├── opcode
│   │   └── Памятка
│   ├── Parse
│   │   ├── Channel.go
│   │   ├── Guilds.go
│   │   ├── Interaction.go
│   │   ├── Message.go
│   │   ├── Parser.go
│   │   └── Reactions.go
│   ├── Payload.go
│   ├── Ready.go
│   ├── Resume.go
│   └── Type
│       ├── Channel.go
│       └── Command.go
├── LICENSE
├── README.md
└── test
├── Commands
│   └── guildc.go
└── test.go
## G4D
The main package for working with data.\
It is controlled by the user and has an Event sub-package that allows you to interact with discord - to parse or send/track data.\
There will be an Event sub-package that will expand
## CONNECT
The main package for interacting with the Discord API.\ Responsible for listening to received events and connecting to Discord
## JSON
Independent storage of data models.\
It contains all the Discord API structures (Message, User, Embed) and tools for parsing them from raw JSON.