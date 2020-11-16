package ugly

import (
	"errors"
	"io"

	"github.com/foxcapades/lib-go-discord/v0/internal/js"
)

// PseudoString is an unsafe string accessor over a mutable byte array.
//
// This type should _never_ be exposed outside of the internal package of this
// library.
type PseudoString []byte

func (p *PseudoString) AppendJSONBytes(buf io.Writer) (e error) {
	if p.IsNil() {
		_, e = buf.Write(js.NullBytesBuf)
		return
	}

	_, e = buf.Write(js.SingleQuoteBuf)

	if e == nil {
		_, e = buf.Write(*p)
	}

	if e == nil {
		_, e = buf.Write(js.SingleQuoteBuf)
	}

	return
}

func (p PseudoString) MarshalJSON() ([]byte, error) {
	sz := p.JSONSize()
	out := make([]byte, sz)
	out[0] = '"'
	out[sz-1] = '"'
	copy(out[1:], p)

	return out, nil
}

func (p *PseudoString) UnmarshalJSON(in []byte) error {
	ln := len(in)

	if ln < 2 || in[0] != '"' || in[ln-1] != '"' {
		return errors.New("") //FIXME: invalid json type
	}

	*p = make([]byte, ln-2)
	copy(*p, in[1:])

	return nil
}

func (p PseudoString) ToJSONBytes() []byte {
	return p
}

func (p PseudoString) String() string {
	return UnsafeString(p[1 : len(p)-1])
}

func (p *PseudoString) IsNil() bool {
	return p == nil
}

func (p PseudoString) JSONSize() uint32 {
	if p == nil {
		return 4
	}

	return uint32(len(p) + 2)
}
