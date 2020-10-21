package image

import (
	"encoding/json"
	"errors"
)

var (
	ErrBadImageDataJsonType = errors.New("image data passed to" +
		" Icon.UnmarshalJSON was not a string value")
	ErrBadImageDataSize = errors.New("image data passed to Icon.UnmarshalJSON" +
		" is too short to be a valid image")
	ErrBadImageDataFormat = errors.New("unrecognized or invalid image data" +
		" string format")
	ErrTypeMismatch = errors.New("the type declared in the image data prefix" +
		" does not match the actual image type")
)

const (
	// EncodingPrefix is the leader string prepended on image data strings.
	EncodingPrefix = "data:"

	// EncodingInfix is the dividing string between the image format and the image
	// data, informing the consumer of the encoding format.
	EncodingInfix  = ";base64,"

	// MinPossibleSize is the minimum possible size of a valid image in bytes.
	//
	// This is based on a 1x1 pixel image with no metadata and the highest
	// possible compression.
	//
	// Sizes reached per valid format:
	//    gif  = 799 bytes
	//    jpeg = 280 bytes
	//    png  = 67  bytes
	MinPossibleSize uint8 = 67

	// MinPossibleB64Size is the minimum possible base64 encoded size of a valid
	// image in bytes.
	//
	// Sizes reached per valid format using the above test images as input:
	//    gif  = 1065 bytes
	//    jpeg = 372  bytes
	//    png  = 88   bytes
	MinPossibleB64Size uint8 = 88

	MinPossibleImageString = uint16(len(EncodingPrefix) +
		len(EncodingInfix) +
		int(MinPossibleB64Size))
)

type Data interface {
	json.Marshaler
	json.Unmarshaler

	String() (string, error)
	Format() Format
	Width() uint16
	Height() uint16
	AspectRatio() AspectRatio

	// Close attempts to close the underlying data source (if it is closable).
	//
	// If the underlying data source is not closable then this method does
	// nothing.
	Close() error
}
