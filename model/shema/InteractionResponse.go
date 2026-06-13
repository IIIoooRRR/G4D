package shema

import (
	"github.com/IIIoooRRR/G4D/model/dependencies"
	"github.com/IIIoooRRR/G4D/model/dependencies/ui"
)

type InteractionResponse struct {
	Type int                     `json:"type"`
	Data InteractionResponseData `json:"data"`
}

type InteractionResponseData struct {
	Content    string               `json:"content,omitempty"`
	Components []ui.ActionRow       `json:"components,omitempty"`
	CustomID   string               `json:"custom_id,omitempty"`
	Title      string               `json:"title,omitempty"`
	Flags      int                  `json:"flags,omitempty"`
	Embeds     []dependencies.Embed `json:"embeds,omitempty"`
}

func NewInteractionResponse(Type int) InteractionResponse {
	return InteractionResponse{
		Type: Type,
	}
}

func (ir InteractionResponse) SetData(irData InteractionResponseData) InteractionResponse {
	ir.Data = irData
	return ir
}

func (ir InteractionResponse) SetType(Type int) InteractionResponse {
	ir.Type = Type
	return ir
}

func NewInteractionResponseData(Content string) *InteractionResponseData {
	return &InteractionResponseData{
		Content: Content,
	}
}
func (irData InteractionResponseData) AddEmbed(em ...dependencies.Embed) InteractionResponseData {
	irData.Embeds = append(irData.Embeds, em...)
	return irData
}
func (irData InteractionResponseData) AddFlags(flags int) InteractionResponseData {
	irData.Flags += flags
	return irData
}
func (irData InteractionResponseData) AddActionRow(row ...ui.ActionRow) InteractionResponseData {
	irData.Components = append(irData.Components, row...)
	return irData
}
func (irData InteractionResponseData) SetTitle(title string) InteractionResponseData {
	irData.Title = title
	return irData
}
func (irData InteractionResponseData) SetCustomID(customID string) InteractionResponseData {
	irData.CustomID = customID
	return irData
}
