package discord

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/comm"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)

type UserApplication interface {
	// the id of the app
	ID() dlib.Snowflake
	SetID(id dlib.Snowflake) UserApplication

	// the name of the app
	Name() string
	SetName(name string) UserApplication

	// the icon hash of the app
	Icon() comm.ImageHash
	IconIsNull() bool
	SetIcon(hash comm.ImageHash) UserApplication
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
