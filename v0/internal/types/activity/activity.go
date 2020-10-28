package activity

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/internal/types"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils/gj"
	"time"

	"github.com/foxcapades/lib-go-discord/v0/internal/utils"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/foxcapades/lib-go-discord/v0/pkg/e"
	"github.com/francoispqt/gojay"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
)

const (
	actNullURL uint8 = 1 << iota
	actNullDetails
	actNullState
	actNullEmoji
)

func NewActivity() discord.Activity {
	return new(activity)
}

//╔══════════════════════════════════════════════════════════════════════════╗//
//║                                                                          ║//
//║   Implementation Type                                                    ║//
//║                                                                          ║//
//╚══════════════════════════════════════════════════════════════════════════╝//

type activity struct {
	// Meta fields
	nulls uint8

	// IO fields
	name       string
	kind       discord.ActivityType
	url        *string
	createdAt  time.Time
	timestamps discord.ActivityTimestamps
	appID      discord.Snowflake
	details    *string
	state      *string
	emoji      discord.ActivityEmoji
	party      discord.ActivityParty
	assets     discord.ActivityAssets
	secrets    discord.ActivitySecrets
	instance   *bool
	flags      *discord.ActivityFlag
}

//╔══════════════════════════════════════════════════════════════════════════╗//
//║                                                                          ║//
//║   GoJay Serialization Methods                                            ║//
//║                                                                          ║//
//╚══════════════════════════════════════════════════════════════════════════╝//

func (a *activity) MarshalJSONObject(enc *gojay.Encoder) {
	gj.NewEncWrapper(enc).
		AddString(KeyName, a.name).
		AddUint8(KeyType, uint8(a.kind)).
		AddTriStateString(KeyURL, a.url, a.has(actNullURL)).
		AddUint64(KeyCreatedAt, *utils.UTCTimeToMillis(&a.createdAt)).
		AddOptionalObject(KeyTimestamps, a.timestamps).
		AddOptionalSnowflake(KeyApplicationID, a.appID).
		AddTriStateString(KeyDetails, a.details, a.has(actNullDetails)).
		AddTriStateString(KeyState, a.state, a.has(actNullState)).
		AddTriStateObject(KeyEmoji, a.emoji, a.has(actNullEmoji)).
		AddOptionalObject(KeyParty, a.party).
		AddOptionalObject(KeyAssets, a.assets).
		AddOptionalObject(KeySecrets, a.secrets).
		AddOptionalBool(KeyInstance, a.instance).
		AddOptionalUint8(KeyInstance, (*uint8)(a.flags))
}

func (a *activity) IsNil() bool {
	return a == nil
}

func (a *activity) UnmarshalJSONObject(d *gojay.Decoder, k string) error {
	switch Key(k) {

	case KeyName:
		return d.DecodeString(&a.name)

	case KeyType:
		var tmp uint8

		if err := d.DecodeUint8(&tmp); err != nil {
			return err
		}

		a.kind = discord.ActivityType(tmp)

	case KeyURL:
		return d.DecodeString(a.url)

	case KeyCreatedAt:
		var tmp uint64

		if err := d.DecodeUint64(&tmp); err != nil {
			return err
		}

		// TODO: confirm that the created at date is in milliseconds
		sec := tmp / 1000
		nan := (tmp % 1000) * 1000000

		a.createdAt = time.Unix(int64(sec), int64(nan))

	case KeyTimestamps:
		a.timestamps = NewTimestamps()
		return d.DecodeObject(a.timestamps)

	case KeyApplicationID:
		var tmp string

		a.appID = types.NewSnowflake()

		if err := d.DecodeString(&tmp); err != nil {
			return err
		}

		if err := a.appID.UnmarshalString(tmp); err != nil {
			return err
		}

	case KeyDetails:
		return d.DecodeString(a.details)

	case KeyState:
		return d.DecodeString(a.state)

	case KeyEmoji:
		a.emoji = NewEmoji()
		return d.DecodeObject(a.emoji)

	case KeyParty:
		a.party = NewParty()
		return d.DecodeObject(a.party)

	case KeyAssets:
		a.assets = NewAssets()
		return d.DecodeObject(a.assets)

	case KeySecrets:
		a.secrets = NewSecrets()
		return d.DecodeObject(a.secrets)

	case KeyInstance:
		return d.DecodeBool(a.instance)

	case KeyFlags:
		var tmp *uint8

		if err := d.DecodeUint8(tmp); err != nil {
			return err
		}

		a.flags = (*discord.ActivityFlag)(tmp)
	}

	return nil
}

