package dio

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/internal/types"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type CreateGuildRequest interface {
	json.Marshaler
	lib.Validatable

	// Name returns the current value of this request's `name` field.
	//
	// The `name` field contains the name of the guild (2-100 characters).
	Name() GuildName

	// SetName overwrites the current value of this request's `name` field.
	SetName(GuildName) CreateGuildRequest

	// Region returns the current value of this request's `region` field.
	//
	// The `region` field contains the voice region id.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use RegionIsSet to check if the field is present before use.
	Region() string

	// RegionIsSet returns whether this request's `region` field is currently
	// present.
	RegionIsSet() bool

	// SetRegion overwrites the current value of this request's `region` field.
	SetRegion(string) CreateGuildRequest

	// UnsetRegion removes this request's `region` field.
	UnsetRegion() CreateGuildRequest

	// Icon returns the current value of this request's `icon` field.
	//
	// The `icon` field contains the base64 128x128 image for the guild icon.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use IconIsSet to check if the field is present before use.
	Icon() GuildIcon

	// IconIsSet returns whether this request's `icon` field is currently present.
	IconIsSet() bool

	// SetIcon overwrites the current value of this request's `icon` field.
	SetIcon(GuildIcon) CreateGuildRequest

	// UnsetIcon removes this request's `icon` field.
	UnsetIcon() CreateGuildRequest

	// VerificationLevel returns the current value of this request's
	// `verification_level` field.
	//
	// The `verification_level` field contains the verification level.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use VerificationLevelIsSet to check if the field is present before use.
	VerificationLevel() VerificationLevel

	// VerificationLevelIsSet returns whether this request's `verification_level`
	// field is currently present.
	VerificationLevelIsSet() bool

	// SetVerificationLevel overwrites the current value of this request's
	// `verification_level` field.
	SetVerificationLevel(VerificationLevel) CreateGuildRequest

	// UnsetVerificationLevel removes this request's `verification_level` field.
	UnsetVerificationLevel() CreateGuildRequest

	// DefaultMessageNotifications returns the current value of this request's
	// `default_message_notifications` field.
	//
	// The `default_message_notifications` field contains the default message
	// notification level.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use DefaultMessageNotificationsIsSet to check if the field is present
	// before use.
	DefaultMessageNotifications() MessageNotificationLevel

	// DefaultMessageNotificationsIsSet returns whether this request's
	// `default_message_notifications` field is currently present.
	DefaultMessageNotificationsIsSet() bool

	// SetDefaultMessageNotifications overwrites the current value of this
	// request's `default_message_notifications` field.
	SetDefaultMessageNotifications(MessageNotificationLevel) CreateGuildRequest

	// UnsetDefaultMessageNotifications removes this request's
	// `default_message_notifications` field.
	UnsetDefaultMessageNotifications() CreateGuildRequest

	// ExplicitContentFilter returns the current value of this request's
	// `explicit_content_filter` field.
	//
	// The `explicit_content_filter` field contains the explicit content filter
	// level.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ExplicitContentFilterIsSet to check if the field is present before use.
	ExplicitContentFilter() ExplicitContentFilterLevel

	// ExplicitContentFilterIsSet returns whether this request's
	// `explicit_content_filter` field is currently present.
	ExplicitContentFilterIsSet() bool

	// SetExplicitContentFilter overwrites the current value of this request's
	// `explicit_content_filter` field.
	SetExplicitContentFilter(ExplicitContentFilterLevel) CreateGuildRequest

	// UnsetExplicitContentFilter removes this request's `explicit_content_filter`
	// field.
	UnsetExplicitContentFilter() CreateGuildRequest

	// Roles returns the current value of this request's `roles` field.
	//
	// The `roles` field contains new guild roles.
	//
	// When using the roles parameter, the first member of the array is used to
	// change properties of the guild's @everyone role. If you are trying to
	// bootstrap a guild with additional roles, keep this in mind.
	//
	// When using the roles parameter, the required id field within each role
	// object is an integer placeholder, and will be replaced by the API upon
	// consumption. Its purpose is to allow you to overwrite a role's permissions
	// in a channel when also passing in channels with the channels array.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use RolesIsSet to check if the field is present before use.
	Roles() []Role

	// RolesIsSet returns whether this request's `roles` field is currently
	// present.
	RolesIsSet() bool

	// SetRoles overwrites the current value of this request's `roles` field.
	//
	// When using the roles parameter, the first member of the array is used to
	// change properties of the guild's @everyone role. If you are trying to
	// bootstrap a guild with additional roles, keep this in mind.
	//
	// When using the roles parameter, the required id field within each role
	// object is an integer placeholder, and will be replaced by the API upon
	// consumption. Its purpose is to allow you to overwrite a role's permissions
	// in a channel when also passing in channels with the channels array.
	SetRoles([]Role) CreateGuildRequest

	// UnsetRoles removes this request's `roles` field.
	UnsetRoles() CreateGuildRequest

	// Channels returns the current value of this request's `channels` field.
	//
	// The `channels` field contains the new guild's channels.
	//
	// When using the channels parameter, the position field is ignored, and none
	// of the default channels are created.
	//
	// When using the channels parameter, the id field within each channel object
	// may be set to an integer placeholder, and will be replaced by the API upon
	// consumption. Its purpose is to allow you to create GUILD_CATEGORY channels
	// by setting the parent_id field on any children to the category's id field.
	// Category channels must be listed before any children.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ChannelsIsSet to check if the field is present before use.
	Channels() []Channel

	// ChannelsIsSet returns whether this request's `channels` field is currently
	// present.
	ChannelsIsSet() bool

	// SetChannels overwrites the current value of this request's `channels`
	// field.
	//
	// When using the channels parameter, the position field is ignored, and none
	// of the default channels are created.
	//
	// When using the channels parameter, the id field within each channel object
	// may be set to an integer placeholder, and will be replaced by the API upon
	// consumption. Its purpose is to allow you to create GUILD_CATEGORY channels
	// by setting the parent_id field on any children to the category's id field.
	// Category channels must be listed before any children.
	SetChannels([]Channel) CreateGuildRequest

	// UnsetChannels removes this request's `channels` field.
	UnsetChannels() CreateGuildRequest

	// AFKChannelID returns the current value of this request's `afk_channel_id`
	// field.
	//
	// The `afk_channel_id` field contains the id for afk channel.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use AFKChannelIDIsSet to check if the field is present before use.
	AFKChannelID() uint16

	// AFKChannelIDIsSet returns whether this request's `afk_channel_id` field is
	// currently present.
	AFKChannelIDIsSet() bool

	// SetAFKChannelID overwrites the current value of this request's
	// `afk_channel_id` field.
	SetAFKChannelID(uint16) CreateGuildRequest

	// UnsetAFKChannelID removes this request's `afk_channel_id` field.
	UnsetAFKChannelID() CreateGuildRequest

	// AFKTimeout returns the current value of this request's `afk_timeout` field.
	//
	// The `afk_timeout` field contains the afk timeout in seconds.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use AFKTimeoutIsSet to check if the field is present before use.
	AFKTimeout() uint64

	// AFKTimeoutIsSet returns whether this request's `afk_timeout` field is
	// currently present.
	AFKTimeoutIsSet() bool

	// SetAFKTimeout overwrites the current value of this request's `afk_timeout`
	// field.
	SetAFKTimeout(uint64) CreateGuildRequest

	// UnsetAFKTimeout removes this request's `afk_timeout` field.
	UnsetAFKTimeout() CreateGuildRequest

	// SystemChannelID returns the current value of this request's
	// `system_channel_id` field.
	//
	// The `system_channel_id` field contains the the id of the channel where
	// guild notices such as welcome messages and boost events are posted.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use SystemChannelIDIsSet to check if the field is present before use.
	SystemChannelID() uint16

	// SystemChannelIDIsSet returns whether this request's `system_channel_id`
	// field is currently present.
	SystemChannelIDIsSet() bool

	// SetSystemChannelID overwrites the current value of this request's
	// `system_channel_id` field.
	SetSystemChannelID(uint16) CreateGuildRequest

	// UnsetSystemChannelID removes this request's `system_channel_id` field.
	UnsetSystemChannelID() CreateGuildRequest
}

