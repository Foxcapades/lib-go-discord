package permission

import (
	"bytes"

	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/internal/types/com"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils/gj"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"github.com/foxcapades/lib-go-discord/v0/pkg/e"
)

const (
	overwriteBaseSize = uint32(2 +   // {}
		len(serial.KeyID) + 3 +        // "":
		len(serial.KeyType) + 4 +      // ,"":
		len(serial.KeyAllow) + 4 +     // ,"":
		len(serial.KeyDeny) + 4)       // ,"":
)

type overwrite struct {
	id    discord.Snowflake
	kind  discord.OverwriteType
	allow discord.Permission
	deny  discord.Permission
}

func (o *overwrite) JSONSize() int {
	return overwriteBaseSize +
		utils.OptionalSize(o.id) +
		utils.OptionalSize(o.kind) +
		utils.OptionalSize(o.allow) +
		utils.OptionalSize(o.deny)
}

func (o *overwrite) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, o.JSONSize()))
	enc := gojay.BorrowEncoder(buf)
	defer enc.Release()

	if err := enc.EncodeObject(o); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (o *overwrite) UnmarshalJSON(in []byte) error {
	dec := gojay.BorrowDecoder(bytes.NewBuffer(in))
	defer dec.Release()

	return dec.DecodeObject(o)
}

func (o *overwrite) MarshalJSONObject(enc *gojay.Encoder) {
	gj.NewEncWrapper(enc).
		AddSnowflake(serial.KeyID, o.id).
		AddUint8(serial.KeyType, uint8(o.kind)).
		AddUint32(serial.KeyAllow, uint32(o.allow)).
		AddUint32(serial.KeyDeny, uint32(o.deny))
}

func (o *overwrite) IsNil() bool {
	return o == nil
}

func (o *overwrite) UnmarshalJSONObject(
	dec *gojay.Decoder,
	key string,
) (err error) {
	switch serial.Key(key) {
	case serial.KeyID:
		o.id, err = com.DecodeSnowflake(dec)
	case serial.KeyType:
		o.kind, err = DecodeType(dec)
	case serial.KeyAllow:
		o.allow, err = DecodePermission(dec)
	case serial.KeyDeny:
		o.deny, err = DecodePermission(dec)
	}

	return
}

func (o *overwrite) NKeys() int {
	return 0
}

func (o *overwrite) IsValid() bool {
	return nil == o.Validate()
}

func (o *overwrite) Validate() error {
	out := lib.NewValidationErrorSet()

	out.AppendRawKeyedError(serial.KeyID, o.id.Validate())
	out.AppendRawKeyedError(serial.KeyType, o.kind.Validate())
	out.AppendRawKeyedError(serial.KeyAllow, o.allow.Validate())
	out.AppendRawKeyedError(serial.KeyDeny, o.deny.Validate())

	if out.Len() > 0 {
		return out
	}

	return nil
}

func (o *overwrite) ID() discord.Snowflake {
	return o.id
}

func (o *overwrite) SetID(id discord.Snowflake) discord.PermissionOverwrite {
	if id == nil {
		panic(e.ErrSetNil)
	}

	o.id = id

	return o
}

func (o *overwrite) Type() discord.OverwriteType {
	return o.kind
}

func (o *overwrite) SetType(t discord.OverwriteType) discord.PermissionOverwrite {
	o.kind = t
	return o
}

func (o *overwrite) Allow() discord.Permission {
	return o.allow
}

func (o *overwrite) SetAllow(p discord.Permission) discord.PermissionOverwrite {
	o.allow = p
	return o
}

func (o *overwrite) Deny() discord.Permission {
	return o.deny
}

func (o *overwrite) SetDeny(p discord.Permission) discord.PermissionOverwrite {
	o.deny = p
	return o
}
