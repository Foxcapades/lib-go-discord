package embed

import "errors"

var (
	ErrBadEmbedType = errors.New("unrecognized embed type value")
)

// Type
// TODO: Document me
//
// Embed types are "loosely defined" and, for the most part, are not used by our
// clients for rendering. Embed attributes power what is rendered. Embed types
// should be considered deprecated and might be removed in a future API version.
type Type string

const (
	// Generic embed rendered from embed attributes.
	TypeRich Type = "rich"

	// Image embed.
	TypeImage Type = "image"

	// Video embed.
	TypeVideo Type = "video"

	// Animated gif image embed rendered as a video embed.
	TypeGifV Type = "gifv"

	// Article embed.
	TypeArticle Type = "article"

	// Link embed.
	TypeLink Type = "link"
)

func (e Type) IsValid() bool {
	return nil == e.Validate()
}

var embedTypes = []Type{
	TypeRich,
	TypeImage,
	TypeVideo,
	TypeGifV,
	TypeArticle,
	TypeLink,
}

func (e Type) Validate() error {
	for i := range embedTypes {
		if e == embedTypes[i] {
			return nil
		}
	}

	return ErrBadEmbedType
}
