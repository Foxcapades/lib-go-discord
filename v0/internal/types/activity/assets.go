package activity

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/francoispqt/gojay"
)

func NewAssets() discord.ActivityAssets {
	return new(assets)
}

type assets struct {
	LgImage *string
	LgText  *string
	SmImage *string
	SmText  *string
}

func (a *assets) MarshalJSONObject(enc *gojay.Encoder) {
	panic("implement me")
}

func (a *assets) IsNil() bool {
	panic("implement me")
}

func (a *assets) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	panic("implement me")
}

func (a *assets) NKeys() int {
	panic("implement me")
}

func (a *assets) MarshalJSON() ([]byte, error) {
	panic("implement me")
}

func (a *assets) UnmarshalJSON(bytes []byte) error {
	panic("implement me")
}

func (a *assets) LargeImage() string {
	panic("implement me")
}

func (a *assets) LargeImageIsSet() bool {
	panic("implement me")
}

func (a *assets) SetLargeImage(s string) discord.ActivityAssets {
	panic("implement me")
}

func (a *assets) UnsetLargeImage() discord.ActivityAssets {
	panic("implement me")
}

func (a *assets) LargeText() string {
	panic("implement me")
}

func (a *assets) LargeTextIsSet() bool {
	panic("implement me")
}

func (a *assets) SetLargeText(s string) discord.ActivityAssets {
	panic("implement me")
}

func (a *assets) UnsetLargeText() discord.ActivityAssets {
	panic("implement me")
}

func (a *assets) SmallImage() string {
	panic("implement me")
}

func (a *assets) SmallImageIsSet() bool {
	panic("implement me")
}

func (a *assets) SetSmallImage(s string) discord.ActivityAssets {
	panic("implement me")
}

func (a *assets) UnsetSmallImage() discord.ActivityAssets {
	panic("implement me")
}

func (a *assets) SmallText() string {
	panic("implement me")
}

func (a *assets) SmallTextIsSet() bool {
	panic("implement me")
}

func (a *assets) SetSmallText(s string) discord.ActivityAssets {
	panic("implement me")
}

func (a *assets) UnsetSmallText() discord.ActivityAssets {
	panic("implement me")
}
