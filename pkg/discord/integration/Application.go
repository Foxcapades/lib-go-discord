package integration

import (
	"github.com/foxcapades/lib-go-discord/pkg/discord/comm"
	"github.com/foxcapades/lib-go-discord/pkg/discord/user"
	"github.com/foxcapades/lib-go-discord/pkg/dlib"
)

type Application interface {
	// the id of the app
	ID() dlib.Snowflake
	SetID(id dlib.Snowflake) Application

	// the name of the app
	Name() string
	SetName(name string) Application

	// the icon hash of the app
	Icon() comm.ImageHash
	IconIsNull() bool
	SetIcon(hash comm.ImageHash) Application
	SetNullIcon() Application

	// the description of the app
	Description() string
	SetDescription(desc string) Application

	// the description of the app
	Summary() string
	SetSummary(sum string) Application

	// the bot associated with this application
	Bot() user.User
	BotIsSet() bool
	SetBot(user user.User) Application
	UnsetBot() Application
}
