package utils

func GCD_U16(a, b uint16) uint16 {
	for b != 0 {
		t := b
		b = a % b
		a= t
	}

	return a
}
