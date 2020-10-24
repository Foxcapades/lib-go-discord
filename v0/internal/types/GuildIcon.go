package types

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type IconImpl struct {
	Root ImageData
}

func (i *IconImpl) IsValid() bool {
	return nil == i.Validate()
}

func (i *IconImpl) Validate() error {
	if err := i.Format().Validate(); err != nil {
		return err
	}

	if i.Root.Width() > IconMaxWidth || i.Root.Height() > IconMaxHeight {
		return ErrBadIconSize
	}

	return nil
}

func (i *IconImpl) MarshalJSON() ([]byte, error) {
	return i.Root.MarshalJSON()
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
// recognized by this library.  See image.ImageFormat for the list of allowed types.
//
// image.ErrTypeMismatch = The image format declared in the image data string's
// prefix does not match the actual type of the base64 encoded image.
//
// In addition to the above errors, this method may return errors from the
// standard library's `image` or `base64` packages, or any error returned from
// attempting to close the previous image data held by this GuildIcon instance.
func (i *IconImpl) UnmarshalJSON(in []byte) (err error) {
	if err = i.Root.Close(); err != nil {
		return
	}

	i.Root, err = NewJSONImageData(in)

	return
}

func (i *IconImpl) String() (string, error) {
	return i.Root.String()
}

func (i *IconImpl) Format() ImageFormat {
	return i.Root.Format()
}

func (i *IconImpl) Width() uint16 {
	return i.Root.Width()
}

func (i *IconImpl) Height() uint16 {
	return i.Root.Height()
}

func (i *IconImpl) AspectRatio() lib.AspectRatio {
	return i.Root.AspectRatio()
}

// Close will attempt to close the input io.ReadSeeker instance if it also
// happens to be an instance of io.Closer.
//
// If the input io.ReadSeeker is not also an instance of io.Closer, then this
// method will do nothing.
func (i *IconImpl) Close() error {
	return i.Root.Close()
}
