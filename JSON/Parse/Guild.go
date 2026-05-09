package Parse

import (
	"encoding/json"

	"github.com/IIIoooRRR/G4D/Connect"
	"github.com/IIIoooRRR/G4D/G4D/Event"
)

func ToGuildMemberAdd(event Connect.RawEvent) Event.GuildMemberAdd {
	var d Event.GuildMemberAdd
	json.Unmarshal(event.Data, &d)
	return d
}
func ToGuildMemberRemove(event Connect.RawEvent) Event.GuildMemberRemove {
	var d Event.GuildMemberRemove
	json.Unmarshal(event.Data, &d)
	return d
}
