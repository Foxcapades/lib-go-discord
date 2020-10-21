package guild

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"strings"

	img "image"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/foxcapades/lib-go-discord/v0/internal/utils"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/comm"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/comm/image"
)

var (
	ErrBadIconSize = errors.New("given image is too large to use as a guild" +
		" icon; images must be no larger than 1024x1024")
	ErrBadIconFormat = errors.New("unrecognized image format passed to " +
		"guild.NewIcon; must be one of gif, png, or jpeg")
)

const (
	IconMaxWidth  uint16 = 1024
	IconMaxHeight uint16 = 1024
)

type Icon interface {
	comm.Validatable
	image.Data
}

// RequireNewIcon constructs a new Guild Icon from the given input file or
// stream.
//
// If this function encounters any error, including ErrBadIconSize, it will
// panic.
//
// See NewIcon() for more information.
func RequireNewIcon(in io.ReadSeeker) Icon {
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
func NewIcon(in io.ReadSeeker) (Icon, error) {
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	out := new(iconImpl)

	if conf, str, err := img.DecodeConfig(in); err != nil {
		return nil, err
	} else {
		if _, err := in.Seek(0, io.SeekStart); err != nil {
			return nil, err
		}

		out.width = uint16(conf.Width)
		out.height = uint16(conf.Height)

		switch str {
		case "png":
			out.form = image.FormatPNG
		case "gif":
			out.form = image.FormatGIF
		case "jpeg", "jpg":
			out.form = image.FormatJPG
		default:
			return nil, ErrBadIconFormat
		}
	}

	if out.height > IconMaxHeight || out.width > IconMaxWidth {
		return out, ErrBadIconSize
	}

	return out, nil
}

type iconImpl struct {
	input  io.ReadSeeker
	form   image.Format
	width  uint16
	height uint16
}

func (i *iconImpl) IsValid() bool {
	return nil == i.Validate()
}

func (i *iconImpl) Validate() error {
	if err := i.form.Validate(); err != nil {
		return err
	}

	if i.width > IconMaxWidth || i.height > IconMaxHeight {
		return ErrBadIconSize
	}

	return nil
}

func (i *iconImpl) MarshalJSON() ([]byte, error) {
	out := new(bytes.Buffer)
	out.WriteByte('"')

	if tmp, err := i.String(); err != nil {
		return nil, err
	} else {
		out.WriteString(tmp)
	}

	out.WriteByte('"')

	return out.Bytes(), nil
}

// UnmarshalJSON attempts to deserialize a string encoded image data value into
// the current Icon instance.
//
// This method may return one of the following library errors:
//
// image.ErrBadImageDataSize = The base64 encoded image data is too small to be
// a valid image, it may have been truncated by the source. See
// image.MinPossibleImageString.
//
// image.ErrBadImageDataJsonType = The value passed to this method was not a
// string.
//
// image.ErrBadImageDataFormat = The given image data string is corrupt or
// malformed and contains no divider between the prefix and the actual image
// data.
//
// image.ErrBadImageFormat = The input image data string is not of a type
// recognized by this library.  See image.Format for the list of allowed types.
//
// image.ErrTypeMismatch = The image format declared in the image data string's
// prefix does not match the actual type of the base64 encoded image.
//
// In addition to the above errors, this method may return errors from the
// standard library's `image` or `base64` packages.
func (i *iconImpl) UnmarshalJSON(in []byte) error {
	ln := len(in)

	// +2 to account for double quotes
	if ln < int(image.MinPossibleImageString)+2 {
		return image.ErrBadImageDataSize
	}

	if in[0] != '"' {
		return image.ErrBadImageDataJsonType
	}

	dataStart := bytes.IndexByte(in, ',')
	if dataStart == -1 {
		return image.ErrBadImageDataFormat
	}

	i.form = image.Format(in[len(image.EncodingPrefix):dataStart])

	if err := i.form.Validate(); err != nil {
		return err
	}

	// +2 because we don't need the leading comma or the trailing double quote
	buf := make([]byte, 0, ln - dataStart + 2)
	if _, err := base64.StdEncoding.Decode(in[dataStart + 1:], buf); err != nil {
		return err
	}

	i.input = bytes.NewReader(buf)

	conf, tp, err := img.DecodeConfig(i.input)
	_, _ = i.input.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	if !i.form.Is(tp) {
		return image.ErrTypeMismatch
	}

	i.width  = uint16(conf.Width)
	i.height = uint16(conf.Height)

	return nil
}

func (i *iconImpl) String() (string, error) {
	out := new(strings.Builder)
	out.Grow(4096)
	out.Reset()
	out.WriteString(image.EncodingPrefix)
	out.WriteString(string(i.form))
	out.WriteString(image.EncodingInfix)

	b := base64.NewEncoder(base64.StdEncoding, out)
	buf := make([]byte, 0, 1024)

	for {
		if n, err := i.input.Read(buf); err != nil && err != io.EOF {
			return "", err
		} else if n > 0 {
			if _, err := b.Write(buf[:n]); err != nil {
				return "", err
			}
		} else {
			break
		}
	}

	return out.String(), nil
}

func (i *iconImpl) Format() image.Format {
	return i.form
}

func (i *iconImpl) Width() uint16 {
	return i.width
}

func (i *iconImpl) Height() uint16 {
	return i.height
}

func (i *iconImpl) AspectRatio() image.AspectRatio {
	gdc := utils.GCD_U16(i.width, i.height)
	return image.AspectRatio{
		Width:  i.width  / gdc,
		Height: i.height / gdc,
	}
}

// Close will attempt to close the input io.ReadSeeker instance if it also
// happens to be an instance of io.Closer.
//
// If the input io.ReadSeeker is not also an instance of io.Closer, then this
// method will do nothing.
func (i *iconImpl) Close() error {
	if tmp, ok := i.input.(io.Closer); ok {
		return tmp.Close()
	}

	return nil
}
