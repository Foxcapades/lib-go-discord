package discord

import (
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/internal/js"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils"
	"io"
)

var (
	ErrBadContentLength = errors.New("content length must not exceed 2000 characters")
)

type MessageContent string

func (c MessageContent) MarshalJSON() ([]byte, error) {
	return c.ToJSONBytes(), nil
}

func (c *MessageContent) UnmarshalJSON(bytes []byte) error {
	return utils.UnmarshalStringInto(bytes, (*string)(c))
}

func (c *MessageContent) IsNil() bool {
	return c == nil
}

func (c MessageContent) ToJSONBytes() (buf []byte) {
	sz := len(c) + js.QuoteSize
	buf = make([]byte, sz)
	buf[0] = '"'
	copy(buf[1:], c)
	buf[sz-1] = '"'

	return
}

func (c *MessageContent) AppendJSONBytes(writer io.Writer) (err error) {
	_, err = writer.Write(c.ToJSONBytes())
	return
}

func (c MessageContent) JSONSize() uint32 {
	return uint32(len(c) + 2)
}

func (c MessageContent) IsValid() bool {
	return nil == c.Validate()
}

func (c MessageContent) Validate() error {
	if len(c) > 2000 {
		return ErrBadContentLength
	}

	return nil
}
