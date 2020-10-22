package discord

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/integration"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
	"time"
)

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║    Errors                                                              ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║    Public Types                                                        ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// Integration
// TODO: Document me
type Integration interface {
	// integration id
	ID() dlib.Snowflake
	SetID(id dlib.Snowflake)

	// integration name
	Name() string
	SetName(name string) Integration

	// integration type (twitch, youtube, or discord)
	Type() string
	SetType(val string) Integration

	// is this integration enabled
	Enabled() bool
	SetEnabled(bool) Integration

	// is this integration syncing
	Syncing() bool
	SetSyncing(bool) Integration

	// id that this integration uses for "subscribers", or the guild id for discord integrations
	RoleID() dlib.Snowflake
	SetRoleID(id dlib.Snowflake) Integration

	// whether emoticons should be synced for this integration (twitch only currently)
	EnableEmoticons() bool
	EnableEmoticonsIsSet() bool
	SetEnableEmoticons(bool) Integration
	UnsetEnableEmoticons() Integration

	// the behavior of expiring subscribers
	ExpireBehavior() integration.ExpireBehavior
	SetExpireBehavior(integration.ExpireBehavior) Integration

	// the grace period (in days) before expiring subscribers
	ExpireGracePeriod() uint32
	SetExpireGracePeriod(days uint32) Integration

	// user for this integration
	User() User
	UserIsSet() bool
	SetUser(User) Integration
	UnsetUser() Integration

	// integration account information
	Account() integration.Account
	SetAccount(integration.Account) Integration

	// when this integration was last synced
	SyncedAt() time.Time
	SetSyncedAt(time.Time) Integration

	// how many subscribers this integration has (0 for discord integrations)
	SubscriberCount() uint32
	SetSubscriberCount(count uint32) Integration

	// has this integration been revoked
	Revoked() bool
	SetRevoked(bool) Integration

	// The bot/OAuth2 application for discord integrations
	Application() UserApplication
	ApplicationIsSet() bool
	SetApplication(UserApplication) Integration
	UnsetApplication() Integration
}
