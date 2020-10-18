package user

type Connection interface {
	Id() string
	SetId(string) Connection

	Name() string
	SetName(string) Connection

	Type() string
	SetType(string) Connection

	Revoked() bool
	RevokedIsSet() bool
	SetRevoked(bool) Connection
	UnsetRevoked() Connection

	Integrations() []server.Integration
	IntegrationsIsSet() bool
	SetIntegrations([]server.Integration) Connection
	UnsetIntegrations() Connection

	Verified() bool
	SetVerified(bool) Connection

	FriendSync() bool
	SetFriendSync(bool) Connection

	ShowActivity() bool
	SetShowActivity(bool) Connection

	Visibility() ConnectionVisibility
	SetVisibility(ConnectionVisibility) Connection
}