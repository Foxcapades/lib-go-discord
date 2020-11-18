package discord

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/foxcapades/go-bytify/v0/bytify"
	"io"
	"strconv"
)

var (
	ErrDiscriminatorLength = errors.New("discriminator values must be " +
		"exactly 4 digits in length")
)

type Discriminator uint16

func (d Discriminator) JSONSize() uint32 {
	switch true {
	case d < 10:
		return 1
	case d < 100:
		return 2
	case d < 1_000:
		return 3
	case d < 10_000:
		return 4
	default:
		return 5
	}
}

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

func (d Discriminator) String() string {
	return strconv.FormatUint(uint64(d), 10)
}

func (d *Discriminator) IsNil() bool {
	return d == nil
}

func (d *Discriminator) ToJSONBytes() []byte {
	return bytify.Uint16ToByteSlice(uint16(*d))
}

func (d *Discriminator) AppendJSONBytes(writer io.Writer) error {
	_, err := writer.Write(d.ToJSONBytes())
	return err
}
