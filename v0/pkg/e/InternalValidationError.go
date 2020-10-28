package e

type InternalValidationError interface {
	error

	Cause() error
	Value() interface{}
	Field() string
}

func NewInternalValidationError(
	field string,
	value interface{},
	cause error,
) InternalValidationError {
	return &internalValidationError{
		field: field,
		value: value,
		cause: cause,
	}
}

type internalValidationError struct {
	field string
	value interface{}
	cause error
}

func (i *internalValidationError) Cause() error {
	return i.cause
}

func (i *internalValidationError) Error() string {
	return i.cause.Error()
}

func (i *internalValidationError) Value() interface{} {
	return i.value
}

func (i *internalValidationError) Field() string {
	return i.field
}
