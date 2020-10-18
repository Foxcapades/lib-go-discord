package message

import "errors"

var (
	ErrBadEmbedType = errors.New("unrecognized embed type value")
)

// EmbedType
// TODO: Document me
//
// Embed types are "loosely defined" and, for the most part, are not used by our
// clients for rendering. Embed attributes power what is rendered. Embed types
// should be considered deprecated and might be removed in a future API version.
type EmbedType string

const (
	// Generic embed rendered from embed attributes.
	EmbedTypeRich EmbedType = "rich"

	// Image embed.
	EmbedTypeImage EmbedType = "image"

	// Video embed.
	EmbedTypeVideo EmbedType = "video"

	// Animated gif image embed rendered as a video embed.
	EmbedTypeGifV EmbedType = "gifv"

	// Article embed.
	EmbedTypeArticle EmbedType = "article"

	// Link embed.
	EmbedTypeLink EmbedType = "link"
)

func (e EmbedType) IsValid() bool {
	return nil == e.Validate()
}

var embedTypes = []EmbedType{
	EmbedTypeRich,
	EmbedTypeImage,
	EmbedTypeVideo,
	EmbedTypeGifV,
	EmbedTypeArticle,
	EmbedTypeLink,
}

func (e EmbedType) Validate() error {
	for i := range embedTypes {
		if e == embedTypes[i] {
			return nil
		}
	}

	return ErrBadEmbedType
}
