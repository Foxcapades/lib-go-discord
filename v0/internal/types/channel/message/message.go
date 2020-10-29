package message

import (
	"strconv"
	"time"

	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/internal/types"
	"github.com/foxcapades/lib-go-discord/v0/internal/types/guild"
	"github.com/foxcapades/lib-go-discord/v0/internal/types/user"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils/gj"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"github.com/foxcapades/lib-go-discord/v0/pkg/e"
)

func NewMessage() discord.Message {
	return new(message)
}

const (
	msgMaskAbsentReactions uint8 = 1 << iota
	msgMaskAbsentMentionChans
)

// TODO: Slices can't be handled with a simple nil check.  There will need to be
//       some other mechanism for tracking absent slices.  See channel.

type message struct {
	metaFlags uint8

	id              discord.Snowflake
	channelID       discord.Snowflake
	guildID         discord.Snowflake
	author          discord.User
	member          discord.GuildMember
	content         discord.MessageContent
	tStamp          time.Time
	editTStamp      *time.Time
	tts             bool
	mentionEveryone bool
	mentions        user.Slice
	mentionRoles    types.SnowflakeSlice
	mentionChannels ChannelMentionSlice
	attachments     AttachmentSlice
	embeds          EmbedSlice
	reactions       ReactionSlice
	nonce           discord.Nonce
	pinned          bool
	webhookID       discord.Snowflake
	kind            discord.MessageType
	activity        discord.MessageActivity
	application     discord.MessageApplication
	msgReference    discord.MessageRef
	flags           *discord.MessageFlag
}

func (m *message) MarshalJSON() ([]byte, error) {
	panic("implement me")
}

func (m *message) UnmarshalJSON(in []byte) error {
	panic("implement me")
}

func (m *message) MarshalJSONObject(enc *gojay.Encoder) {
	gj.NewEncWrapper(enc).
		AddSnowflake(serial.KeyID, m.id).
		AddSnowflake(serial.KeyChannelID, m.channelID).
		AddOptionalSnowflake(serial.KeyGuildID, m.guildID).
		AddObject(serial.KeyAuthor, m.author).
		AddOptionalObject(serial.KeyMember, m.member).
		AddString(serial.KeyContent, string(m.content)).
		AddString(serial.KeyTimestamp, m.tStamp.Format(time.RFC3339Nano)).
		AddNullableString(serial.KeyEditedTimestamp, gj.EncodeTime(m.editTStamp)).
		AddBool(serial.KeyTTS, m.tts).
		AddBool(serial.KeyMentionEveryone, m.mentionEveryone).
		AddSlice(serial.KeyMentions, m.mentions).
		AddSlice(serial.KeyMentionRoles, m.mentionRoles)
}

func (m *message) IsNil() bool {
	return m == nil
}

func (m *message) UnmarshalJSONObject(dec *gojay.Decoder, k string) (err error) {
	switch serial.Key(k) {
	case serial.KeyID:
		m.id, err = types.DecodeSnowflake(dec)
	case serial.KeyChannelID:
		m.channelID, err = types.DecodeSnowflake(dec)
	case serial.KeyGuildID:
		m.guildID, err = types.DecodeSnowflake(dec)
	case serial.KeyAuthor:
		m.author = user.NewUser()
		err = dec.DecodeObject(m.author)
	case serial.KeyMember:
		m.member = guild.NewMember()
		err = dec.DecodeObject(m.member)
	case serial.KeyContent:
		err = dec.DecodeString((*string)(&m.content))
	case serial.KeyTimestamp:
		m.tStamp, err = gj.DecodeReqTime(dec)
	case serial.KeyEditedTimestamp:
		m.editTStamp, err = gj.DecodeTime(dec)
	case serial.KeyTTS:
		err = dec.DecodeBool(&m.tts)
	case serial.KeyMentionEveryone:
		err = dec.DecodeBool(&m.mentionEveryone)
	case serial.KeyMentions:
		err = dec.DecodeArray(&m.mentions)
	case serial.KeyMentionRoles:
		err = dec.DecodeArray(&m.mentionRoles)
	case serial.KeyMentionChannels:
		err = dec.DecodeArray(&m.mentionChannels)
	case serial.KeyAttachments:
		err = dec.DecodeArray(&m.attachments)
	case serial.KeyEmbeds:
		err = dec.DecodeArray(&m.embeds)
	case serial.KeyReactions:
		err = dec.DecodeArray(&m.reactions)
	case serial.KeyNonce:
		m.nonce, err = DecodeNonce(dec)
	case serial.KeyPinned:
		err = dec.DecodeBool(&m.pinned)
	case serial.KeyWebhookID:
		m.webhookID, err = types.DecodeSnowflake(dec)
	case serial.KeyType:
		err = dec.DecodeUint8((*uint8)(&m.kind))
	case serial.KeyActivity:
		err = dec.DecodeObject(m.activity)
	case serial.KeyApplication:
		err = dec.DecodeObject(m.application)
	case serial.KeyMessageReference:
		err = dec.DecodeObject(m.msgReference)
	case serial.KeyFlags:
		err = dec.DecodeUint8((*uint8)(m.flags))
	}

	return
}

