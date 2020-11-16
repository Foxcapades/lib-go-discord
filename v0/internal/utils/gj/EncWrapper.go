package gj

import (
	"fmt"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"github.com/francoispqt/gojay"
)

type EncWrapper interface {
	// AddBool appends the given boolean value to the output encoder.
	AddBool(key serial.Key, val bool) EncWrapper

	// AddOptionalBool appends the given boolean pointer to the output encoder
	// only if the pointer is not nil.
	AddOptionalBool(key serial.Key, val *bool) EncWrapper

	// AddNullableBool appends either the given boolean pointer or a null value to
	// the output encoder based on whether the given pointer is nil.
	AddNullableBool(key serial.Key, val *bool) EncWrapper

	// AddTriStateBool may append either the given boolean pointer or a null value
	// to the output encoder based on whether the given pointer is nil and the
	// value of the given null flag.
	//
	// If the given boolean pointer is not nil, the value will be appended to the
	// output encoder.
	//
	// Else, if the null flag is false and the given pointer is nil, nothing will
	// be appended to the output encoder.
	//
	// Else, if the null flag is true and the given pointer is nil, a null value
	// will be appended to the output encoder.
	AddTriStateBool(key serial.Key, val *bool, null bool) EncWrapper

	// AddSlice appends the given slice/array value to the output encoder.
	//
	// If the given slice is nil, an empty json array will be appended to the
	// output encoder.
	AddSlice(key serial.Key, val gojay.MarshalerJSONArray) EncWrapper

	// AddOptionalSlice appends the given slice/array value to the output encoder
	// only if the value is not nil.
	AddOptionalSlice(
		key serial.Key,
		val gojay.MarshalerJSONArray,
		empty bool,
	) EncWrapper

	// AddNullableSlice appends either the given slice/array value or a null value
	// to the output encoder based on whether the given value is nil.
	AddNullableSlice(
		key serial.Key,
		val gojay.MarshalerJSONArray,
		null bool,
	) EncWrapper

	// AddTriStateSlice may append either the given slice/array value or a null
	// value to the output encoder based on whether the given pointer is nil and
	// the value of the given null flag.
	//
	// If the given slice/array value is not nil, the value will be appended to
	// the output encoder.
	//
	// Else, if the empty flag is true and the given value is nil, nothing will
	// be appended to the output encoder.
	//
	// Else, if the null flag is true and the given value is nil, a null value
	// will be appended to the output encoder.
	//
	// Else, if neither the empty flag or null flag is true, and the given value
	// is nil, an empty json array will be appended to the output encoder.
	//
	// Note: If both the empty flag and the null flag are true, the empty flag
	// takes priority, meaning nothing will be appended to the output encoder.
	AddTriStateSlice(
		key serial.Key,
		val gojay.MarshalerJSONArray,
		empty bool,
		null bool,
	) EncWrapper

	// AddString appends the given string value to the output encoder.
	AddString(key serial.Key, val string) EncWrapper

	// AddOptionalString appends the given string pointer to the output encoder
	// only if the pointer is not nil.
	AddOptionalString(key serial.Key, val *string) EncWrapper

	// AddNullableString appends either the given string pointer or a null value
	// to the output encoder based on whether the given pointer is nil.
	AddNullableString(key serial.Key, val *string) EncWrapper

	// AddStringer appends the string value of the given value to the output
	// encoder.
	AddStringer(key serial.Key, val fmt.Stringer) EncWrapper

	// AddOptionalStringer appends the string value of the given value to the
	// output encoder only if the pointer is not nil.
	AddOptionalStringer(key serial.Key, val fmt.Stringer) EncWrapper

	// AddNullableStringer appends the string value of the given value or a null
	// value to the output encoder based on whether the given pointer is nil.
	AddNullableStringer(key serial.Key, val fmt.Stringer) EncWrapper

	// AddTriStateStringer may append either the string value of the given value
	// or a null value to the output encoder based on whether the given pointer is
	// nil and the value of the given null flag.
	//
	// If the given string pointer is not nil, the value will be appended to the
	// output encoder.
	//
	// Else, if the null flag is false and the given pointer is nil, nothing will
	// be appended to the output encoder.
	//
	// Else, if the null flag is true and the given pointer is nil, a null value
	// will be appended to the output encoder.
	AddTriStateString(key serial.Key, val *string, null bool) EncWrapper

	// AddUint8 appends the given uint8 value to the output encoder.
	AddUint8(key serial.Key, val uint8) EncWrapper

	// AddOptionalUint8 appends the given uint8 pointer to the output encoder
	// only if the pointer is not nil.
	AddOptionalUint8(key serial.Key, val *uint8) EncWrapper

	// AddNullableUint8 appends either the given uint8 pointer or a null value to
	// the output encoder based on whether the given pointer is nil.
	AddNullableUint8(key serial.Key, val *uint8) EncWrapper

	// AddTriStateUint8 may append either the given uint8 pointer or a null value
	// to the output encoder based on whether the given pointer is nil and the
	// value of the given null flag.
	//
	// If the given uint8 pointer is not nil, the value will be appended to the
	// output encoder.
	//
	// Else, if the null flag is false and the given pointer is nil, nothing will
	// be appended to the output encoder.
	//
	// Else, if the null flag is true and the given pointer is nil, a null value
	// will be appended to the output encoder.
	AddTriStateUint8(key serial.Key, val *uint8, null bool) EncWrapper

	// AddUint16 appends the given uint16 value to the output encoder.
	AddUint16(key serial.Key, val uint16) EncWrapper

	// AddOptionalUint16 appends the given uint16 pointer to the output encoder
	// only if the pointer is not nil.
	AddOptionalUint16(key serial.Key, val *uint16) EncWrapper

	// AddNullableUint16 appends either the given uint16 pointer or a null value
	// to the output encoder based on whether the given pointer is nil.
	AddNullableUint16(key serial.Key, val *uint16) EncWrapper

	// AddTriStateUint16 may append either the given uint16 pointer or a null
	// value to the output encoder based on whether the given pointer is nil and
	// the value of the given null flag.
	//
	// If the given uint16 pointer is not nil, the value will be appended to the
	// output encoder.
	//
	// Else, if the null flag is false and the given pointer is nil, nothing will
	// be appended to the output encoder.
	//
	// Else, if the null flag is true and the given pointer is nil, a null value
	// will be appended to the output encoder.
	AddTriStateUint16(key serial.Key, val *uint16, null bool) EncWrapper

	// AddUint32 appends the given uint32 value to the output encoder.
	AddUint32(key serial.Key, val uint32) EncWrapper

	// AddOptionalUint32 appends the given uint32 pointer to the output encoder
	// only if the pointer is not nil.
	AddOptionalUint32(key serial.Key, val *uint32) EncWrapper

	// AddNullableUint32 appends either the given uint32 pointer or a null value
	// to the output encoder based on whether the given pointer is nil.
	AddNullableUint32(key serial.Key, val *uint32) EncWrapper

	// AddTriStateUint32 may append either the given uint32 pointer or a null
	// value to the output encoder based on whether the given pointer is nil and
	// the value of the given null flag.
	//
	// If the given uint32 pointer is not nil, the value will be appended to the
	// output encoder.
	//
	// Else, if the null flag is false and the given pointer is nil, nothing will
	// be appended to the output encoder.
	//
	// Else, if the null flag is true and the given pointer is nil, a null value
	// will be appended to the output encoder.
	AddTriStateUint32(key serial.Key, val *uint32, null bool) EncWrapper

	// AddUint64 appends the given uint64 value to the output encoder.
	AddUint64(key serial.Key, val uint64) EncWrapper

	// AddOptionalUint64 appends the given uint64 pointer to the output encoder
	// only if the pointer is not nil.
	AddOptionalUint64(key serial.Key, val *uint64) EncWrapper

	// AddNullableUint64 appends either the given uint64 pointer or a null value
	// to the output encoder based on whether the given pointer is nil.
	AddNullableUint64(key serial.Key, val *uint64) EncWrapper

	// AddTriStateUint64 may append either the given uint64 pointer or a null
	// value to the output encoder based on whether the given pointer is nil and
	// the value of the given null flag.
	//
	// If the given uint64 pointer is not nil, the value will be appended to the
	// output encoder.
	//
	// Else, if the null flag is false and the given pointer is nil, nothing will
	// be appended to the output encoder.
	//
	// Else, if the null flag is true and the given pointer is nil, a null value
	// will be appended to the output encoder.
	AddTriStateUint64(key serial.Key, val *uint64, null bool) EncWrapper

	// AddInt64 appends the given uint64 value to the output encoder.
	AddInt64(key serial.Key, val int64) EncWrapper

	// AddOptionalInt64 appends the given int64 pointer to the output encoder
	// only if the pointer is not nil.
	AddOptionalInt64(key serial.Key, val *int64) EncWrapper

	// AddNullableInt64 appends either the given int64 pointer or a null value
	// to the output encoder based on whether the given pointer is nil.
	AddNullableInt64(key serial.Key, val *int64) EncWrapper

	// AddTriStateInt64 may append either the given int64 pointer or a null
	// value to the output encoder based on whether the given pointer is nil and
	// the value of the given null flag.
	//
	// If the given int64 pointer is not nil, the value will be appended to the
	// output encoder.
	//
	// Else, if the null flag is false and the given pointer is nil, nothing will
	// be appended to the output encoder.
	//
	// Else, if the null flag is true and the given pointer is nil, a null value
	// will be appended to the output encoder.
	AddTriStateInt64(key serial.Key, val *int64, null bool) EncWrapper

	// AddSnowflake appends the given Snowflake value to the output encoder.
	//
	// If the given Snowflake value is nil, the literal string "0" will be
	// appended to the output encoder.
	AddSnowflake(key serial.Key, val discord.Snowflake) EncWrapper

	// AddOptionalSnowflake appends the given Snowflake to the output encoder
	// only if the pointer is not nil.
	AddOptionalSnowflake(key serial.Key, val discord.Snowflake) EncWrapper

	// AddNullableSnowflake appends either the given Snowflake or a null value to
	// the output encoder based on whether the given pointer is nil.
	AddNullableSnowflake(key serial.Key, val discord.Snowflake) EncWrapper

	// AddTriStateSnowflake may append either the given Snowflake or a null value
	// to the output encoder based on whether the given pointer is nil and the
	// value of the given null flag.
	//
	// If the given Snowflake is not nil, the value will be appended to the output
	// encoder.
	//
	// Else, if the null flag is false and the given pointer is nil, nothing will
	// be appended to the output encoder.
	//
	// Else, if the null flag is true and the given pointer is nil, a null value
	// will be appended to the output encoder.
	AddTriStateSnowflake(
		key serial.Key,
		val discord.Snowflake,
		null bool,
	) EncWrapper

	// AddObject appends the given object value to the output encoder.
	//
	// If the given object value is nil, an empty json object will be appended to
	// the output encoder.
	AddObject(key serial.Key, val gojay.MarshalerJSONObject) EncWrapper

	// AddOptionalObject appends the given object to the output encoder only if
	// the pointer is not nil.
	AddOptionalObject(key serial.Key, val gojay.MarshalerJSONObject) EncWrapper

	// AddNullableObject appends either the given object or a null value to the
	// output encoder based on whether the given pointer is nil.
	AddNullableObject(key serial.Key, val gojay.MarshalerJSONObject) EncWrapper

	// AddTriStateObject may append either the given object or a null value to the
	// output encoder based on whether the given pointer is nil and the value of
	// the given null flag.
	//
	// If the given object is not nil, the value will be appended to the output
	// encoder.
	//
	// Else, if the empty flag is true and the given pointer is nil, nothing will
	// be appended to the output encoder.
	//
	// Else, if the null flag is true and the given pointer is nil, a null value
	// will be appended to the output encoder.
	AddTriStateObject(
		key serial.Key,
		val gojay.MarshalerJSONObject,
		null bool,
	) EncWrapper
}

