package gateway

import (
	"github.com/IIIoooRRR/G4D/model/customize"
)

func (r *Receiver) SetActivity(activity *customize.Activity) error {
	err := r.connectWS.WriteJSON(activity)
	if err != nil {
		return err
	}
	return nil
}
