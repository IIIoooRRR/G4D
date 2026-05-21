package _const

const (
	InteractionPing                           = 1
	InteractionApplicationCommand             = 2
	InteractionMessageComponent               = 3
	InteractionApplicationCommandAutocomplete = 4
	InteractionModalSubmit                    = 5
)
const (
	ResponsePong                                 = 1
	ResponseChannelMessageWithSource             = 4
	ResponseDeferredChannelMessageWithSource     = 5
	ResponseDeferredUpdateMessage                = 6
	ResponseUpdateMessage                        = 7
	ResponseApplicationCommandAutocompleteResult = 8
	ResponseModal                                = 9
)
const (
	FlagsEphemeral = 1 << 6
)
