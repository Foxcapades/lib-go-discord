package message

import (
	"bytes"
	"github.com/foxcapades/lib-go-discord/v0/internal/types/com"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dmeta"
	"strconv"
	"time"

	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/internal/types/guild"
	"github.com/foxcapades/lib-go-discord/v0/internal/types/user"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils/gj"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"github.com/foxcapades/lib-go-discord/v0/pkg/e"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

func NewMessage() Message {
	return new(message)
}

const (
	msgMaskAbsentReactions uint8 = 1 << iota
	msgMaskAbsentMentionChans
)

const (
	// Base buffer size, not counting field values.
	msgBufferSize = uint32(2 + // {}
		len(serial.KeyID) + 3 + // "":
		len(serial.KeyChannelID) + 4 + // ,"":
		len(serial.KeyGuildID) + 4 + // ,"":
		len(serial.KeyAuthor) + 4 + // ,"":
		len(serial.KeyMember) + 4 + // ,"":
		len(serial.KeyContent) + 4 + // ,"":
		len(serial.KeyTimestamp) + 4 + // ,"":
		len(serial.KeyEditedTimestamp) + 4 + // ,"":
		len(serial.KeyTTS) + 4 + // ,"":
		len(serial.KeyMentionEveryone) + 4 + // ,"":
		len(serial.KeyMentions) + 4 + // ,"":
		len(serial.KeyMentionRoles) + 4 + // ,"":
		len(serial.KeyMentionChannels) + 4 + // ,"":
		len(serial.KeyAttachments) + 4 + // ,"":
		len(serial.KeyEmbeds) + 4 + // ,"":
		len(serial.KeyReactions) + 4 + // ,"":
		len(serial.KeyNonce) + 4 + // ,"":
		len(serial.KeyPinned) + 4 + // ,"":
		len(serial.KeyWebhookID) + 4 + // ,"":
		len(serial.KeyType) + 4 + // ,"":
		len(serial.KeyActivity) + 4 + // ,"":
		len(serial.KeyApplication) + 4 + // ,"":
		len(serial.KeyMessageReference) + 4 + // ,"":
		len(serial.KeyFlags) + 4) // ,"":
)

type message struct {
	metaFlags uint8

	id              Snowflake
	channelID       Snowflake
	guildID         Snowflake
	author          User
	member          GuildMember
	content         MessageContent
	tStamp          time.Time
	editTStamp      *time.Time
	tts             bool
	mentionEveryone bool
	mentions        user.Slice
	mentionRoles    com.SnowflakeSlice
	mentionChannels ChannelMentionSlice
	attachments     AttachmentSlice
	embeds          EmbedSlice
	reactions       ReactionSlice
	nonce           Nonce
	pinned          bool
	webhookID       Snowflake
	kind            MessageType
	activity        MessageActivity
	application     MessageApplication
	msgReference    MessageRef
	flags           *MessageFlag
}

func (m *message) JSONSize() uint32 {
	return msgBufferSize +
		utils.OptionalSize(m.id) +
		utils.OptionalSize(m.channelID) +
		utils.OptionalSize(m.guildID) +
		utils.OptionalSize(m.author) +
		utils.OptionalSize(&m.content) +
		utils.OptionalTimeSize(&m.tStamp) +
		utils.OptionalTimeSize(m.editTStamp) +
		utils.BoolSize(m.tts) +
		utils.BoolSize(m.mentionEveryone) +
		utils.OptionalSize(m.mentions) +
		utils.OptionalSize(m.mentionRoles) +
		utils.OptionalSize(m.mentionChannels)
}

func (m *message) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, m.JSONSize()))
	enc := gojay.BorrowEncoder(buf)
	defer enc.Release()

	if err := enc.EncodeObject(m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (m *message) UnmarshalJSON(in []byte) error {
	dec := gojay.BorrowDecoder(bytes.NewBuffer(in))
	defer dec.Release()

	return dec.DecodeObject(m)
}

func (m *message) MarshalJSONObject(enc *gojay.Encoder) {
	tmp := gj.NewEncWrapper(enc).
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
		AddSlice(serial.KeyMentionRoles, &m.mentionRoles).
		AddOptionalSlice(
			serial.KeyMentionChannels,
			m.mentionChannels,
			m.metaFlags&msgMaskAbsentMentionChans > 0,
		).
		AddSlice(serial.KeyAttachments, m.attachments).
		AddSlice(serial.KeyEmbeds, m.embeds).
		AddOptionalSlice(
			serial.KeyReactions,
			m.reactions,
			m.metaFlags&msgMaskAbsentReactions > 0,
		)

	if m.nonce != nil {
		if m.nonce.IsInteger() {
			tmp.AddInt64(serial.KeyNonce, m.nonce.IntegerValue())
		} else {
			tmp.AddString(serial.KeyNonce, m.nonce.StringValue())
		}
	}

	tmp.AddBool(serial.KeyPinned, m.pinned).
		AddOptionalSnowflake(serial.KeyWebhookID, m.webhookID).
		AddUint8(serial.KeyType, uint8(m.kind)).
		AddOptionalObject(serial.KeyActivity, m.activity).
		AddOptionalObject(serial.KeyApplication, m.application).
		AddOptionalObject(serial.KeyMessageReference, m.msgReference).
		AddOptionalUint8(serial.KeyFlags, (*uint8)(m.flags))
}

func (m *message) IsNil() bool {
	return m == nil
}

func (m *message) UnmarshalJSONObject(dec *gojay.Decoder, k string) (err error) {
	switch serial.Key(k) {
	case serial.KeyID:
		m.id, err = com.DecodeSnowflake(dec)
	case serial.KeyChannelID:
		m.channelID, err = com.DecodeSnowflake(dec)
	case serial.KeyGuildID:
		m.guildID, err = com.DecodeSnowflake(dec)
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
		m.webhookID, err = com.DecodeSnowflake(dec)
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

	valids := map[serial.Key]dmeta.Validatable{
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

	if out.Len() > 0 {
		return out
	}

	return nil
}

func (m *message) ID() Snowflake {
	return m.id
}

func (m *message) SetID(id Snowflake) Message {
	if id == nil {
		panic(e.ErrSetNil)
	}

	m.id = id

	return m
}

func (m *message) ChannelID() Snowflake {
	return m.channelID
}

func (m *message) SetChannelID(id Snowflake) Message {
	if id == nil {
		panic(e.ErrSetNil)
	}

	m.channelID = id

	return m
}

func (m *message) GuildID() Snowflake {
	if m.guildID == nil {
		panic(e.ErrGetUnset)
	}

	return m.guildID
}

func (m *message) GuildIDIsSet() bool {
	return m.guildID != nil
}

func (m *message) SetGuildID(id Snowflake) Message {
	if id == nil {
		panic(e.ErrSetNil)
	}

	m.guildID = id

	return m
}

func (m *message) UnsetGuildID() Message {
	m.guildID = nil
	return m
}

func (m *message) Author() User {
	return m.author
}

func (m *message) SetAuthor(user User) Message {
	if user == nil {
		panic(e.ErrSetNil)
	}

	m.author = user

	return m
}

func (m *message) Member() GuildMember {
	if m.member == nil {
		panic(e.ErrGetUnset)
	}

	return m.member
}

func (m *message) MemberIsSet() bool {
	return m.member != nil
}

func (m *message) SetMember(member GuildMember) Message {
	if member == nil {
		panic(e.ErrSetNil)
	}

	m.member = member

	return m
}

func (m *message) UnsetMember() Message {
	m.member = nil
	return m
}

func (m *message) Content() MessageContent {
	return m.content
}

func (m *message) SetContent(content MessageContent) Message {
	m.content = content
	return m
}

func (m *message) Timestamp() time.Time {
	return m.tStamp
}

func (m *message) SetTimestamp(t time.Time) Message {
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

func (m *message) SetEditedTimestamp(t time.Time) Message {
	m.editTStamp = &t
	return m
}

func (m *message) SetNullEditedTimestamp() Message {
	m.editTStamp = nil
	return m
}

func (m *message) TTS() bool {
	return m.tts
}

func (m *message) SetTTS(b bool) Message {
	m.tts = b
	return m
}

func (m *message) MentionEveryone() bool {
	return m.mentionEveryone
}

func (m *message) SetMentionEveryone(b bool) Message {
	m.mentionEveryone = b
	return m
}

func (m *message) Mentions() []User {
	return m.mentions
}

func (m *message) SetMentions(users []User) Message {
	m.mentions = users
	return m
}

func (m *message) MentionRoles() []Snowflake {
	return m.mentionRoles
}

func (m *message) SetMentionRoles(roles []Snowflake) Message {
	m.mentionRoles = roles
	return m
}

func (m *message) MentionChannels() []ChannelMention {
	if !m.MentionChannelsIsSet() {
		panic(e.ErrGetUnset)
	}

	return m.mentionChannels
}

func (m *message) MentionChannelsIsSet() bool {
	return m.metaFlags&msgMaskAbsentMentionChans == 0
}

func (m *message) SetMentionChannels(mentions []ChannelMention) Message {
	m.mentionChannels = mentions
	m.metaFlags &= ^msgMaskAbsentMentionChans

	return m
}

func (m *message) UnsetMentionChannels() Message {
	m.mentionChannels = nil
	m.metaFlags |= msgMaskAbsentMentionChans

	return m
}

func (m *message) Attachments() []MessageAttachment {
	return m.attachments
}

func (m *message) SetAttachments(attachments []MessageAttachment) Message {
	m.attachments = attachments
	return m
}

func (m *message) Embeds() []MessageEmbed {
	return m.embeds
}

func (m *message) SetEmbeds(embeds []MessageEmbed) Message {
	m.embeds = embeds
	return m
}

func (m *message) Reactions() []MessageReaction {
	if m.metaFlags&msgMaskAbsentReactions > 0 {
		panic(e.ErrGetUnset)
	}

	return m.reactions
}

func (m *message) ReactionsIsSet() bool {
	return m.metaFlags&msgMaskAbsentReactions == 0
}

func (m *message) SetReactions(reactions []MessageReaction) Message {
	m.reactions = reactions
	m.metaFlags &= ^msgMaskAbsentReactions

	return m
}

func (m *message) UnsetReactions() Message {
	m.reactions = nil
	m.metaFlags |= msgMaskAbsentReactions

	return m
}

func (m *message) Nonce() Nonce {
	if m.nonce == nil {
		panic(e.ErrGetUnset)
	}

	return m.nonce
}

func (m *message) NonceIsSet() bool {
	return m.nonce != nil
}

func (m *message) SetNonce(nonce Nonce) Message {
	if nonce == nil {
		panic(e.ErrSetNil)
	}

	m.nonce = nonce

	return m
}

func (m *message) UnsetNonce() Message {
	m.nonce = nil
	return m
}

func (m *message) Pinned() bool {
	return m.pinned
}

func (m *message) SetPinned(b bool) Message {
	m.pinned = b
	return m
}

func (m *message) WebhookID() Snowflake {
	if m.webhookID == nil {
		panic(e.ErrGetUnset)
	}

	return m.webhookID
}

func (m *message) WebhookIDIsSet() bool {
	return m.webhookID != nil
}

func (m *message) SetWebhookID(id Snowflake) Message {
	if id == nil {
		panic(e.ErrSetNil)
	}

	m.webhookID = id

	return m
}

func (m *message) UnsetWebhookID() Message {
	m.webhookID = nil
	return m
}

func (m *message) Type() MessageType {
	return m.kind
}

func (m *message) SetType(messageType MessageType) Message {
	m.kind = messageType
	return m
}

func (m *message) Activity() MessageActivity {
	if m.activity == nil {
		panic(e.ErrGetUnset)
	}

	return m.activity
}

func (m *message) ActivityIsSet() bool {
	return m.activity != nil
}

func (m *message) SetActivity(activity MessageActivity) Message {
	if activity == nil {
		panic(e.ErrSetNil)
	}

	m.activity = activity

	return m
}

func (m *message) UnsetActivity() Message {
	m.activity = nil
	return m
}

func (m *message) Application() MessageApplication {
	if m.application == nil {
		panic(e.ErrGetUnset)
	}

	return m.application
}

func (m *message) ApplicationIsSet() bool {
	return m.application != nil
}

func (m *message) SetApplication(application MessageApplication) Message {
	if application == nil {
		panic(e.ErrSetNil)
	}

	m.application = application

	return m
}

func (m *message) UnsetApplication() Message {
	m.application = nil
	return m
}

func (m *message) MessageReference() MessageRef {
	if m.msgReference == nil {
		panic(e.ErrGetUnset)
	}

	return m.msgReference
}

func (m *message) MessageReferenceIsSet() bool {
	return m.msgReference != nil
}

func (m *message) SetMessageReference(ref MessageRef) Message {
	if ref == nil {
		panic(e.ErrSetNil)
	}

	m.msgReference = ref

	return m
}

func (m *message) UnsetMessageReference() Message {
	m.msgReference = nil
	return m
}

func (m *message) Flags() MessageFlag {
	if m.flags == nil {
		panic(e.ErrGetUnset)
	}

	return *m.flags
}

func (m *message) FlagsIsSet() bool {
	return m.flags != nil
}

func (m *message) SetFlags(flag MessageFlag) Message {
	m.flags = &flag
	return m
}

func (m *message) UnsetFlags() Message {
	m.flags = nil
	return m
}

func (m *message) AddFlag(flag MessageFlag) Message {
	if m.flags == nil {
		m.flags = &flag
	} else {
		*m.flags |= flag
	}

	return m
}

func (m *message) RemoveFlag(flag MessageFlag) Message {
	if m.flags != nil {
		*m.flags &= ^flag
	}

	return m
}

func (m *message) FlagsContains(flag MessageFlag) bool {
	if m.flags != nil {
		return *m.flags&flag == flag
	}

	return false
}