func (a *activity) NKeys() int {
	return 0
}

//╔══════════════════════════════════════════════════════════════════════════╗//
//║                                                                          ║//
//║   Standard Library Serialization Methods                                 ║//
//║                                                                          ║//
//╚══════════════════════════════════════════════════════════════════════════╝//

func (a *activity) MarshalJSON() ([]byte, error) {
	return json.Marshal(make(utils.OutBuilder, 14).
		RawValue(KeyName, &a.name).
		RawValue(KeyType, &a.kind).
		TriState(KeyURL, a.has(actNullURL), &a.kind).
		RawValue(KeyCreatedAt, *utils.UTCTimeToMillis(&a.createdAt)).
		Nullable(KeyTimestamps, a.timestamps).
		Nullable(KeyApplicationID, a.appID).
		TriState(KeyDetails, a.has(actNullDetails), a.details).
		TriState(KeyState, a.has(actNullState), a.state).
		TriState(KeyEmoji, a.has(actNullEmoji), a.emoji).
		Optional(KeyParty, a.party).
		Optional(KeyAssets, a.assets).
		Optional(KeySecrets, a.secrets).
		Optional(KeyInstance, a.instance).
		Optional(KeyFlags, a.flags))
}

func (a *activity) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[Key]json.RawMessage, 14)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if val, ok := tmp[KeyName]; ok {
		if err := json.Unmarshal(val, &a.name); err != nil {
			return err
		}
	}

	if val, ok := tmp[KeyType]; ok {
		if err := json.Unmarshal(val, &a.kind); err != nil {
			return err
		}
	}

	if val, ok := tmp[KeyURL]; ok {
		if err := json.Unmarshal(val, &a.url); err != nil {
			return err
		}

		if a.url == nil {
			a.nulls |= actNullURL
		}
	}

	if val, ok := tmp[KeyCreatedAt]; ok {
		var tmp uint64

		if err := json.Unmarshal(val, &tmp); err != nil {
			return err
		}

		// TODO: confirm that the created at date is in milliseconds
		sec := tmp / 1000
		nan := (tmp % 1000) * 1000000

		a.createdAt = time.Unix(int64(sec), int64(nan))
	}

	if val, ok := tmp[KeyTimestamps]; ok {
		a.timestamps = NewTimestamps()
		if err := json.Unmarshal(val, a.timestamps); err != nil {
			return err
		}
	}

	if val, ok := tmp[KeyApplicationID]; ok {
		a.appID = types.NewSnowflake()
		if err := a.appID.UnmarshalJSON(val); err != nil {
			return err
		}
	}

	if val, ok := tmp[KeyDetails]; ok {
		if err := json.Unmarshal(val, a.details); err != nil {
			return err
		}

		if a.details == nil {
			a.nulls |= actNullDetails
		}
	}

	if val, ok := tmp[KeyState]; ok {
		if err := json.Unmarshal(val, a.state); err != nil {
			return err
		}

		if a.state == nil {
			a.nulls |= actNullState
		}
	}

	if val, ok := tmp[KeyEmoji]; ok {
		a.emoji = NewEmoji()

		if err := a.emoji.UnmarshalJSON(val); err != nil {
			return err
		}

		if a.emoji == nil {
			a.nulls |= actNullEmoji
		}
	}

	if val, ok := tmp[KeyParty]; ok {
		a.party = NewParty()
		if err := a.party.UnmarshalJSON(val); err != nil {
			return err
		}
	}

	if val, ok := tmp[KeyAssets]; ok {
		a.assets = NewAssets()
		if err := a.assets.UnmarshalJSON(val); err != nil {
			return err
		}
	}

	if val, ok := tmp[KeySecrets]; ok {
		a.secrets = NewSecrets()
		if err := a.secrets.UnmarshalJSON(val); err != nil {
			return err
		}
	}

	if val, ok := tmp[KeyInstance]; ok {
		if err := json.Unmarshal(val, a.instance); err != nil {
			return err
		}
	}

	if val, ok := tmp[KeyFlags]; ok {
		if err := json.Unmarshal(val, a.flags); err != nil {
			return err
		}
	}

	return nil
}

//╔══════════════════════════════════════════════════════════════════════════╗//
//║                                                                          ║//
//║   Validation Methods                                                     ║//
//║                                                                          ║//
//╚══════════════════════════════════════════════════════════════════════════╝//

func (a *activity) IsValid() bool {
	return nil == a.Validate()
}

