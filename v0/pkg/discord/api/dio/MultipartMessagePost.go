package dio

import "io"

type MultipartMessagePost interface {
	JSONMessagePost

	// File returns the current value of this record's `file` field.
	//
	// The `file` field contains the contents of the file being sent.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use FileIsSet to check if the field is present before use.
	File() io.ReadCloser

	// FileIsSet returns whether this record's `file` field is currently present.
	FileIsSet() bool

	// SetFile overwrites the current value of this record's `file` field.
	SetFile(io.ReadCloser) MultipartMessagePost

	// UnsetFile removes this record's `file` field.
	UnsetFile() MultipartMessagePost
}

func NewMultipartMessagePost(validate bool) MultipartMessagePost {
	panic("implement me")
}
