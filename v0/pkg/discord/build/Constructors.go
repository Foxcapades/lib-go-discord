package build

import (
	"image"
	"io"

	"github.com/foxcapades/lib-go-discord/v0/internal/types/activity"
	"github.com/foxcapades/lib-go-discord/v0/internal/types/channel"
	"github.com/foxcapades/lib-go-discord/v0/internal/types/guild"
	"github.com/foxcapades/lib-go-discord/v0/internal/types/user"
	"github.com/foxcapades/lib-go-discord/v0/internal/types/voice"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"

	t "github.com/foxcapades/lib-go-discord/v0/internal/types"
	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

///  A  ////////////////////////////////////////////////////////////////////////

func NewActivity() Activity {
	return activity.NewActivity()
}

func NewActivityAssets() ActivityAssets {
	return activity.NewAssets()
}

func NewActivityEmoji() ActivityEmoji {
	return activity.NewEmoji()
}

func NewActivityParty() ActivityParty {
	return activity.NewParty()
}

func NewActivitySecrets() ActivitySecrets {
	return activity.NewSecrets()
}

func NewActivityTimestamps() ActivityTimestamps {
	return activity.NewTimestamps()
}

///  B  ////////////////////////////////////////////////////////////////////////

func NewBan() Ban {
	return guild.NewBan()
}

///  C  ////////////////////////////////////////////////////////////////////////

func NewChannel() Channel {
	return channel.NewChannel()
}

///  I  ////////////////////////////////////////////////////////////////////////

// RequireNewIcon constructs a new Guild Icon from the given input file or
// stream.
//
// If this function encounters any error, including ErrBadIconSize, it will
// panic.
//
// See NewIcon() for more information.
func RequireNewIcon(in io.ReadSeeker) GuildIcon {
	if out, err := NewIcon(in); err != nil {
		panic(err)
	} else {
		return out
	}
}

// NewIcon constructs a new Guild Icon from the given input file or stream.
//
// If an IO error occurs while attempting to work with the given stream this
// function will return `nil, error`.
//
// If the given image stream is not an image that is a gif, png, or jpeg this
// function will return `nil, ErrBadIconFormat`.
//
// If the given image stream is of an image with a width or height greater than
// the allowed dimensions as defined by the Discord API docs, then the new Icon
// instance will be returned alongside an ErrBadIconSize error.  This error may
// be treated as a warning and ignored, however the Discord API may return an
// error itself if this icon is sent to that API.
func NewIcon(in io.ReadSeeker) (out GuildIcon, err error) {
	inner, err := NewImageData(in)

	if err != nil {
		return
	}

	out = &guild.IconImpl{Root: inner}

	if inner.Height() > IconMaxHeight || inner.Width() > IconMaxWidth {
		return out, ErrBadIconSize
	}

	return
}

// RequireNewImageData constructs a new string/json serializable image container
// from the given input file or stream.
//
// If this function encounters any error it will panic.
//
// See NewImageData for more information.
func RequireNewImageData(in io.ReadSeeker) ImageData {
	if out, err := NewImageData(in); err != nil {
		panic(err)
	} else {
		return out
	}
}

// NewImageData constructs a new string/json serializable image container from
// the given input file or stream.
//
// If an IO error occurs while attempting to work with the given stream this
// function will return `nil, error`.
//
// If the given image stream is not an image that is a gif, png, or jpeg this
// function will return `nil, ErrBadIconFormat`.
func NewImageData(in io.ReadSeeker) (ImageData, error) {
	// FIXME: This should be done in an ImageDataImpl constructor.
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	out := new(t.ImageDataImpl)

	if conf, str, err := image.DecodeConfig(in); err != nil {
		return nil, err
	} else {
		if _, err := in.Seek(0, io.SeekStart); err != nil {
			return nil, err
		}

		out.X = uint16(conf.Width)
		out.Y = uint16(conf.Height)

		switch str {
		case "png":
			out.Form = ImageFormatPNG
		case "gif":
			out.Form = ImageFormatGIF
		case "jpeg", "jpg":
			out.Form = ImageFormatJPG
		default:
			return nil, ErrBadImageFormat
		}
	}

	return out, nil
}

///  J  ////////////////////////////////////////////////////////////////////////

// NewJSONImageData constructs an ImageData instance from the given JSON byte
// slice.
//
// This method is intended to allow creating a new ImageData instance without
// having a local copy of the image data.
func NewJSONImageData(in []byte) (ImageData, error) {
	return t.NewJSONImageData(in)
}

///  R  ////////////////////////////////////////////////////////////////////////

// NewRole
func NewRole(validate bool) Role {
	return (&guild.role{
		Validate: validate,
	}).SetColor(lib.NewColor())
}

///  S  ////////////////////////////////////////////////////////////////////////

// NewSnowflake constructs a new empty Snowflake instance.
//
// The `validate` flag enables on-the-fly validation of this type.  When enabled
// the setters on the Snowflake type will panic if provided an invalid input
// value.  Additionally the Marshal/Unmarshal methods will return validation
// errors if the value being (un)marshalled is invalid.
//
// Note: The Snowflake instance this method creates is empty, and as such is
// invalid.  The snowflake will need to have values set via setters or
// unmarshaler methods to become valid.
func NewSnowflake() Snowflake {
	return t.NewSnowflake()
}

///  U  ////////////////////////////////////////////////////////////////////////

// Returns a new, empty User instance.
//
// Data Validation
//
// The "validate" flag enables/disables internal input validation.  This will
// check that values being set via the type's setter methods are valid according
// to the documented Discord restrictions.
//
// Note: internal input validation does not guarantee that Discord will accept
// the value, there are undocumented internal checks that the Discord API does
// in addition to the documented checks.  This feature is primarily intended to
// help catch internal bugs and/or avoid making requests with obviously bad
// data.
func NewUser() User {
	return user.NewUser()
}

///  V  ////////////////////////////////////////////////////////////////////////

// NewVoiceRegion
// TODO: Document me
func NewVoiceRegion() VoiceRegion {
	return new(voice.VoiceRegionImpl)
}

func NewVoiceState(validate bool) VoiceState {
	return &voice.VoiceStateImpl{Validate: validate}
}

func NewWidget() Widget {
	return guild.NewWidgetImpl()
}
