package user

type Connection interface {
	// id of the connection account
	Id() string
	SetId(string) Connection

	// the username of the connection account
	Name() string
	SetName(string) Connection

	// the service of the connection (twitch, youtube)
	Type() string
	SetType(string) Connection

	// whether the connection is revoked
	Revoked() bool
	RevokedIsSet() bool
	SetRevoked(bool) Connection
	UnsetRevoked() Connection

	// an array of partial server integrations
	Integrations() []Integration
	IntegrationsIsSet() bool
	SetIntegrations([]Integration) Connection
	UnsetIntegrations() Connection

	// whether the connection is verified
	Verified() bool
	SetVerified(bool) Connection

	// whether friend sync is enabled for this connection
	FriendSync() bool
	SetFriendSync(bool) Connection

	// whether activities related to this connection will be shown in presence updates
	ShowActivity() bool
	SetShowActivity(bool) Connection

	// visibility of this connection
	Visibility() ConnectionVisibility
	SetVisibility(ConnectionVisibility) Connection
}