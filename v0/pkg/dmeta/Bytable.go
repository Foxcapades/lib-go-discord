package dmeta

import "io"

type Bytable interface {
	ToJSONBytes() []byte

	// AppendBytes appends the JSON value of this type to the given writer.
	//
	// This method should include any encapsulating characters such as quotes or
	// brackets.  JSON field dividers separating this value from other values,
	// such as commas or colons should be appended by the caller of this method,
	// not the implementer.  Implementers of this method should only include those
	// dividers multiple values are encapsulated in this type itself, such as in a
	// struct or slice type.
	//
	// Example:
	//   Given a field of type `string`, with the value:
	//     foo := "hello"
	//   The bytes that should be written would be:
	//     []byte(`"hello"`).
	//
	//   Given a field of type `[]string` with the value
	//     foo := []string{"yes", "no"}
	//   The bytes that would be written would be:
	//     []byte(`["yes","no"]`)
	AppendJSONBytes(writer io.Writer) error
}
