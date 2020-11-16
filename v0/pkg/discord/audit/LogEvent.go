package audit

import "errors"

var (
	ErrBadLogEvent = errors.New("unrecognized audit log event value")
)

type LogEvent uint8

const (
	LogEventGuildUpdate            LogEvent = 1
	LogEventChannelCreate          LogEvent = 10
	LogEventChannelUpdate          LogEvent = 11
	LogEventChannelDelete          LogEvent = 12
	LogEventChannelOverwriteCreate LogEvent = 13
	LogEventChannelOverwriteUpdate LogEvent = 14
	LogEventChannelOverwriteDelete LogEvent = 15
	LogEventMemberKick             LogEvent = 20
	LogEventMemberPrune            LogEvent = 21
	LogEventMemberBanAdd           LogEvent = 22
	LogEventMemberBanRemove        LogEvent = 23
	LogEventMemberUpdate           LogEvent = 24
	LogEventMemberRoleUpdate       LogEvent = 25
	LogEventMemberMove             LogEvent = 26
	LogEventMemberDisconnect       LogEvent = 27
	LogEventBotAdd                 LogEvent = 28
	LogEventRoleCreate             LogEvent = 30
	LogEventRoleUpdate             LogEvent = 31
	LogEventRoleDelete             LogEvent = 32
	LogEventInviteCreate           LogEvent = 40
	LogEventInviteUpdate           LogEvent = 41
	LogEventInviteDelete           LogEvent = 42
	LogEventWebhookCreate          LogEvent = 50
	LogEventWebhookUpdate          LogEvent = 51
	LogEventWebhookDelete          LogEvent = 52
	LogEventEmojiCreate            LogEvent = 60
	LogEventEmojiUpdate            LogEvent = 61
	LogEventEmojiDelete            LogEvent = 62
	LogEventMessageDelete          LogEvent = 72
	LogEventMessageBulkDelete      LogEvent = 73
	LogEventMessagePin             LogEvent = 74
	LogEventMessageUnpin           LogEvent = 75
	LogEventIntegrationCreate      LogEvent = 80
	LogEventIntegrationUpdate      LogEvent = 81
	LogEventIntegrationDelete      LogEvent = 82
)

func (l LogEvent) JSONSize() int {
	if l > 1 {
		return 2
	}

	return 1
}

func (l LogEvent) IsValid() bool {
	return nil == l.Validate()
}

func (l LogEvent) Validate() error {
	if l > 82 || l < 1 {
		return ErrBadLogEvent
	}

	switch true {
	case l == LogEventGuildUpdate:
		return nil
	case l >= LogEventChannelCreate && l <= LogEventChannelOverwriteDelete:
		return nil
	case l >= LogEventMemberKick && l <= LogEventBotAdd:
		return nil
	case l >= LogEventRoleCreate && l <= LogEventRoleDelete:
		return nil
	case l >= LogEventInviteCreate && l <= LogEventInviteDelete:
		return nil
	case l >= LogEventWebhookCreate && l <= LogEventWebhookDelete:
		return nil
	case l >= LogEventEmojiCreate && l <= LogEventEmojiDelete:
		return nil
	case l >= LogEventMessageDelete && l <= LogEventMessageUnpin:
		return nil
	case l >= LogEventIntegrationCreate && l <= LogEventIntegrationDelete:
		return nil
	}

	return ErrBadLogEvent
}
