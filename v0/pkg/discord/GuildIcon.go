package discord

import (
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dmeta"
)

var (
	ErrBadIconSize = errors.New("given image is too large to use as a guild" +
		" icon; images must be no larger than 1024x1024")
)

const (
	IconMaxWidth  uint16 = 1024
	IconMaxHeight uint16 = 1024
)

type GuildIcon interface {
	dmeta.Validatable
	ImageData
}
