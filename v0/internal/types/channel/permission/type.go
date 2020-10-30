package permission

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/francoispqt/gojay"
)

func DecodeType(dec *gojay.Decoder) (discord.OverwriteType, error) {
	var tmp uint8

	if err := dec.DecodeUint8(&tmp); err != nil {
		return 0, err
	}

	return discord.OverwriteType(tmp), nil
}
