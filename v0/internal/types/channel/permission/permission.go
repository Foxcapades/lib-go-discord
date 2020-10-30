package permission

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/francoispqt/gojay"
)

func DecodePermission(dec *gojay.Decoder) (discord.Permission, error) {
	var tmp uint32

	if err := dec.DecodeUint32(&tmp); err != nil {
		return 0, err
	}

	return discord.Permission(tmp), nil
}
