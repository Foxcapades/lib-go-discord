package audit

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type Log interface {
	// Webhooks returns the current value of this record's `webhooks` field.
	//
	// The `webhooks` field contains a list of webhooks found in the audit log.
	Webhooks() []discord.Webhook

	// SetWebhooks overwrites the current value of this record's `webhooks` field.
	SetWebhooks([]discord.Webhook) Log

	// Users returns the current value of this record's `users` field.
	//
	// The `users` field contains a list of users found in the audit log.
	Users() []discord.User

	// SetUsers overwrites the current value of this record's `users` field.
	SetUsers([]discord.User) Log

	// AuditLogEntries returns the current value of this record's
	// `audit_log_entries` field.
	//
	// The `audit_log_entries` field contains a list of audit log entries.
	AuditLogEntries() []LogEntry

	// SetAuditLogEntries overwrites the current value of this record's
	// `audit_log_entries` field.
	SetAuditLogEntries([]LogEntry) Log

	// Integrations returns the current value of this record's `integrations`
	// field.
	//
	// The `integrations` field contains a list of partial integration objects.
	Integrations() []discord.Integration

	// SetIntegrations overwrites the current value of this record's
	// `integrations` field.
	SetIntegrations([]discord.Integration) Log
}
