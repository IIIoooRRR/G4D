package gateway_test

import (
	way "github.com/IIIoooRRR/G4D/gateway"
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/customize"
)

// Example_receiver init gateway
func ExampleReceiver_InitGateway() {
	// Custom activity (streaming/game)
	details := "Go to coding"
	state := "Believe"

	activity := customize.Activity{
		Name:    "Coding",
		Type:    _const.ActivityStreaming,
		Details: details,
		State:   state,
	}

	// Build gateway with all options
	gateway := way.NewGateway(way.BufferSize(15),
		way.Intents(34307),
		way.Activity(activity),
		way.NetStatus(_const.NetStatusIDLE)).WithDescription("hello!")
	_ = gateway
}
