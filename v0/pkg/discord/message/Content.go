package message

import (
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)

var (
	ErrBadContentLength = errors.New("content length must not exceed 2000 characters")
)

type Content string

func (c Content) IsValid() bool {
	return nil == c.Validate()
}

func (c Content) Validate() error {
	if len(c) > 2000 {
		return ErrBadContentLength
	}

	return nil
}

type TriStateContentField struct {
	value *Content
	null  bool
}

func (t TriStateContentField) IsSet() bool {
	return t.value != nil
}

func (t TriStateContentField) IsUnset() bool {
	return t.value == nil && !t.null
}

func (t TriStateContentField) IsNull() bool {
	return t.null
}

func (t TriStateContentField) IsReadable() bool {
	return t.value != nil
}

func (t TriStateContentField) IsNotNull() bool {
	return !t.null
}

func (t TriStateContentField) SetNull() {
	t.value = nil
	t.null = true
}

func (t TriStateContentField) Unset() {
	t.value = nil
	t.null = false
}

func (t TriStateContentField) Get() Content {
	if t.value == nil {
		if t.null {
			panic(dlib.ErrNullField)
		} else {
			panic(dlib.ErrUnsetField)
		}
	}

	return *t.value
}

func (t TriStateContentField) Set(c Content) {
	t.value = &c
	t.null = false
}
