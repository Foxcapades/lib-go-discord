package discord

import (
	"encoding/json"
	"github.com/francoispqt/gojay"
)

type ActivityAssets interface {
	json.Marshaler
	json.Unmarshaler

	gojay.MarshalerJSONObject
	gojay.UnmarshalerJSONObject

	// LargeImage returns the current value of this record's `large_image` field.
	//
	// The `large_image` field contains the id for a large asset of the activity,
	// usually a snowflake
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use LargeImageIsSet to check if the field is present before use.
	LargeImage() string

	// LargeImageIsSet returns whether this record's `large_image` field is
	// currently present.
	LargeImageIsSet() bool

	// SetLargeImage overwrites the current value of this record's `large_image`
	// field.
	SetLargeImage(string) ActivityAssets

	// UnsetLargeImage removes this record's `large_image` field.
	UnsetLargeImage() ActivityAssets

	// LargeText returns the current value of this record's `large_text` field.
	//
	// The `large_text` field contains the text displayed when hovering over the
	// large image of the activity.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use LargeTextIsSet to check if the field is present before use.
	LargeText() string

	// LargeTextIsSet returns whether this record's `large_text` field is
	// currently present.
	LargeTextIsSet() bool

	// SetLargeText overwrites the current value of this record's `large_text`
	// field.
	SetLargeText(string) ActivityAssets

	// UnsetLargeText removes this record's `large_text` field.
	UnsetLargeText() ActivityAssets

	// SmallImage returns the current value of this record's `small_image` field.
	//
	// The `small_image` field contains the id for a small asset of the activity,
	// usually a snowflake
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use SmallImageIsSet to check if the field is present before use.
	SmallImage() string

	// SmallImageIsSet returns whether this record's `small_image` field is
	// currently present.
	SmallImageIsSet() bool

	// SetSmallImage overwrites the current value of this record's `small_image`
	// field.
	SetSmallImage(string) ActivityAssets

	// UnsetSmallImage removes this record's `small_image` field.
	UnsetSmallImage() ActivityAssets

	// SmallText returns the current value of this record's `small_text` field.
	//
	// The `small_text` field contains the text displayed when hovering over the
	// small image of the activity.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use SmallTextIsSet to check if the field is present before use.
	SmallText() string

	// SmallTextIsSet returns whether this record's `small_text` field is
	// currently present.
	SmallTextIsSet() bool

	// SetSmallText overwrites the current value of this record's `small_text`
	// field.
	SetSmallText(string) ActivityAssets

	// UnsetSmallText removes this record's `small_text` field.
	UnsetSmallText() ActivityAssets
}