func (m *message) NKeys() int {
	return 0
}

func (m *message) IsValid() bool {
	return nil == m.Validate()
}

func (m *message) Validate() error {
	out := lib.NewValidationErrorSet()

	if m.id == nil {
		out.AppendKeyedError(serial.KeyID, "message id is required")
	}
	if m.channelID == nil {
		out.AppendKeyedError(serial.KeyChannelID, "message channel id is required")
	}
	if m.author == nil {
		out.AppendKeyedError(serial.KeyAuthor, "message author is required")
	}

	valids := map[serial.Key]lib.Validatable{
		serial.KeyID:               m.id,
		serial.KeyChannelID:        m.channelID,
		serial.KeyGuildID:          m.guildID,
		serial.KeyAuthor:           m.author,
		serial.KeyMember:           m.member,
		serial.KeyContent:          &m.content,
		serial.KeyWebhookID:        m.webhookID,
		serial.KeyType:             &m.kind,
		serial.KeyActivity:         m.activity,
		serial.KeyApplication:      m.application,
		serial.KeyMessageReference: m.msgReference,
		serial.KeyFlags:            m.flags,
	}

	for k, v := range valids {
		if v != nil {
			out.AppendRawKeyedError(k, v.Validate())
		}
	}

	for i := range m.mentions {
		if err := m.mentions[i].Validate(); err != nil {
			out.AppendRawKeyedError(
				serial.KeyMentions+
					"["+serial.Key(strconv.FormatInt(int64(i), 10))+"]",
				err,
			)
		}
	}

	for i := range m.mentionRoles {
		if err := m.mentionRoles[i].Validate(); err != nil {
			out.AppendRawKeyedError(
				serial.KeyMentionRoles+
					"["+serial.Key(strconv.FormatInt(int64(i), 10))+"]",
				err,
			)
		}
	}

	for i := range m.mentionChannels {
		if err := m.mentionChannels[i].Validate(); err != nil {
			out.AppendRawKeyedError(
				serial.KeyMentionChannels+
					"["+serial.Key(strconv.FormatInt(int64(i), 10))+"]",
				err,
			)
		}
	}

	for i := range m.attachments {
		if err := m.attachments[i].Validate(); err != nil {
			out.AppendRawKeyedError(
				serial.KeyAttachments+
					"["+serial.Key(strconv.FormatInt(int64(i), 10))+"]",
				err,
			)
		}
	}

	for i := range m.embeds {
		if err := m.embeds[i].Validate(); err != nil {
			out.AppendRawKeyedError(
				serial.KeyEmbeds+
					"["+serial.Key(strconv.FormatInt(int64(i), 10))+"]",
				err,
			)
		}
	}

	for i := range m.reactions {
		if err := m.reactions[i].Validate(); err != nil {
			out.AppendRawKeyedError(
				serial.KeyReactions+
					"["+serial.Key(strconv.FormatInt(int64(i), 10))+"]",
				err,
			)
		}
	}

	if out.GetSize() > 0 {
		return out
	}

	return nil
}

func (m *message) ID() discord.Snowflake {
	return m.id
}

func (m *message) SetID(id discord.Snowflake) discord.Message {
	if id == nil {
		panic(e.ErrSetNil)
	}

	m.id = id

	return m
}

func (m *message) ChannelID() discord.Snowflake {
	return m.channelID
}

