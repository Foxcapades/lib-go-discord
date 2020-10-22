package discord

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/comm"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)

type GuildPreview interface {
	// guild id
	ID() dlib.Snowflake
	SetID(id dlib.Snowflake) GuildPreview

	// guild name (2-100 characters)
	Name() GuildName
	SetName(name GuildName) GuildPreview

	// icon hash
	Icon() comm.ImageHash
	IconIsNull() bool
	SetIcon(hash comm.ImageHash) GuildPreview
	SetNullIcon() GuildPreview

	// splash hash
	Splash() comm.ImageHash
	SplashIsNull() bool
	SetSplash(hash comm.ImageHash) GuildPreview
	SetNullSplash() GuildPreview

	// splash hash
	DiscoverySplash() comm.ImageHash
	DiscoverySplashIsNull() bool
	SetDiscoverySplash(hash comm.ImageHash) GuildPreview
	SetNullDiscoverySplash() GuildPreview

	// custom guild emojis
	Emojis() []Emoji
	SetEmojis([]Emoji) GuildPreview

	// enabled guild features
	Features() []GuildFeature
	SetFeatures([]GuildFeature) GuildPreview

	// approximate number of members in this guild
	ApproximateMemberCount() uint32
	SetApproximateMemberCount(count uint32) GuildPreview

	// approximate number of online members in this guild
	ApproximatePresenceCount() uint32
	SetApproximatePresenceCount(count uint32) GuildPreview

	// the description for the guild
	Description() string
	DescriptionIsNull() bool
	SetDescription(desc string) GuildPreview
	SetNullDescription() GuildPreview
}
