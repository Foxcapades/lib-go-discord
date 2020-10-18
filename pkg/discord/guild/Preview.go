package guild

import (
	"github.com/foxcapades/lib-go-discord/pkg/discord/comm"
	"github.com/foxcapades/lib-go-discord/pkg/dlib"
)

type Preview interface {
	// guild id
	ID() dlib.Snowflake
	SetID(id dlib.Snowflake) Preview

	// guild name (2-100 characters)
	Name() Name
	SetName(name Name) Preview

	// icon hash
	Icon() comm.ImageHash
	IconIsNull() bool
	SetIcon(hash comm.ImageHash) Preview
	SetNullIcon() Preview

	// splash hash
	Splash() comm.ImageHash
	SplashIsNull() bool
	SetSplash(hash comm.ImageHash) Preview
	SetNullSplash() Preview

	// splash hash
	DiscoverySplash() comm.ImageHash
	DiscoverySplashIsNull() bool
	SetDiscoverySplash(hash comm.ImageHash) Preview
	SetNullDiscoverySplash() Preview

	// custom guild emojis
	Emojis() []comm.Emoji
	SetEmojis([]comm.Emoji) Preview

	// enabled guild features
	Features() []Feature
	SetFeatures([]Feature) Preview

	// approximate number of members in this guild
	ApproximateMemberCount() uint32
	SetApproximateMemberCount(count uint32) Preview

	// approximate number of online members in this guild
	ApproximatePresenceCount() uint32
	SetApproximatePresenceCount(count uint32) Preview

	// the description for the guild
	Description() string
	DescriptionIsNull() bool
	SetDescription(desc string) Preview
	SetNullDescription() Preview
}
