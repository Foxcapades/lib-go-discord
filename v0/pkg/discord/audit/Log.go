package audit

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/user"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/webhook"
)

type Log interface {
	// Webhooks returns the current value of this record's `webhooks` field.
	//
	// The `webhooks` field contains a list of webhooks found in the audit log.
	Webhooks() []webhook.Webhook

	// SetWebhooks overwrites the current value of this record's `webhooks` field.
	SetWebhooks([]webhook.Webhook) Log

	// Users returns the current value of this record's `users` field.
	//
	// The `users` field contains a list of users found in the audit log.
	Users() []user.User

	// SetUsers overwrites the current value of this record's `users` field.
	SetUsers([]user.User) Log

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
	Integrations() []user.Integration

	// SetIntegrations overwrites the current value of this record's
	// `integrations` field.
	SetIntegrations([]user.Integration) Log
}
