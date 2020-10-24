package types

import (
	"bytes"
	"encoding/base64"
	"image"
	"io"
	"strings"

	"github.com/foxcapades/lib-go-discord/v0/internal/utils"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

// NewJSONImageData constructs an ImageData instance from the given JSON byte
// slice.
//
// This method is intended to allow creating a new ImageData instance without
// having a local copy of the image data.
func NewJSONImageData(in []byte) (ImageData, error) {
	out := new(ImageDataImpl)

	if err := out.UnmarshalJSON(in); err != nil {
		return nil, err
	}

	return out, nil
}

type ImageDataImpl struct {
	Input io.ReadSeeker
	Form  ImageFormat
	X, Y  uint16
}

func (i *ImageDataImpl) MarshalJSON() ([]byte, error) {
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
// the current ImageData instance.
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
// recognized by this library.  See image.ImageFormat for the list of allowed types.
//
// image.ErrTypeMismatch = The image format declared in the image data string's
// prefix does not match the actual type of the base64 encoded image.
//
// In addition to the above errors, this method may return errors from the
// standard library's `image` or `base64` packages, or any error returned from
// attempting to close the previous image data held by this ImageData instance.
func (i *ImageDataImpl) UnmarshalJSON(in []byte) error {
	if err := i.Close(); err != nil {
		return err
	}

	ln := len(in)

	// +2 to account for double quotes
	if ln < int(MinPossibleImageString)+2 {
		return ErrBadImageDataSize
	}

	if in[0] != '"' {
		return ErrBadImageDataJsonType
	}

	dataStart := bytes.IndexByte(in, ',')
	if dataStart == -1 {
		return ErrBadImageDataFormat
	}

	i.Form = ImageFormat(in[len(EncodingPrefix):dataStart])

	if err := i.Form.Validate(); err != nil {
		return err
	}

	// +2 because we don't need the leading comma or the trailing double quote
	buf := make([]byte, 0, ln-dataStart+2)
	if _, err := base64.StdEncoding.Decode(in[dataStart+1:], buf); err != nil {
		return err
	}

	i.Input = bytes.NewReader(buf)

	conf, tp, err := image.DecodeConfig(i.Input)
	_, _ = i.Input.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	if !i.Form.Is(tp) {
		return ErrTypeMismatch
	}

	i.X = uint16(conf.Width)
	i.Y = uint16(conf.Height)

	return nil
}

func (i *ImageDataImpl) String() (string, error) {
	out := new(strings.Builder)
	out.Grow(4096)
	out.Reset()
	out.WriteString(EncodingPrefix)
	out.WriteString(string(i.Form))
	out.WriteString(EncodingInfix)

	b := base64.NewEncoder(base64.StdEncoding, out)
	buf := make([]byte, 0, 1024)

	for {
		if n, err := i.Input.Read(buf); err != nil && err != io.EOF {
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

func (i *ImageDataImpl) Format() ImageFormat {
	return i.Form
}

func (i *ImageDataImpl) Width() uint16 {
	return i.X
}

func (i *ImageDataImpl) Height() uint16 {
	return i.Y
}

func (i *ImageDataImpl) AspectRatio() lib.AspectRatio {
	gdc := utils.GCD_U16(i.X, i.Y)
	return lib.AspectRatio{
		Width:  i.X / gdc,
		Height: i.Y / gdc,
	}
}

// Close will attempt to close the input io.ReadSeeker instance if it is also
// an instance of io.Closer.
//
// If the input io.ReadSeeker is not also an instance of io.Closer, then this
// method will do nothing.
func (i *ImageDataImpl) Close() error {
	if tmp, ok := i.Input.(io.Closer); ok {
		return tmp.Close()
	}

	return nil
}