func NewEncWrapper(enc *gojay.Encoder) EncWrapper {
	return &encWrapper{enc}
}

type encWrapper struct {
	enc *gojay.Encoder
}

func (e *encWrapper) AddBool(key serial.Key, val bool) EncWrapper {
	e.enc.AddBoolKey(string(key), val)
	return e
}

func (e *encWrapper) AddOptionalBool(key serial.Key, val *bool) EncWrapper {
	if val != nil {
		e.enc.AddBoolKey(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddNullableBool(key serial.Key, val *bool) EncWrapper {
	if val == nil {
		e.enc.AddNullKey(string(key))
	} else {
		e.enc.AddBoolKey(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddTriStateBool(
	key serial.Key,
	val *bool,
	null bool,
) EncWrapper {
	if val == nil {
		if null {
			e.enc.AddNullKey(string(key))
		}
	} else {
		e.enc.AddBoolKey(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddString(key serial.Key, val string) EncWrapper {
	e.enc.AddStringKey(string(key), val)
	return e
}

func (e *encWrapper) AddOptionalString(key serial.Key, val *string) EncWrapper {
	if val != nil {
		e.enc.AddStringKey(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddNullableString(key serial.Key, val *string) EncWrapper {
	if val == nil {
		e.enc.AddNullKey(string(key))
	} else {
		e.enc.AddStringKey(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddTriStateString(
	key serial.Key,
	val *string,
	null bool,
) EncWrapper {
	if val == nil {
		if null {
			e.enc.AddNullKey(string(key))
		}
	} else {
		e.enc.AddStringKey(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddUint8(key serial.Key, val uint8) EncWrapper {
	e.enc.AddUint8Key(string(key), val)
	return e
}

func (e *encWrapper) AddOptionalUint8(key serial.Key, val *uint8) EncWrapper {
	if val != nil {
		e.enc.AddUint8Key(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddNullableUint8(key serial.Key, val *uint8) EncWrapper {
	if val == nil {
		e.enc.AddNullKey(string(key))
	} else {
		e.enc.AddUint8Key(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddTriStateUint8(
	key serial.Key,
	val *uint8,
	null bool,
) EncWrapper {
	if val == nil {
		if null {
			e.enc.AddNullKey(string(key))
		}
	} else {
		e.enc.AddUint8Key(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddUint16(key serial.Key, val uint16) EncWrapper {
	e.enc.AddUint16Key(string(key), val)
	return e
}

func (e *encWrapper) AddOptionalUint16(key serial.Key, val *uint16) EncWrapper {
	if val != nil {
		e.enc.AddUint16Key(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddNullableUint16(key serial.Key, val *uint16) EncWrapper {
	if val == nil {
		e.enc.AddNullKey(string(key))
	} else {
		e.enc.AddUint16Key(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddTriStateUint16(
	key serial.Key,
	val *uint16,
	null bool,
) EncWrapper {
	if val == nil {
		if null {
			e.enc.AddNullKey(string(key))
		}
	} else {
		e.enc.AddUint16Key(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddUint32(key serial.Key, val uint32) EncWrapper {
	e.enc.AddUint32Key(string(key), val)
	return e
}

func (e *encWrapper) AddOptionalUint32(key serial.Key, val *uint32) EncWrapper {
	if val != nil {
		e.enc.AddUint32Key(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddNullableUint32(key serial.Key, val *uint32) EncWrapper {
	if val == nil {
		e.enc.AddNullKey(string(key))
	} else {
		e.enc.AddUint32Key(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddTriStateUint32(
	key serial.Key,
	val *uint32,
	null bool,
) EncWrapper {
	if val == nil {
		if null {
			e.enc.AddNullKey(string(key))
		}
	} else {
		e.enc.AddUint32Key(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddUint64(key serial.Key, val uint64) EncWrapper {
	e.enc.AddUint64Key(string(key), val)
	return e
}

func (e *encWrapper) AddOptionalUint64(key serial.Key, val *uint64) EncWrapper {
	if val != nil {
		e.enc.AddUint64Key(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddNullableUint64(key serial.Key, val *uint64) EncWrapper {
	if val == nil {
		e.enc.AddNullKey(string(key))
	} else {
		e.enc.AddUint64Key(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddTriStateUint64(
	key serial.Key,
	val *uint64,
	null bool,
) EncWrapper {
	if val == nil {
		if null {
			e.enc.AddNullKey(string(key))
		}
	} else {
		e.enc.AddUint64Key(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddSnowflake(
	key serial.Key,
	val discord.Snowflake,
) EncWrapper {
	e.enc.AddStringKey(string(key), val.String())
	return e
}

func (e *encWrapper) AddOptionalSnowflake(
	key serial.Key,
	val discord.Snowflake,
) EncWrapper {
	if val != nil {
		e.enc.AddStringKey(string(key), val.String())
	}

	return e
}

func (e *encWrapper) AddNullableSnowflake(
	key serial.Key,
	val discord.Snowflake,
) EncWrapper {
	if val == nil {
		e.enc.AddNullKey(string(key))
	} else {
		e.enc.AddStringKey(string(key), val.String())
	}

	return e
}

func (e *encWrapper) AddTriStateSnowflake(
	key serial.Key,
	val discord.Snowflake,
	null bool,
) EncWrapper {
	if val == nil {
		if null {
			e.enc.AddNullKey(string(key))
		}
	} else {
		e.enc.AddStringKey(string(key), val.String())
	}

	return e
}

func (e *encWrapper) AddObject(
	key serial.Key,
	val gojay.MarshalerJSONObject,
) EncWrapper {
	e.enc.AddObjectKey(string(key), val)
	return e
}

func (e *encWrapper) AddOptionalObject(
	key serial.Key,
	val gojay.MarshalerJSONObject,
) EncWrapper {
	if val != nil {
		e.enc.AddObjectKey(string(key), val)
	}

	return e
}

func (e *encWrapper) AddNullableObject(
	key serial.Key,
	val gojay.MarshalerJSONObject,
) EncWrapper {
	if val == nil {
		e.enc.AddNullKey(string(key))
	} else {
		e.enc.AddObjectKey(string(key), val)
	}

	return e
}

func (e *encWrapper) AddTriStateObject(
	key serial.Key,
	val gojay.MarshalerJSONObject,
	null bool,
) EncWrapper {
	if val == nil {
		if null {
			e.enc.AddNullKey(string(key))
		}
	} else {
		e.enc.AddObjectKey(string(key), val)
	}

	return e
}

func (e *encWrapper) AddSlice(
	key serial.Key,
	val gojay.MarshalerJSONArray,
) EncWrapper {
	e.enc.AddArrayKey(string(key), val)
	return e
}

func (e *encWrapper) AddOptionalSlice(
	key serial.Key,
	value gojay.MarshalerJSONArray,
	empty bool,
) EncWrapper {
	// if slice is nil
	if value == nil {

		// and marked as empty
		if empty {

			// do nothing
			return e
		}

		// and _not_ marked as empty
		// add constructed empty slice
		return e.AddSlice(key, emptySlice{})
	}

	return e.AddSlice(key, value)
}

func (e *encWrapper) AddNullableSlice(
	key serial.Key,
	value gojay.MarshalerJSONArray,
	null bool,
) EncWrapper {
	if value == nil {
		if null {
			e.enc.AddNullKey(string(key))
			return e
		}

		return e.AddSlice(key, emptySlice{})
	}

	return e.AddSlice(key, value)
}

func (e *encWrapper) AddTriStateSlice(
	key serial.Key,
	value gojay.MarshalerJSONArray,
	empty bool,
	null bool,
) EncWrapper {
	if value == nil {
		if empty {
			return e
		} else if null {
			e.enc.AddNullKey(string(key))
			return e
		} else {
			return e.AddSlice(key, emptySlice{})
		}
	}

	return e.AddSlice(key, value)
}

func (e *encWrapper) AddInt64(key serial.Key, val int64) EncWrapper {
	e.enc.AddInt64Key(string(key), val)
	return e
}

func (e *encWrapper) AddOptionalInt64(key serial.Key, val *int64) EncWrapper {
	if val != nil {
		e.enc.AddInt64Key(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddNullableInt64(key serial.Key, val *int64) EncWrapper {
	if val == nil {
		e.enc.AddNullKey(string(key))
	} else {
		e.enc.AddInt64Key(string(key), *val)
	}

	return e
}

func (e *encWrapper) AddTriStateInt64(key serial.Key, val *int64, null bool) EncWrapper {
	if null {
		return e.AddNullableInt64(key, val)
	} else {
		return e.AddOptionalInt64(key, val)
	}
}

func (e *encWrapper) AddStringer(key serial.Key, val fmt.Stringer) EncWrapper {
	return e.AddString(key, val.String())
}

func (e *encWrapper) AddOptionalStringer(
	key serial.Key,
	val fmt.Stringer,
) EncWrapper {
	if val != nil {
		e.enc.AddStringKey(string(key), val.String())
	}

	return e
}

func (e *encWrapper) AddNullableStringer(
	key serial.Key,
	val fmt.Stringer,
) EncWrapper {
	if val != nil {
		e.enc.AddStringKey(string(key), val.String())
	} else {
		e.enc.AddNullKey(string(key))
	}

	return e
}

type emptySlice []uint8

func (emptySlice) MarshalJSONArray(*gojay.Encoder) {}

func (emptySlice) IsNil() bool {
	return false
}
