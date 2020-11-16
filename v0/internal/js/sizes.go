package js

import "github.com/foxcapades/go-bytify/v0/bytify"

const (
	SingleQuoteSize = 1
	QuoteSize       = SingleQuoteSize * 2
	BracketSize     = 2
	FirstFieldSize  = QuoteSize + 1
	NextFieldSize   = QuoteSize + 2

	// FirstKeySize is the number of JSON overhead characters needed for the first
	// json object key.
	//
	//     FirstKeySize = len(`"":`)
	FirstKeySize = QuoteSize + 1

	// NextKeySize is the number of JSON overhead characters needed for object
	// keys after the first key.
	//
	//     NextKeySize = len(`,"":`)
	NextKeySize = QuoteSize + 2

	NullSize = 4
)

func SizeUint16OrNull(v *uint16) uint32 {
	if v == nil {
		return NullSize
	}

	return uint32(bytify.Uint16StringSize(*v))
}