func (m *message) SetChannelID(id discord.Snowflake) discord.Message {
	if id == nil {
		panic(e.ErrSetNil)
	}

	m.channelID = id

	return m
}

func (m *message) GuildID() discord.Snowflake {
	if m.guildID == nil {
		panic(e.ErrGetUnset)
	}

	return m.guildID
}

func (m *message) GuildIDIsSet() bool {
	return m.guildID != nil
}

func (m *message) SetGuildID(id discord.Snowflake) discord.Message {
	if id == nil {
		panic(e.ErrSetNil)
	}

	m.guildID = id

	return m
}

func (m *message) UnsetGuildID() discord.Message {
	m.guildID = nil
	return m
}

func (m *message) Author() discord.User {
	return m.author
}

func (m *message) SetAuthor(user discord.User) discord.Message {
	if user == nil {
		panic(e.ErrSetNil)
	}

	m.author = user

	return m
}

func (m *message) Member() discord.GuildMember {
	if m.member == nil {
		panic(e.ErrGetUnset)
	}

	return m.member
}

func (m *message) MemberIsSet() bool {
	return m.member != nil
}

func (m *message) SetMember(member discord.GuildMember) discord.Message {
	if member == nil {
		panic(e.ErrSetNil)
	}

	m.member = member

	return m
}

func (m *message) UnsetMember() discord.Message {
	m.member = nil
	return m
}

func (m *message) Content() discord.MessageContent {
	return m.content
}

func (m *message) SetContent(content discord.MessageContent) discord.Message {
	m.content = content
	return m
}

func (m *message) Timestamp() time.Time {
	return m.tStamp
}

func (m *message) SetTimestamp(t time.Time) discord.Message {
	m.tStamp = t
	return m
}

func (m *message) EditedTimestamp() time.Time {
	if m.editTStamp == nil {
		panic(e.ErrGetNull)
	}

	return *m.editTStamp
}

func (m *message) EditedTimestampIsNull() bool {
	return m.editTStamp == nil
}

func (m *message) SetEditedTimestamp(t time.Time) discord.Message {
	m.editTStamp = &t
	return m
}

func (m *message) SetNullEditedTimestamp() discord.Message {
	m.editTStamp = nil
	return m
}

func (m *message) TTS() bool {
	return m.tts
}

func (m *message) SetTTS(b bool) discord.Message {
	m.tts = b
	return m
}

func (m *message) MentionEveryone() bool {
	return m.mentionEveryone
}

func (m *message) SetMentionEveryone(b bool) discord.Message {
	m.mentionEveryone = b
	return m
}

func (m *message) Mentions() []discord.User {
	return m.mentions
}

func (m *message) SetMentions(users []discord.User) discord.Message {
	m.mentions = users
	return m
}

func (m *message) MentionRoles() []discord.Snowflake {
	return m.mentionRoles
}

func (m *message) SetMentionRoles(roles []discord.Snowflake) discord.Message {
	m.mentionRoles = roles
	return m
}

func (m *message) MentionChannels() []discord.ChannelMention {
	if !m.MentionChannelsIsSet() {
		panic(e.ErrGetUnset)
	}

	return m.mentionChannels
}

func (m *message) MentionChannelsIsSet() bool {
	return m.metaFlags&msgMaskAbsentMentionChans == 0
}

func (m *message) SetMentionChannels(mentions []discord.ChannelMention) discord.Message {
	m.mentionChannels = mentions
	m.metaFlags &= ^msgMaskAbsentMentionChans

	return m
}

func (m *message) UnsetMentionChannels() discord.Message {
	m.mentionChannels = nil
	m.metaFlags |= msgMaskAbsentMentionChans

	return m
}

func (m *message) Attachments() []discord.MessageAttachment {
	return m.attachments
}

func (m *message) SetAttachments(attachments []discord.MessageAttachment) discord.Message {
	m.attachments = attachments
	return m
}

func (m *message) Embeds() []discord.MessageEmbed {
	return m.embeds
}

func (m *message) SetEmbeds(embeds []discord.MessageEmbed) discord.Message {
	m.embeds = embeds
	return m
}

func (m *message) Reactions() []discord.MessageReaction {
	if m.metaFlags&msgMaskAbsentReactions > 0 {
		panic(e.ErrGetUnset)
	}

	return m.reactions
}

