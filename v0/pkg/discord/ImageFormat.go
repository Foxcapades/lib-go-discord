package discord

type ImageFormat string

const (
	ImageFormatGIF ImageFormat = "image/gif"
	ImageFormatJPG ImageFormat = "image/jpeg"
	ImageFormatPNG ImageFormat = "image/png"
)

// Is returns whether the given string is a valid representation of the current
// format.
//
// For example, an input string of "jpg", "jpeg", or "image/jpeg" would return
// true for the value ImageFormatJPG
func (f ImageFormat) Is(tp string) bool {
	switch f {
	case ImageFormatPNG:
		return tp == "png" || tp == string(ImageFormatPNG)
	case ImageFormatGIF:
		return tp == "gif" || tp == string(ImageFormatGIF)
	case ImageFormatJPG:
		return tp == "jpg" || tp == "jpeg" || tp == string(ImageFormatJPG)
	}

	return false
}

func (f ImageFormat) IsValid() bool {
	return f.Validate() == nil
}

func (f ImageFormat) Validate() error {
	switch f {
	case ImageFormatPNG, ImageFormatGIF, ImageFormatJPG:
		return nil
	default:
		return ErrBadImageFormat
	}
}


