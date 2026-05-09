├── Bot
│        ├──
│        ├── Message.go
│        ├── Guilds.go
│    ├── AddCommand.go  
│    ├── Bot.go
│    ├── Command.go
│    ├── EventProcessor.go
│    ├── Guilds.go
│    └── Message.go
├── Bot.md
├── ConnectToDiscord
│         ├── Connect.go
│         ├── CreateBot.go
│         ├── dispatch.go
│         ├── gateway.go
│         ├── heartbeat.go
│         ├── hello.go
│         ├── identify.go
│         ├── Receiver.go
│         ├── resume.go
│         └── Stop.go
├── CreateCommand.md
├── go.mod
├── go.sum
├── JSON
│         ├── Dependencies
│         │ ├── Attachment.go
│         │ ├── Channel.go
│         │ ├── Embed.go
│         │ ├── Emoji.go
│         │ ├── Role.go
│         │ └── User.go
│         ├── Identify.go
│         ├── opcode
│         │ └── Памятка
│         ├── Parse
│         │ └── Message.go
│         ├── Payload.go
│         ├── Ready.go
│         └── Resume.go
├── LICENSE
├── README.md
└── Structure.md

## BOT
The main package for working with data.\
It is controlled by the user and has an Event sub-package that allows you to interact with discord - to parse or send/track data.\
There will be an Event sub-package that will expand
## CONNECT
The main package for interacting with the Discord API.\ Responsible for listening to received events and connecting to Discord
## JSON
Independent storage of data models.\
It contains all the Discord API structures (Message, User, Embed) and tools for parsing them from raw JSON.