func (m *message) ReactionsIsSet() bool {
	return m.metaFlags&msgMaskAbsentReactions == 0
}

func (m *message) SetReactions(reactions []discord.MessageReaction) discord.Message {
	m.reactions = reactions
	m.metaFlags &= ^msgMaskAbsentReactions

	return m
}

func (m *message) UnsetReactions() discord.Message {
	m.reactions = nil
	m.metaFlags |= msgMaskAbsentReactions

	return m
}

func (m *message) Nonce() discord.Nonce {
	if m.nonce == nil {
		panic(e.ErrGetUnset)
	}

	return m.nonce
}

func (m *message) NonceIsSet() bool {
	return m.nonce != nil
}

func (m *message) SetNonce(nonce discord.Nonce) discord.Message {
	if nonce == nil {
		panic(e.ErrSetNil)
	}

	m.nonce = nonce

	return m
}

func (m *message) UnsetNonce() discord.Message {
	m.nonce = nil
	return m
}

func (m *message) Pinned() bool {
	return m.pinned
}

func (m *message) SetPinned(b bool) discord.Message {
	m.pinned = b
	return m
}

func (m *message) WebhookID() discord.Snowflake {
	if m.webhookID == nil {
		panic(e.ErrGetUnset)
	}

	return m.webhookID
}

func (m *message) WebhookIDIsSet() bool {
	return m.webhookID != nil
}

func (m *message) SetWebhookID(id discord.Snowflake) discord.Message {
	if id == nil {
		panic(e.ErrSetNil)
	}

	m.webhookID = id

	return m
}

func (m *message) UnsetWebhookID() discord.Message {
	m.webhookID = nil
	return m
}

func (m *message) Type() discord.MessageType {
	return m.kind
}

func (m *message) SetType(messageType discord.MessageType) discord.Message {
	m.kind = messageType
	return m
}

func (m *message) Activity() discord.MessageActivity {
	if m.activity == nil {
		panic(e.ErrGetUnset)
	}

	return m.activity
}

func (m *message) ActivityIsSet() bool {
	return m.activity != nil
}

func (m *message) SetActivity(activity discord.MessageActivity) discord.Message {
	if activity == nil {
		panic(e.ErrSetNil)
	}

	m.activity = activity

	return m
}

func (m *message) UnsetActivity() discord.Message {
	m.activity = nil
	return m
}

func (m *message) Application() discord.MessageApplication {
	if m.application == nil {
		panic(e.ErrGetUnset)
	}

	return m.application
}

func (m *message) ApplicationIsSet() bool {
	return m.application != nil
}

func (m *message) SetApplication(application discord.MessageApplication) discord.Message {
	if application == nil {
		panic(e.ErrSetNil)
	}

	m.application = application

	return m
}

func (m *message) UnsetApplication() discord.Message {
	m.application = nil
	return m
}

func (m *message) MessageReference() discord.MessageRef {
	if m.msgReference == nil {
		panic(e.ErrGetUnset)
	}

	return m.msgReference
}

func (m *message) MessageReferenceIsSet() bool {
	return m.msgReference != nil
}

func (m *message) SetMessageReference(ref discord.MessageRef) discord.Message {
	if ref == nil {
		panic(e.ErrSetNil)
	}

	m.msgReference = ref

	return m
}

func (m *message) UnsetMessageReference() discord.Message {
	m.msgReference = nil
	return m
}

func (m *message) Flags() discord.MessageFlag {
	if m.flags == nil {
		panic(e.ErrGetUnset)
	}

	return *m.flags
}

func (m *message) FlagsIsSet() bool {
	return m.flags != nil
}

func (m *message) SetFlags(flag discord.MessageFlag) discord.Message {
	m.flags = &flag
	return m
}

func (m *message) UnsetFlags() discord.Message {
	m.flags = nil
	return m
}

func (m *message) AddFlag(flag discord.MessageFlag) discord.Message {
	if m.flags == nil {
		m.flags = &flag
	} else {
		*m.flags |= flag
	}

	return m
}

func (m *message) RemoveFlag(flag discord.MessageFlag) discord.Message {
	if m.flags != nil {
		*m.flags &= ^flag
	}

	return m
}

func (m *message) FlagsContains(flag discord.MessageFlag) bool {
	if m.flags != nil {
		return *m.flags&flag == flag
	}

	return false
}
