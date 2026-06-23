package _const

const (
	EventMessageCreate     = "MESSAGE_CREATE"
	EventMessageUpdate     = "MESSAGE_UPDATE"
	EventMessageDelete     = "MESSAGE_DELETE"
	EventMessageDeleteBulk = "MESSAGE_DELETE_BULK"

	EventInteractionCreate = "INTERACTION_CREATE"
	EventInteractionUpdate = "INTERACTION_UPDATE"
	EventInteractionDelete = "INTERACTION_DELETE"

	EventMessageReactionAdd         = "MESSAGE_REACTION_ADD"
	EventMessageReactionRemove      = "MESSAGE_REACTION_REMOVE"
	EventMessageReactionRemoveAll   = "MESSAGE_REACTION_REMOVE_ALL"
	EventMessageReactionRemoveEmoji = "MESSAGE_REACTION_REMOVE_EMOJI"

	/*
		EventVoiceStateUpdate  = "VOICE_STATE_UPDATE"
		EventVoiceServerUpdate = "VOICE_SERVER_UPDATE"

		EventUserUpdate     = "USER_UPDATE"
		EventPresenceUpdate = "PRESENCE_UPDATE"
		EventTypingStart    = "TYPING_START"
	*/

	EventChannelCreate = "CHANNEL_CREATE"
	EventChannelUpdate = "CHANNEL_UPDATE"
	EventChannelDelete = "CHANNEL_DELETE"

	EventGuildBanAdd       = "GUILD_BAN_ADD"
	EventGuildBanRemove    = "GUILD_BAN_REMOVE"
	EventGuildEmojisUpdate = "GUILD_EMOJIS_UPDATE"
	EventGuildCreate       = "GUILD_CREATE"
	EventGuildUpdate       = "GUILD_UPDATE"
	EventGuildMemberAdd    = "GUILD_MEMBER_ADD"
	EventGuildMemberRemove = "GUILD_MEMBER_REMOVE"

	/*
		    GUILD_INTEGRATIONS_UPDATE
		    GUILD_MEMBER_UPDATE
		    GUILD_MEMBERS_CHUNK
		    GUILD_ROLE_CREATE
		    GUILD_ROLE_UPDATE
		    GUILD_ROLE_DELETE
		    GUILD_SCHEDULED_EVENT_CREATE
		    GUILD_SCHEDULED_EVENT_UPDATE
		    GUILD_SCHEDULED_EVENT_DELETE
		    GUILD_SCHEDULED_EVENT_USER_ADD
			GUILD_SCHEDULED_EVENT_USER_REMOVE

	*/
	//EventChannelPinsUpdate = "CHANNEL_PINS UPDATE"
	//EventWebhooksUpdate = "WEBHOOKS_UPDATE"
	/*
		EventMessagePollVoteAdd = "MESSAGE_POLL_VOTE_ADD"
		EventMessagePollVoteRemove = "MESSAGE_POLL_VOTE_REMOVE"
	*/
	/*
		EventStageInstanceCreate = STAGE_INSTANCE_CREATE
		EventStageInstanceUpdate = STAGE_INSTANCE_UPDATE
		EventStageInstanceDelete = STAGE_INSTANCE_DELETE
	*/
	/*
		EventThreadCreate = =THREAD_CREATE
		EventThreadUpdate = THREAD_UPDATE
		EventThreadDelete = THREAD_DELETE
		EventThreadListSync = THREAD_LIST_SYNC
		EventThreadMemberUpdate = THREAD_MEMBER_UPDATE
	*/
)
