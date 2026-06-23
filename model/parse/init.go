package parse

import (
	"encoding/json"
	"runtime"

	"github.com/bytedance/sonic"
	"go.uber.org/zap"
)

/*
The encoder itself is determined based on your processor architecture - arm, amd, 64/32
This parser uses TikTok's Sonic technology to increase the speed of message analysis, or the basic encoder/JSON
is designed to save processor and RAM resources on amd64.
Sorry, I'm addicted to optimizing various operations and commands.
You don't need to worry about errors or conflicts when using this library.
Thank you if you read this.
*/

var (
	Marshal   func(any) ([]byte, error)
	Unmarshal func([]byte, any) error
	logger    = zap.Must(zap.NewProduction()).Named("parser")
)

func init() {
	if runtime.GOARCH == "amd64" {
		Marshal = sonic.Marshal
		Unmarshal = sonic.ConfigFastest.Unmarshal
		logger.Info("Encoder: sonic")
	} else {
		Marshal = json.Marshal
		Unmarshal = json.Unmarshal
		logger.Info("Encoder: encoder/json")
	}
}
