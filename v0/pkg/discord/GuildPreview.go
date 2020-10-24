package discord

type GuildPreview interface {
	// guild id
	ID() Snowflake
	SetID(id Snowflake) GuildPreview

	// guild name (2-100 characters)
	Name() GuildName
	SetName(name GuildName) GuildPreview

	// icon hash
	Icon() ImageHash
	IconIsNull() bool
	SetIcon(hash ImageHash) GuildPreview
	SetNullIcon() GuildPreview

	// splash hash
	Splash() ImageHash
	SplashIsNull() bool
	SetSplash(hash ImageHash) GuildPreview
	SetNullSplash() GuildPreview

	// splash hash
	DiscoverySplash() ImageHash
	DiscoverySplashIsNull() bool
	SetDiscoverySplash(hash ImageHash) GuildPreview
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
