package user

type Nickname string

func (n Nickname) IsValid() bool {
	ln := len(scrubName(string(n)))

	return ln > 0 && ln < 33
}
