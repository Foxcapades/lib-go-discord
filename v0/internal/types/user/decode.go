package user

import (
	"github.com/foxcapades/lib-go-discord/v0/internal/utils/gj"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/francoispqt/gojay"
	"strconv"
)

func DecodeDiscriminator(dec *gojay.Decoder) (discord.Discriminator, error) {
	var tmp string

	if err := dec.DecodeString(&tmp); err != nil {
		return 0, err
	}

	if raw, err := strconv.ParseUint(tmp, 10, 16); err != nil {
		return 0, err
	} else {
		return discord.Discriminator(raw), nil
	}
}

func DecodeFlags(dec *gojay.Decoder) (*discord.UserFlag, error) {
	var tmp *uint32

	if err := dec.Uint32Null(&tmp); err != nil {
		return nil, err
	}

	return (*discord.UserFlag)(tmp), nil
}

func DecodePremiumType(dec *gojay.Decoder) (*discord.UserPremiumType, error) {
	tmp, err := gj.DecodeOptUint8(dec)
	return (*discord.UserPremiumType)(tmp), err
}
