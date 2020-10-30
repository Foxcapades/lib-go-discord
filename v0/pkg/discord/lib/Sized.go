package lib

type Sized interface {
	// BufferSize returns the number of bytes needed to store the serialized form
	// of this object.
	BufferSize() uint32
}
