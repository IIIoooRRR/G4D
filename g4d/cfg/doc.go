// Package cfg loads bot configuration from g4d.yaml.
//
// It allows you to initialize the bot via the g4d.yaml file. The file is
// searched in the project root, so be careful about conflicts with other
// configs.
//
// To initialize the bot via the config, two fields are required in the code:
// a logger and a panic handler.
//
// Example:
//
//	logger := zap.Must(zap.NewProduction())
//	bot := cfg.LoadBot(logger, panicHandler{})
//
// Configuration structure:
//
//	bot:
//	  token: ${TOKEN}  # or hardcore, but u can use .env
//	  prefix: "r."
//	  description: "My Discord bot powered by G4D"
//
//	gateway:
//	  intents: 34307
//	  queue_size: 15
//	  presence_update:
//	    since: 0
//	    afk: false
//	    status: "idle"          # _const.NetStatus: online | idle | dnd | offline
//	    activities:
//	      - name: "With G4D"
//	        type: 1              # _const.Activity: 0=Playing, 1=Streaming, 2=Listening, 3=Watching, 5=Competing
//	        url: "https://twitch.tv/example"   # required for streaming
//	        details: "with g4d"
//	        state: "On the G4D framework"
//
// G4D resolves `${VAR}` placeholders in `g4d.yaml` from the process environment at load time
// // if the variable is not set via the environment, the bot will not start
package cfg
