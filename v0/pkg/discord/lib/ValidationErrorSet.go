package lib

import (
	"encoding/json"
	"fmt"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"github.com/francoispqt/gojay"
)

const (
	GeneralErrorsKey = "general_errors"
)

type ValidationErrorSet interface {
	error

	json.Marshaler

	gojay.MarshalerJSONObject

	// Returns the set top level error message.
	//
	// Same as calling `.Error()`.
	Message() string

	// SetMessage sets the top level message for this error.
	//
	// This message will be returned when the `.Error()` method is called.
	SetMessage(string) ValidationErrorSet

	// AppendErrorSet appends all the errors from the given error set to this
	// error set.
	//
	// If the given error value is nil, this method does nothing.
	AppendErrorSet(err ValidationErrorSet) ValidationErrorSet

	// AppendKeyedError adds the given keyed error string to the list of field
	// specific errors.
	AppendKeyedError(key serial.Key, err string) ValidationErrorSet

	// AppendRawKeyedError adds the given keyed error to the list of field
	// specific errors.
	//
	// If the given error value is nil, this method does nothing.
	//
	// If the given error is an instance of ValidationErrorSet, this method
	// will unpack the given error set into this error set.
	AppendRawKeyedError(key serial.Key, err error) ValidationErrorSet

	// AppendGeneralError adds the given error text to the list of general, or
	// non-field specific errors.
	AppendGeneralError(err string) ValidationErrorSet

	// AppendRawError adds the given error's text value to the list of general
	// errors contained by this error set.
	//
	// If the given error value is nil, this method does nothing.
	AppendRawError(err error) ValidationErrorSet

	// GetKeyedErrors returns a map of keys (such as json field name) to a list
	// of errors relating to each key.
	//
	// General/unkeyed errors are not included in this mapping.
	GetKeyedErrors() map[string][]string

	// GetGeneralErrors returns only the unkeyed errors.
	GetGeneralErrors() []string

	// GetAllErrors returns all contained errors in a map of key to value.
	//
	// General/non-keyed errors will be available using the GeneralErrorsKey const
	// as the map key.
	GetAllErrors() map[string][]string

	// GetSize returns the number of errors contained in this error set.
	GetSize() uint16
}

func NewValidationErrorSet() ValidationErrorSet {
	return new(validationErrorSet)
}

type validationErrorSet struct {
	message string
	general []string
	keyed   map[string][]string
	size    uint16
}

func (v *validationErrorSet) Error() string {
	return v.Message()
}

func (v *validationErrorSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.GetAllErrors())
}

func (v *validationErrorSet) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddSliceStringKey(GeneralErrorsKey, v.general)
	for k, v := range v.keyed {
		enc.AddSliceStringKey(k, v)
	}
}

func (v *validationErrorSet) IsNil() bool {
	return v == nil
}

func (v *validationErrorSet) Message() string {
	msg := v.message

	if len(msg) == 0 {
		msg = "no message provided"
	}

	if v.size == 0 {
		return msg
	} else if v.size == 1 {
		return fmt.Sprintf("%s (1 error)", msg)
	} else {
		return fmt.Sprintf("%s (%d errors)", msg, v.size)
	}
}

func (v *validationErrorSet) SetMessage(s string) ValidationErrorSet {
	v.message = s
	return v
}

func (v *validationErrorSet) AppendErrorSet(err ValidationErrorSet) ValidationErrorSet {
	if err == nil {
		return v
	}

	for _, e := range err.GetGeneralErrors() {
		v.general = append(v.general, e)
		v.size++
	}

	for k, e := range err.GetKeyedErrors() {
		v.keyed[k] = e
		v.size += uint16(len(e))
	}

	return v
}

func (v *validationErrorSet) AppendKeyedError(key serial.Key, err string) ValidationErrorSet {
	sk := string(key)

	if _, ok := v.keyed[sk]; ok {
		v.keyed[sk] = append(v.keyed[sk], err)
	}

	v.size++

	return v
}

func (v *validationErrorSet) AppendRawKeyedError(key serial.Key, err error) ValidationErrorSet {
	if err == nil {
		return v
	}

	if tmp, ok := err.(ValidationErrorSet); ok {
		for k, val := range tmp.GetKeyedErrors() {
			v.keyed[string(key)+"."+k] = val
		}

		v.general = append(v.general, tmp.GetGeneralErrors()...)

		return v
	}

	return v.AppendKeyedError(key, err.Error())
}

func (v *validationErrorSet) AppendGeneralError(err string) ValidationErrorSet {
	v.general = append(v.general, err)
	v.size++

	return v
}

func (v *validationErrorSet) AppendRawError(err error) ValidationErrorSet {
	if err == nil {
		return v
	}

	if tmp, ok := err.(ValidationErrorSet); ok {
		return v.AppendErrorSet(tmp)
	}

	v.general = append(v.general, err.Error())
	v.size++

	return v
}

func (v *validationErrorSet) GetKeyedErrors() map[string][]string {
	return v.keyed
}

func (v *validationErrorSet) GetGeneralErrors() []string {
	return v.general
}

func (v *validationErrorSet) GetAllErrors() map[string][]string {
	tmp := make(map[string][]string, len(v.keyed)+1)

	for k, v := range v.keyed {
		tmp[k] = v
	}

	tmp[GeneralErrorsKey] = v.general

	return tmp
}

func (v *validationErrorSet) GetSize() uint16 {
	return v.size
}
