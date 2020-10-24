package discord

import (
	"errors"
)

var (
	ErrBadContentLength = errors.New("content length must not exceed 2000 characters")
)

type MessageContent string

func (c MessageContent) IsValid() bool {
	return nil == c.Validate()
}

func (c MessageContent) Validate() error {
	if len(c) > 2000 {
		return ErrBadContentLength
	}

	return nil
}

type TriStateContentField struct {
	value *MessageContent
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

func (t TriStateContentField) Get() MessageContent {
	if t.value == nil {
		if t.null {
			panic(ErrNullField)
		} else {
			panic(ErrUnsetField)
		}
	}

	return *t.value
}

func (t TriStateContentField) Set(c MessageContent) {
	t.value = &c
	t.null = false
}
