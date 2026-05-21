package connect

import (
	"github.com/IIIoooRRR/G4D/JSON/customize"
)

func (r *Receiver) SetActivity(activity *customize.Activity) error {
	err := r.connectWS.WriteJSON(activity)
	if err != nil {
		return err
	}
	return nil
}
