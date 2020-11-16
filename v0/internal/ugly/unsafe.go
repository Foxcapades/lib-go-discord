package ugly

import (
	"unsafe"
)

// UnsafeString returns a mutable "string" by raw casting the given byte slice.
//
// This is primarily intended to do string "casts" without copying the
// underlying data.
//
// This also means that the caller of this function takes full ownership of the
// backing byte slice and never passes that slice around while this "string"
// exists.
func UnsafeString(in []byte) string {
	return *(*string)(unsafe.Pointer(&in))
}
