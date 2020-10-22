package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

var (
	ErrDiscriminatorLength = errors.New("discriminator values must be " +
		"exactly 4 digits in length")
)

type Discriminator uint16

func (d *Discriminator) UnmarshalJSON(bytes []byte) error {
	var tmp string

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if val, err := strconv.ParseUint(tmp, 10, 16); err != nil {
		return err
	} else {
		*d = Discriminator(val)
	}

	return nil
}

func (d Discriminator) MarshalJSON() ([]byte, error) {
	out := new(bytes.Buffer)
	out.WriteByte('"')
	out.WriteString(strconv.FormatUint(uint64(d), 10))
	out.WriteByte('"')

	return out.Bytes(), nil
}

func (d Discriminator) IsValid() bool {
	return nil == d.Validate()
}

func (d Discriminator) Validate() error {
	if d < 1000 || d > 9999 {
		return ErrDiscriminatorLength
	}

	return nil
}

