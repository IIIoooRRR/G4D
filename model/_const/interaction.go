package _const

const (
	InteractionPing                           = 1
	InteractionApplicationCommand             = 2
	InteractionMessageComponent               = 3
	InteractionApplicationCommandAutocomplete = 4
	InteractionModalSubmit                    = 5
)
const (
	IResponsePong                                 = 1
	IResponseChannelMessageWithSource             = 4
	IResponseDeferredChannelMessageWithSource     = 5
	IResponseDeferredUpdateMessage                = 6
	IResponseUpdateMessage                        = 7
	IResponseApplicationCommandAutocompleteResult = 8
	IResponseModal                                = 9
)
const (
	FlagsEphemeral = 1 << 6
)

const (
	InteractionNumber = 4
	InteractionString = 3
)