func (a *activity) Validate() error {
	errs := lib.NewValidationErrorSet()

	if a.TimestampsIsSet() {
		errs.AppendRawKeyedError(KeyTimestamps, a.timestamps.Validate())
	}

	if a.ApplicationIDIsSet() {
		errs.AppendRawKeyedError(KeyApplicationID, a.appID.Validate())
	}

	if a.EmojiIsReadable() {
		errs.AppendRawKeyedError(KeyEmoji, a.emoji.Validate())
	}

	if a.PartyIsSet() {
		errs.AppendRawKeyedError(KeyParty, a.party.Validate())
	}

	if a.FlagsIsSet() {
		errs.AppendRawKeyedError(KeyFlags, a.flags.Validate())
	}

	if errs.GetSize() > 0 {
		return errs
	}

	return nil
}

//╔══════════════════════════════════════════════════════════════════════════╗//
//║                                                                          ║//
//║   Property Accessors                                                     ║//
//║                                                                          ║//
//╚══════════════════════════════════════════════════════════════════════════╝//

func (a *activity) Name() string {
	return a.name
}

func (a *activity) SetName(s string) discord.Activity {
	a.name = s
	return a
}

func (a *activity) Type() discord.ActivityType {
	return a.kind
}

func (a *activity) SetType(activityType discord.ActivityType) discord.Activity {
	a.kind = activityType
	return a
}

func (a *activity) URL() string {
	if a.url == nil {
		if a.has(actNullURL) {
			panic(e.ErrGetNull)
		} else {
			panic(e.ErrGetUnset)
		}
	}

	return *a.url
}

func (a *activity) URLIsNull() bool {
	return a.url == nil && a.has(actNullURL)
}

func (a *activity) URLIsSet() bool {
	return a.url != nil
}

func (a *activity) URLIsReadable() bool {
	return a.url != nil
}

func (a *activity) SetURL(s string) discord.Activity {
	a.url = &s
	a.nulls &= ^actNullURL

	return a
}

func (a *activity) SetNullURL() discord.Activity {
	a.url = nil
	a.nulls |= actNullURL

	return a
}

func (a *activity) UnsetURL() discord.Activity {
	a.url = nil
	a.nulls &= ^actNullURL

	return a
}

func (a *activity) CreatedAt() time.Time {
	return a.createdAt
}

func (a *activity) SetCreatedAt(time time.Time) discord.Activity {
	a.createdAt = time
	return a
}

func (a *activity) Timestamps() discord.ActivityTimestamps {
	if a.timestamps == nil {
		panic(e.ErrGetUnset)
	}

	return a.timestamps
}

func (a *activity) TimestampsIsSet() bool {
	return a.timestamps != nil
}

func (a *activity) SetTimestamps(timestamps discord.ActivityTimestamps) discord.Activity {
	if timestamps == nil {
		panic(e.ErrSetNil)
	}

	a.timestamps = timestamps

	return a
}

func (a *activity) UnsetTimestamps() discord.Activity {
	a.timestamps = nil
	return a
}

func (a *activity) ApplicationID() discord.Snowflake {
	if a.appID == nil {
		panic(e.ErrGetUnset)
	}

	return a.appID
}

func (a *activity) ApplicationIDIsSet() bool {
	return a.appID != nil
}

func (a *activity) SetApplicationID(snowflake discord.Snowflake) discord.Activity {
	if snowflake == nil {
		panic(e.ErrSetNil)
	}

	a.appID = snowflake

	return a
}

func (a *activity) UnsetApplicationID() discord.Activity {
	a.appID = nil
	return a
}

func (a *activity) Details() string {
	if a.details == nil {
		if a.has(actNullDetails) {
			panic(e.ErrGetNull)
		} else {
			panic(e.ErrGetUnset)
		}
	}

	return *a.details
}

func (a *activity) DetailsIsNull() bool {
	return a.details == nil && a.has(actNullDetails)
}

func (a *activity) DetailsIsSet() bool {
	return a.details != nil
}

func (a *activity) DetailsIsReadable() bool {
	return a.details != nil
}

func (a *activity) SetDetails(s string) discord.Activity {
	a.details = &s
	a.nulls &= ^actNullDetails

	return a
}

func (a *activity) SetNullDetails() discord.Activity {
	a.details = nil
	a.nulls |= actNullDetails

	return a
}

func (a *activity) UnsetDetails() discord.Activity {
	a.details = nil
	a.nulls &= ^actNullDetails

	return a
}

func (a *activity) State() string {
	if a.state == nil {
		if a.has(actNullState) {
			panic(e.ErrGetNull)
		} else {
			panic(e.ErrGetUnset)
		}
	}

	return *a.state
}

