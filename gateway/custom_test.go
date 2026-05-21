package gateway_test

import (
	"github.com/IIIoooRRR/G4D/gateway"
	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/IIIoooRRR/G4D/model/customize"
	"github.com/IIIoooRRR/G4D/model/dependencies"
)

// Example_customization demonstrates all gateway customization options
func Example_customization() {
	// Custom activity (streaming/game)
	details := "Go to coding"
	state := "Believe"

	activity := customize.Activity{
		Name:    "Coding",
		Type:    _const.ActivityStreaming,
		Details: &details,
		State:   &state,
	}

	// Build gateway with all options
	gateway := gateway.NewGateway().
		WithActivity(activity).                // Add activity
		WithNetStatus(_const.NetStatusOnline). // Set status (online/idle/dnd/invisible)
		WithIntents(34307).                    // Set event intents
		WithQueueSize(20)                      // Set event buffer size

	_ = gateway
}

// Example_customStatus demonstrates custom status (text + emoji)
func Example_customStatus() {

	gateway := gateway.NewGateway().
		WithDescription("Building a bot 💗", dependencies.Emoji{}). // Custom status text
		WithNetStatus(_const.NetStatusIDLE)                        // DND status

	_ = gateway
}

// Example_multipleActivities demonstrates rich presence with multiple activities
func Example_multipleActivities() {
	game := customize.Activity{
		Name: "My Awesome Game",
		Type: _const.ActivityGame,
	}

	listening := customize.Activity{
		Name: "Lo-Fi Beats",
		Type: _const.ActivityListening,
	}

	gateway := gateway.NewGateway().
		WithActivity(game).
		WithActivity(listening).
		WithNetStatus(_const.NetStatusDND)

	_ = gateway
}
