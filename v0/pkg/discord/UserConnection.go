package discord

type UserConnection interface {
	// id of the connection account
	Id() string
	SetId(string) UserConnection

	// the username of the connection account
	Name() string
	SetName(string) UserConnection

	// the service of the connection (twitch, youtube)
	Type() string
	SetType(string) UserConnection

	// whether the connection is revoked
	Revoked() bool
	RevokedIsSet() bool
	SetRevoked(bool) UserConnection
	UnsetRevoked() UserConnection

	// an array of partial server integrations
	Integrations() []Integration
	IntegrationsIsSet() bool
	SetIntegrations([]Integration) UserConnection
	UnsetIntegrations() UserConnection

	// whether the connection is verified
	Verified() bool
	SetVerified(bool) UserConnection

	// whether friend sync is enabled for this connection
	FriendSync() bool
	SetFriendSync(bool) UserConnection

	// whether activities related to this connection will be shown in presence updates
	ShowActivity() bool
	SetShowActivity(bool) UserConnection

	// visibility of this connection
	Visibility() ConnectionVisibility
	SetVisibility(ConnectionVisibility) UserConnection
}