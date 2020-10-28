package types_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/build"

	. "github.com/foxcapades/lib-go-discord/v0/internal/types"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNullableInt8(t *testing.T) {
	Convey("NullableInt8", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableInt8

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableInt8

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(120)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableInt8

					target.Set(120)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableInt8

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableInt8

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(120)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableInt8

					target.Set(120)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableInt8

				target.Set(120)

				So(target.Get(), ShouldEqual, 120)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableInt8

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableInt8

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableInt8

				target.Set(120)

				a, b := target.MarshalJSON()

				So(string(a), ShouldEqual, "120")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableInt16(t *testing.T) {
	Convey("NullableInt16", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableInt16

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableInt16

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(32_215)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableInt16

					target.Set(32_215)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableInt16

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableInt16

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(32_215)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableInt16

					target.Set(32_215)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableInt16

				target.Set(32_215)

				So(target.Get(), ShouldEqual, 32_215)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableInt16

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableInt16

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableInt16

				target.Set(32_215)

				a, b := target.MarshalJSON()

				So(string(a), ShouldEqual, "32215")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableInt32(t *testing.T) {
	Convey("NullableInt32", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableInt32

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableInt32

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(3_331_115)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableInt32

					target.Set(3_331_115)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableInt32

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableInt32

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(3_331_115)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableInt32

					target.Set(3_331_115)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableInt32

				target.Set(3_331_115)

				So(target.Get(), ShouldEqual, 3_331_115)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableInt32

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableInt32

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableInt32

				target.Set(3_331_115)

				a, b := target.MarshalJSON()

				So(string(a), ShouldEqual, "3331115")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableInt64(t *testing.T) {
	Convey("NullableInt64", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableInt64

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableInt64

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(123_456_789_123)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableInt64

					target.Set(123_456_789_123)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableInt64

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableInt64

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(123_456_789_123)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableInt64

					target.Set(123_456_789_123)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableInt64

				target.Set(123_456_789_123)

				So(target.Get(), ShouldEqual, 123_456_789_123)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableInt64

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableInt64

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableInt64

				target.Set(123_456_789_123)

				a, b := target.MarshalJSON()

				So(string(a), ShouldEqual, "123456789123")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableUint8(t *testing.T) {
	Convey("NullableUint8", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableUint8

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableUint8

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(250)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableUint8

					target.Set(250)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableUint8

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableUint8

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(250)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableUint8

					target.Set(250)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableUint8

				target.Set(250)

				So(target.Get(), ShouldEqual, 250)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableUint8

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableUint8

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableUint8

				target.Set(250)

				a, b := target.MarshalJSON()

				So(string(a), ShouldEqual, "250")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableUint16(t *testing.T) {
	Convey("NullableUint16", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableUint16

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableUint16

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(62_000)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableUint16

					target.Set(62_000)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableUint16

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableUint16

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(62_000)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableUint16

					target.Set(62_000)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableUint16

				target.Set(62_000)

				So(target.Get(), ShouldEqual, 62_000)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableUint16

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableUint16

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableUint16

				target.Set(62_000)

				a, b := target.MarshalJSON()

				So(string(a), ShouldEqual, "62000")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableUint32(t *testing.T) {
	Convey("NullableUint32", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableUint32

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableUint32

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(10_880_000)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableUint32

					target.Set(10_880_000)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableUint32

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableUint32

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(10_880_000)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableUint32

					target.Set(10_880_000)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableUint32

				target.Set(10_880_000)

				So(target.Get(), ShouldEqual, 10_880_000)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableUint32

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableUint32

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableUint32

				target.Set(10_880_000)

				a, b := target.MarshalJSON()

				So(string(a), ShouldEqual, "10880000")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableUint64(t *testing.T) {
	Convey("NullableUint64", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableUint64

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableUint64

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(999_999_999_999)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableUint64

					target.Set(999_999_999_999)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableUint64

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableUint64

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(999_999_999_999)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableUint64

					target.Set(999_999_999_999)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableUint64

				target.Set(999_999_999_999)

				So(target.Get(), ShouldEqual, 999_999_999_999)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableUint64

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableUint64

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableUint64

				target.Set(999_999_999_999)

				a, b := target.MarshalJSON()

				So(string(a), ShouldEqual, "999999999999")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableFloat32(t *testing.T) {
	Convey("NullableFloat32", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableFloat32

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableFloat32

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(1.234512)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableFloat32

					target.Set(1.234512)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableFloat32

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableFloat32

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(1.234512)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableFloat32

					target.Set(1.234512)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableFloat32

				target.Set(1.234512)

				So(target.Get(), ShouldAlmostEqual, 1.234512, 0.00001)
			})
			// [0.00001]

			Convey("panics when no value is set", func() {
				var target NullableFloat32

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableFloat32

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableFloat32

				target.Set(1.234512)

				a, b := target.MarshalJSON()

				So(func() float32 {
					v, e := strconv.ParseFloat(string(a), 32)
					if e != nil {
						panic(e)
					}
					return float32(v)
				}(), ShouldAlmostEqual, 1.234512, 0.00001)
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableFloat64(t *testing.T) {
	Convey("NullableFloat64", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableFloat64

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableFloat64

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(12345.123456789)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableFloat64

					target.Set(12345.123456789)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableFloat64

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableFloat64

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(12345.123456789)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableFloat64

					target.Set(12345.123456789)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableFloat64

				target.Set(12345.123456789)

				So(target.Get(), ShouldAlmostEqual, 12345.123456789)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableFloat64

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableFloat64

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableFloat64

				target.Set(12345.1234567891)

				a, b := target.MarshalJSON()

				So(func() float64 {
					v, e := strconv.ParseFloat(string(a), 64)
					if e != nil {
						panic(e)
					}
					return v
				}(), ShouldAlmostEqual, 12345.1234567891)
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableBool(t *testing.T) {
	Convey("NullableBool", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableBool

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableBool

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(true)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableBool

					target.Set(true)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableBool

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableBool

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(true)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableBool

					target.Set(true)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableBool

				target.Set(true)

				So(target.Get(), ShouldEqual, true)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableBool

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableBool

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableBool

				target.Set(true)

				a, b := target.MarshalJSON()

				So(string(a), ShouldEqual, "true")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableString(t *testing.T) {
	Convey("NullableString", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableString

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableString

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set("hello world")
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableString

					target.Set("hello world")

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableString

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableString

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set("hello world")
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableString

					target.Set("hello world")

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableString

				target.Set("hello world")

				So(target.Get(), ShouldEqual, "hello world")
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableString

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableString

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableString

				target.Set("hello world")

				a, b := target.MarshalJSON()

				So(string(a), ShouldEqual, "\"hello world\"")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableSnowflake(t *testing.T) {
	Convey("NullableSnowflake", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableSnowflake

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableSnowflake

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(func() discord.Snowflake {
						out := build.NewSnowflake(false)
						out.SetRawValue(123_456_789_123)
						return out
					}())
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableSnowflake

					target.Set(func() discord.Snowflake {
						out := build.NewSnowflake(false)
						out.SetRawValue(123_456_789_123)
						return out
					}())

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableSnowflake

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableSnowflake

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(func() discord.Snowflake {
						out := build.NewSnowflake(false)
						out.SetRawValue(123_456_789_123)
						return out
					}())
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableSnowflake

					target.Set(func() discord.Snowflake {
						out := build.NewSnowflake(false)
						out.SetRawValue(123_456_789_123)
						return out
					}())

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableSnowflake

				target.Set(func() discord.Snowflake {
					out := build.NewSnowflake(false)
					out.SetRawValue(123_456_789_123)
					return out
				}())

				So(target.Get().RawValue(), ShouldEqual, 123456789123)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableSnowflake

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableSnowflake

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableSnowflake

				target.Set(func() discord.Snowflake {
					out := build.NewSnowflake(false)
					out.SetRawValue(123_456_789_123)
					return out
				}())

				a, b := target.MarshalJSON()

				So(string(a), ShouldEqual, "123456789123")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableTime(t *testing.T) {
	Convey("NullableTime", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableTime

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableTime

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(time.Unix(0, 0).UTC())
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableTime

					target.Set(time.Unix(0, 0).UTC())

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableTime

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableTime

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(time.Unix(0, 0).UTC())
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableTime

					target.Set(time.Unix(0, 0).UTC())

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableTime

				target.Set(time.Unix(0, 0).UTC())

				So(target.Get(), ShouldResemble, time.Unix(0, 0).UTC())
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableTime

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableTime

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableTime

				target.Set(time.Unix(0, 0).UTC())

				a, b := target.MarshalJSON()

				So(string(a), ShouldEqual, "\"1970-01-01T00:00:00Z\"")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableVerificationLevel(t *testing.T) {
	Convey("NullableVerificationLevel", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableVerificationLevel

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableVerificationLevel

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(discord.VerificationLevelHigh)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableVerificationLevel

					target.Set(discord.VerificationLevelHigh)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableVerificationLevel

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableVerificationLevel

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(discord.VerificationLevelHigh)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableVerificationLevel

					target.Set(discord.VerificationLevelHigh)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableVerificationLevel

				target.Set(discord.VerificationLevelHigh)

				So(target.Get(), ShouldEqual, discord.VerificationLevelHigh)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableVerificationLevel

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableVerificationLevel

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableVerificationLevel

				target.Set(discord.VerificationLevelLow)

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'1'})
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableMessageNotificationLevel(t *testing.T) {
	Convey("NullableMessageNotificationLevel", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableMessageNotificationLevel

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableMessageNotificationLevel

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(discord.MsgNoteLvlOnlyMentions)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableMessageNotificationLevel

					target.Set(discord.MsgNoteLvlOnlyMentions)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableMessageNotificationLevel

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableMessageNotificationLevel

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(discord.MsgNoteLvlOnlyMentions)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableMessageNotificationLevel

					target.Set(discord.MsgNoteLvlOnlyMentions)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableMessageNotificationLevel

				target.Set(discord.MsgNoteLvlOnlyMentions)

				So(target.Get(), ShouldEqual, discord.MsgNoteLvlOnlyMentions)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableMessageNotificationLevel

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableMessageNotificationLevel

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableMessageNotificationLevel

				target.Set(discord.MsgNoteLvlAllMessages)

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'0'})
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableExplicitContentFilterLevel(t *testing.T) {
	Convey("NullableExplicitContentFilterLevel", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableExplicitContentFilterLevel

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableExplicitContentFilterLevel

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(discord.ExpConFilterLvlDisabled)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableExplicitContentFilterLevel

					target.Set(discord.ExpConFilterLvlDisabled)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableExplicitContentFilterLevel

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableExplicitContentFilterLevel

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(discord.ExpConFilterLvlDisabled)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableExplicitContentFilterLevel

					target.Set(discord.ExpConFilterLvlDisabled)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableExplicitContentFilterLevel

				target.Set(discord.ExpConFilterLvlDisabled)

				So(target.Get(), ShouldEqual, discord.ExpConFilterLvlDisabled)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableExplicitContentFilterLevel

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableExplicitContentFilterLevel

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableExplicitContentFilterLevel

				target.Set(discord.ExpConFilterLvlMembersWithoutRoles)

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'1'})
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableChannelTopic(t *testing.T) {
	Convey("NullableChannelTopic", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableChannelTopic

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableChannelTopic

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set("testing")
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableChannelTopic

					target.Set("testing")

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableChannelTopic

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableChannelTopic

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set("testing")
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableChannelTopic

					target.Set("testing")

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableChannelTopic

				target.Set("testing")

				So(target.Get(), ShouldEqual, discord.ChannelTopic("testing"))
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableChannelTopic

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableChannelTopic

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableChannelTopic

				target.Set("testing")

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte("\"testing\""))
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableAny(t *testing.T) {
	Convey("NullableAny", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableAny

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableAny

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(nil)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableAny

					target.Set(nil)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableAny

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableAny

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(nil)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableAny

					target.Set(nil)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableAny

				target.Set(nil)

				So(target.Get(), ShouldEqual, nil)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableAny

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableAny

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableAny

				target.Set(nil)

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte("null"))
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableActivityEmoji(t *testing.T) {
	Convey("NullableActivityEmoji", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableActivityEmoji

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableActivityEmoji

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(nil)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableActivityEmoji

					target.Set(nil)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableActivityEmoji

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableActivityEmoji

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(nil)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableActivityEmoji

					target.Set(nil)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableActivityEmoji

				target.Set(nil)

				So(target.Get(), ShouldEqual, nil)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableActivityEmoji

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableActivityEmoji

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableActivityEmoji

				target.Set(nil)

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte("null"))
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableActivityParty(t *testing.T) {
	Convey("NullableActivityParty", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableActivityParty

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableActivityParty

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(nil)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableActivityParty

					target.Set(nil)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableActivityParty

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableActivityParty

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(nil)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableActivityParty

					target.Set(nil)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableActivityParty

				target.Set(nil)

				So(target.Get(), ShouldEqual, nil)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableActivityParty

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableActivityParty

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableActivityParty

				target.Set(nil)

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte("null"))
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableActivityAssets(t *testing.T) {
	Convey("NullableActivityAssets", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableActivityAssets

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableActivityAssets

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(nil)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableActivityAssets

					target.Set(nil)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableActivityAssets

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableActivityAssets

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(nil)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableActivityAssets

					target.Set(nil)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableActivityAssets

				target.Set(nil)

				So(target.Get(), ShouldEqual, nil)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableActivityAssets

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableActivityAssets

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableActivityAssets

				target.Set(nil)

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte("null"))
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableActivitySecrets(t *testing.T) {
	Convey("NullableActivitySecrets", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target NullableActivitySecrets

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target NullableActivitySecrets

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(nil)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target NullableActivitySecrets

					target.Set(nil)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target NullableActivitySecrets

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target NullableActivitySecrets

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(nil)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target NullableActivitySecrets

					target.Set(nil)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target NullableActivitySecrets

				target.Set(nil)

				So(target.Get(), ShouldEqual, nil)
			})
			// []

			Convey("panics when no value is set", func() {
				var target NullableActivitySecrets

				So(func() { target.Get() }, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target NullableActivitySecrets

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target NullableActivitySecrets

				target.Set(nil)

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte("null"))
				So(b, ShouldBeNil)
			})
		})
	})
}
