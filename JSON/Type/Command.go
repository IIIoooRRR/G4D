package Type

const (
	// EventReady Основные события жизненного цикла
	EventReady   = "READY"
	EventResumed = "RESUMED"

	// EventMessage События сообщений
	EventMessageCreate = "MESSAGE_CREATE"
	EventMessageUpdate = "MESSAGE_UPDATE"
	EventMessageDelete = "MESSAGE_DELETE"

	// EventGuild События сервера и участников
	EventGuildCreate       = "GUILD_CREATE"
	EventGuildUpdate       = "GUILD_UPDATE"
	EventGuildMemberAdd    = "GUILD_MEMBER_ADD"
	EventGuildMemberRemove = "GUILD_MEMBER_REMOVE"

	// EventInteraction Взаимодействия (кнопки, слэш-команды)
	EventInteractionCreate = "INTERACTION_CREATE"

	//EventReaction Реакции
	EventMessageReactionAdd       = "MESSAGE_REACTION_ADD"
	EventMessageReactionRemove    = "MESSAGE_REACTION_REMOVE"
	EventMessageReactionRemoveAll = "MESSAGE_REACTION_REMOVE_ALL"
)
