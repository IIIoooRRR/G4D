package gateway_test

import (
	"github.com/IIIoooRRR/G4D/gateway"
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/customize"
)

// Example_customization demonstrates all gateway customization options
func Example_customization() {
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
	gateway := gateway.NewGateway(23).
		WithActivity(activity).
		WithNetStatus(_const.NetStatusDND)

	_ = gateway
}

// Example_customStatus demonstrates custom status (text + emoji)
func Example_customStatus() {

	gateway := gateway.NewGateway(23).
		WithDescription("Building a bot 💗"). // Custom status text
		WithNetStatus(_const.NetStatusIDLE)  // DND status

	_ = gateway
}

// Example_multipleActivities demonstrates rich presence with multiple activities
func Example_multipleActivities() {

	listening := customize.Activity{
		Name: "Lo-Fi Beats",
		Type: _const.ActivityListening,
	}

	gateway := gateway.NewGateway(23).
		WithActivity(listening).
		WithNetStatus(_const.NetStatusDND)

	_ = gateway
}
