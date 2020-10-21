package image

import "errors"

var (
	ErrBadImageFormat = errors.New("unrecognized or invalid image format")
)

type Format string

const (
	FormatGIF Format = "image/gif"
	FormatJPG Format = "image/jpeg"
	FormatPNG Format = "image/png"
)

// Is returns whether the given string is a valid representation of the current
// format.
//
// For example, an input string of "jpg", "jpeg", or "image/jpeg" would return
// true for the value FormatJPG
func (f Format) Is(tp string) bool {
	switch f {
	case FormatPNG:
		return tp == "png" || tp == string(FormatPNG)
	case FormatGIF:
		return tp == "gif" || tp == string(FormatGIF)
	case FormatJPG:
		return tp == "jpg" || tp == "jpeg" || tp == string(FormatJPG)
	}

	return false
}

func (f Format) IsValid() bool {
	return f.Validate() == nil
}

func (f Format) Validate() error {
	switch f {
	case FormatPNG, FormatGIF, FormatJPG:
		return nil
	default:
		return ErrBadImageFormat
	}
}