func (a *activity) StateIsNull() bool {
	return a.state == nil && a.has(actNullState)
}

func (a *activity) StateIsSet() bool {
	return a.state != nil
}

func (a *activity) StateIsReadable() bool {
	return a.state != nil
}

func (a *activity) SetState(s string) discord.Activity {
	a.state = &s
	a.nulls &= ^actNullState

	return a
}

func (a *activity) SetNullState() discord.Activity {
	a.state = nil
	a.nulls |= actNullState

	return nil
}

func (a *activity) UnsetState() discord.Activity {
	a.state = nil
	a.nulls &= ^actNullState

	return nil
}

func (a *activity) Emoji() discord.ActivityEmoji {
	if a.emoji == nil {
		if a.has(actNullEmoji) {
			panic(e.ErrGetNull)
		} else {
			panic(e.ErrGetUnset)
		}
	}

	return a.emoji
}

func (a *activity) EmojiIsNull() bool {
	return a.emoji == nil && a.has(actNullEmoji)
}

func (a *activity) EmojiIsSet() bool {
	return a.emoji != nil
}

func (a *activity) EmojiIsReadable() bool {
	return a.emoji != nil
}

func (a *activity) SetEmoji(emoji discord.ActivityEmoji) discord.Activity {
	if emoji == nil {
		panic(e.ErrSetNil)
	}

	a.emoji = emoji
	a.nulls &= ^actNullEmoji

	return a
}

func (a *activity) SetNullEmoji() discord.Activity {
	a.emoji = nil
	a.nulls |= actNullEmoji

	return a
}

func (a *activity) UnsetEmoji() discord.Activity {
	a.emoji = nil
	a.nulls &= ^actNullEmoji

	return a
}

func (a *activity) Party() discord.ActivityParty {
	if a.party == nil {
		panic(e.ErrGetUnset)
	}

	return a.party
}

func (a *activity) PartyIsSet() bool {
	return a.party != nil
}

func (a *activity) SetParty(party discord.ActivityParty) discord.Activity {
	if party == nil {
		panic(e.ErrSetNil)
	}

	a.party = party

	return a
}

func (a *activity) UnsetParty() discord.Activity {
	a.party = nil

	return a
}

func (a *activity) Assets() discord.ActivityAssets {
	if a.assets == nil {
		panic(e.ErrGetUnset)
	}

	return a.assets
}

func (a *activity) AssetsIsSet() bool {
	return a.assets == nil
}

func (a *activity) SetAssets(ass discord.ActivityAssets) discord.Activity {
	if ass == nil {
		panic(e.ErrSetNil)
	}

	return a
}

func (a *activity) UnsetAssets() discord.Activity {
	a.assets = nil
	return a
}

func (a *activity) Secrets() discord.ActivitySecrets {
	if a.secrets == nil {
		panic(e.ErrGetUnset)
	}

	return a.secrets
}

func (a *activity) SecretsIsSet() bool {
	return a.secrets != nil
}

func (a *activity) SetSecrets(sec discord.ActivitySecrets) discord.Activity {
	if sec == nil {
		panic(e.ErrSetNil)
	}

	a.secrets = sec

	return a
}

func (a *activity) UnsetSecrets() discord.Activity {
	a.secrets = nil
	return a
}

func (a *activity) Instance() bool {
	if a.instance == nil {
		panic(e.ErrGetUnset)
	}

	return *a.instance
}

func (a *activity) InstanceIsSet() bool {
	return a.instance != nil
}

func (a *activity) SetInstance(b bool) discord.Activity {
	a.instance = &b
	return a
}

func (a *activity) UnsetInstance() discord.Activity {
	a.instance = nil
	return a
}

func (a *activity) Flags() discord.ActivityFlag {
	if a.flags == nil {
		panic(e.ErrGetUnset)
	}

	return *a.flags
}

func (a *activity) FlagsIsSet() bool {
	return a.flags != nil
}

func (a *activity) SetFlags(flag discord.ActivityFlag) discord.Activity {
	a.flags = &flag
	return a
}

func (a *activity) UnsetFlags() discord.Activity {
	a.flags = nil
	return a
}

//╔══════════════════════════════════════════════════════════════════════════╗//
//║                                                                          ║//
//║   Internal Helpers                                                       ║//
//║                                                                          ║//
//╚══════════════════════════════════════════════════════════════════════════╝//

func (a *activity) has(v uint8) bool {
	return a.nulls&v == v
}