func NewCreateGuildRequest(validate bool) CreateGuildRequest {
	return &createGuildRequest{validate: validate}
}

type createGuildRequest struct {
	validate bool

	name   GuildName
	region types.OptionalString
	// TODO: optional icon requires imagedata
	vLvl   types.OptionalUint8
	defMes types.OptionalUint8
	expFil types.OptionalUint8
	// TODO: optional roles requires roles
	// TODO: optional channels requires channels
	afkChan types.OptionalUint16
	afkTime types.OptionalUint64
	sysChan types.OptionalUint16
}

func (c *createGuildRequest) IsValid() bool {
	return nil == c.Validate()
}

func (c *createGuildRequest) Validate() error {
	panic("implement me")
}

func (c *createGuildRequest) MarshalJSON() ([]byte, error) {
	panic("implement me")
}

func (c *createGuildRequest) Name() GuildName {
	return c.name
}

func (c *createGuildRequest) SetName(name GuildName) CreateGuildRequest {
	c.name = name
	return c
}

func (c *createGuildRequest) Region() string {
	return c.region.Get()
}

func (c *createGuildRequest) RegionIsSet() bool {
	return c.region.IsSet()
}

func (c *createGuildRequest) SetRegion(s string) CreateGuildRequest {
	c.region.Set(s)
	return c
}

