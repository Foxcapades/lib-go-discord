package meta

type Sized interface {
	// JSONSize returns the projected number of bytes required to store this value
	// as a JSON string.
	JSONSize() uint32
}
