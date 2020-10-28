package permission

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/francoispqt/gojay"
)

type overwrite struct{}

func (o *overwrite) MarshalJSON() ([]byte, error) {

}

func (o *overwrite) UnmarshalJSON(bytes []byte) error {

}

func (o *overwrite) MarshalJSONObject(enc *gojay.Encoder) {

}

func (o *overwrite) IsNil() bool {

}

func (o *overwrite) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {

}

func (o *overwrite) NKeys() int {

}

func (o *overwrite) IsValid() bool {

}

func (o *overwrite) Validate() error {

}

func (o *overwrite) ID() discord.Snowflake {

}

func (o *overwrite) SetID(id discord.Snowflake) discord.PermissionOverwrite {

}

func (o *overwrite) Type() discord.Type {

}

func (o *overwrite) SetType(t discord.Type) discord.PermissionOverwrite {

}

func (o *overwrite) Allow() discord.Permission {

}

func (o *overwrite) SetAllow(p discord.Permission) discord.PermissionOverwrite {

}

func (o *overwrite) Deny() discord.Permission {

}

func (o *overwrite) SetDeny(p discord.Permission) discord.PermissionOverwrite {

}
