package discord

import (
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/internal/js"
	"github.com/francoispqt/gojay"
	"io"
)

var (
	ErrBadChannelNameLength = errors.New("channel names must be 2-100" +
		" characters in length")
)

type ChannelName string

func (n ChannelName) JSONSize() uint32 {
	return uint32(len(n)) + js.QuoteSize
}

func (n ChannelName) IsValid() bool {
	return nil == n.Validate()
}

func (n ChannelName) Validate() error {
	l := len(n)
	if l < 2 || l > 100 {
		return ErrBadChannelNameLength
	}

	return nil
}

func (n ChannelName) MarshalJSON() ([]byte, error) {
	sz := n.JSONSize()

	tmp := make([]byte, sz)
	tmp[0] = js.SingleQuote
	tmp[sz-1] = js.SingleQuote

	copy(tmp[1:], n)

	return tmp, nil
}

func (n *ChannelName) UnmarshalJSON(bytes []byte) error {
	var tmp string

	if err := gojay.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	*n = ChannelName(tmp)

	return nil
}

func (n *ChannelName) IsNil() bool {
	return n == nil
}

func (n ChannelName) ToJSONBytes() []byte {
	panic("implement me")
}

func (n *ChannelName) AppendJSONBytes(writer io.Writer) error {
	tmp, _ := n.MarshalJSON()
	_, err := writer.Write(tmp)

	return err
}

func DecodeChannelName(dec *gojay.Decoder) (*ChannelName, error) {
	var tmp *string

	if err := dec.StringNull(&tmp); err != nil {
		return nil, err
	}

	return (*ChannelName)(tmp), nil
}
