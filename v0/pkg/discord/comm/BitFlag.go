package comm

type BitFlag interface {
	// Add returns the value of the current flag bitwise ORd with the given flag.
	//
	// This method will not modify the value of the current flag.
	Add(flag BitFlag) BitFlag

	// Remove returns the value of the current flag bitwise ANDed with the inverse
	// of the given flag.
	//
	// This method will not modify the value of the current flag.
	Remove(flag BitFlag) BitFlag

	// Mask returns the value of the current flag bitwise ANDed with the given
	// flag.
	Mask(flag BitFlag) BitFlag
}