func (c *createGuildRequest) UnsetRegion() CreateGuildRequest {
	c.region.Unset()
	return c
}

func (c *createGuildRequest) Icon() GuildIcon {
	panic("implement me")
}

func (c *createGuildRequest) IconIsSet() bool {
	panic("implement me")
}

func (c *createGuildRequest) SetIcon(icon GuildIcon) CreateGuildRequest {
	panic("implement me")
}

func (c *createGuildRequest) UnsetIcon() CreateGuildRequest {
	panic("implement me")
}

func (c *createGuildRequest) VerificationLevel() VerificationLevel {
	return VerificationLevel(c.vLvl.Get())
}

func (c *createGuildRequest) VerificationLevelIsSet() bool {
	return c.vLvl.IsSet()
}

func (c *createGuildRequest) SetVerificationLevel(level VerificationLevel) CreateGuildRequest {
	c.vLvl.Set(uint8(level))
	return c
}

func (c *createGuildRequest) UnsetVerificationLevel() CreateGuildRequest {
	c.vLvl.Unset()
	return c
}

func (c *createGuildRequest) DefaultMessageNotifications() MessageNotificationLevel {
	return MessageNotificationLevel(c.defMes.Get())
}

func (c *createGuildRequest) DefaultMessageNotificationsIsSet() bool {
	return c.defMes.IsSet()
}

func (c *createGuildRequest) SetDefaultMessageNotifications(level MessageNotificationLevel) CreateGuildRequest {
	c.defMes.Set(uint8(level))
	return c
}

func (c *createGuildRequest) UnsetDefaultMessageNotifications() CreateGuildRequest {
	c.defMes.Unset()
	return c
}

func (c *createGuildRequest) ExplicitContentFilter() ExplicitContentFilterLevel {
	return ExplicitContentFilterLevel(c.expFil.Get())
}

func (c *createGuildRequest) ExplicitContentFilterIsSet() bool {
	return c.expFil.IsSet()
}

func (c *createGuildRequest) SetExplicitContentFilter(level ExplicitContentFilterLevel) CreateGuildRequest {
	c.expFil.Set(uint8(level))
	return c
}

func (c *createGuildRequest) UnsetExplicitContentFilter() CreateGuildRequest {
	c.expFil.Unset()
	return c
}

func (c *createGuildRequest) Roles() []Role {
	panic("implement me")
}

func (c *createGuildRequest) RolesIsSet() bool {
	panic("implement me")
}

func (c *createGuildRequest) SetRoles(roles []Role) CreateGuildRequest {
	panic("implement me")
}

func (c *createGuildRequest) UnsetRoles() CreateGuildRequest {
	panic("implement me")
}

func (c *createGuildRequest) Channels() []Channel {
	panic("implement me")
}

func (c *createGuildRequest) ChannelsIsSet() bool {
	panic("implement me")
}

func (c *createGuildRequest) SetChannels(channels []Channel) CreateGuildRequest {
	panic("implement me")
}

func (c *createGuildRequest) UnsetChannels() CreateGuildRequest {
	panic("implement me")
}

func (c *createGuildRequest) AFKChannelID() uint16 {
	return c.afkChan.Get()
}

func (c *createGuildRequest) AFKChannelIDIsSet() bool {
	return c.afkChan.IsSet()
}

func (c *createGuildRequest) SetAFKChannelID(u uint16) CreateGuildRequest {
	c.afkChan.Set(u)
	return c
}

func (c *createGuildRequest) UnsetAFKChannelID() CreateGuildRequest {
	c.afkChan.Unset()
	return c
}

func (c *createGuildRequest) AFKTimeout() uint64 {
	return c.afkTime.Get()
}

func (c *createGuildRequest) AFKTimeoutIsSet() bool {
	return c.afkTime.IsSet()
}

func (c *createGuildRequest) SetAFKTimeout(u uint64) CreateGuildRequest {
	c.afkTime.Set(u)
	return c
}

func (c *createGuildRequest) UnsetAFKTimeout() CreateGuildRequest {
	c.afkTime.Unset()
	return c
}

func (c *createGuildRequest) SystemChannelID() uint16 {
	return c.sysChan.Get()
}

func (c *createGuildRequest) SystemChannelIDIsSet() bool {
	return c.sysChan.IsSet()
}

func (c *createGuildRequest) SetSystemChannelID(u uint16) CreateGuildRequest {
	c.sysChan.Set(u)
	return c
}

func (c *createGuildRequest) UnsetSystemChannelID() CreateGuildRequest {
	c.sysChan.Unset()
	return c
}
