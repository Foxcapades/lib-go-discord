package discord

type UserApplication interface {
	// the id of the app
	ID() Snowflake
	SetID(id Snowflake) UserApplication

	// the name of the app
	Name() string
	SetName(name string) UserApplication

	// the icon hash of the app
	Icon() ImageHash
	IconIsNull() bool
	SetIcon(hash ImageHash) UserApplication
	SetNullIcon() UserApplication

	// the description of the app
	Description() string
	SetDescription(desc string) UserApplication

	// the description of the app
	Summary() string
	SetSummary(sum string) UserApplication

	// the bot associated with this application
	Bot() User
	BotIsSet() bool
	SetBot(user User) UserApplication
	UnsetBot() UserApplication
}
