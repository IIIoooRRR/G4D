package connect_test

import (
	"github.com/IIIoooRRR/G4D/JSON/Dependencies"
	"github.com/IIIoooRRR/G4D/JSON/Type"
	"github.com/IIIoooRRR/G4D/JSON/customize"
	"github.com/IIIoooRRR/G4D/connect"
)

// Example_customization demonstrates all gateway customization options
func Example_customization() {
	// Custom activity (streaming/game)
	details := "Go to coding"
	state := "Believe"

	activity := customize.Activity{
		Name:    "Coding",
		Type:    Type.ActivityStreaming,
		Details: &details,
		State:   &state,
	}

	// Build gateway with all options
	gateway := connect.NewGateway().
		WithActivity(activity).               // Add activity
		WithNetStatus(Type.NetStatusOnline).  // Set status (online/idle/dnd/invisible)
		WithIntents(34307).                   // Set event intents
		WithQueueSize(20)                     // Set event buffer size

	_ = gateway
}

// Example_customStatus demonstrates custom status (text + emoji)
func Example_customStatus() {

	gateway := connect.NewGateway().
		WithDescription("Building a bot 💗", Dependencies.Emoji{}). // Custom status text
		WithNetStatus(Type.NetStatusIDLE)  // DND status

	_ = gateway
}

// Example_multipleActivities demonstrates rich presence with multiple activities
func Example_multipleActivities() {
	game := customize.Activity{
		Name: "My Awesome Game",
		Type: Type.ActivityGame,
	}

	listening := customize.Activity{
		Name: "Lo-Fi Beats",
		Type: Type.ActivityListening,
	}

	gateway := connect.NewGateway().
		WithActivity(game).
		WithActivity(listening).
		WithNetStatus(Type.NetStatusDND)

	_ = gateway
}