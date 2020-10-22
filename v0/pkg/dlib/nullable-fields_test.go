package dlib_test

import (
	"strconv"
	"testing"
	"time"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
	. "github.com/smartystreets/goconvey/convey"
)

func TestI8NullableField(t *testing.T) {
	Convey("I8NullableField", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target I8NullableField

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target I8NullableField

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(120)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target I8NullableField

					target.Set(120)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target I8NullableField

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target I8NullableField

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(120)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target I8NullableField

					target.Set(120)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target I8NullableField

				target.Set(120)

				So(target.Get(),ShouldEqual,120)
			})
			// []

			Convey("panics when no value is set", func() {
				var target I8NullableField

				So(func() {target.Get()}, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target I8NullableField

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target I8NullableField

				target.Set(120)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"120")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestI16NullableField(t *testing.T) {
	Convey("I16NullableField", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target I16NullableField

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target I16NullableField

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(32_215)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target I16NullableField

					target.Set(32_215)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target I16NullableField

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target I16NullableField

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(32_215)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target I16NullableField

					target.Set(32_215)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target I16NullableField

				target.Set(32_215)

				So(target.Get(),ShouldEqual,32_215)
			})
			// []

			Convey("panics when no value is set", func() {
				var target I16NullableField

				So(func() {target.Get()}, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target I16NullableField

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target I16NullableField

				target.Set(32_215)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"32215")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestI32NullableField(t *testing.T) {
	Convey("I32NullableField", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target I32NullableField

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target I32NullableField

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(3_331_115)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target I32NullableField

					target.Set(3_331_115)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target I32NullableField

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target I32NullableField

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(3_331_115)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target I32NullableField

					target.Set(3_331_115)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target I32NullableField

				target.Set(3_331_115)

				So(target.Get(),ShouldEqual,3_331_115)
			})
			// []

			Convey("panics when no value is set", func() {
				var target I32NullableField

				So(func() {target.Get()}, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target I32NullableField

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target I32NullableField

				target.Set(3_331_115)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"3331115")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestI64NullableField(t *testing.T) {
	Convey("I64NullableField", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target I64NullableField

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target I64NullableField

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(123_456_789_123)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target I64NullableField

					target.Set(123_456_789_123)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target I64NullableField

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target I64NullableField

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(123_456_789_123)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target I64NullableField

					target.Set(123_456_789_123)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target I64NullableField

				target.Set(123_456_789_123)

				So(target.Get(),ShouldEqual,123_456_789_123)
			})
			// []

			Convey("panics when no value is set", func() {
				var target I64NullableField

				So(func() {target.Get()}, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target I64NullableField

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target I64NullableField

				target.Set(123_456_789_123)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"123456789123")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestU8NullableField(t *testing.T) {
	Convey("U8NullableField", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target U8NullableField

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target U8NullableField

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(250)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target U8NullableField

					target.Set(250)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target U8NullableField

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target U8NullableField

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(250)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target U8NullableField

					target.Set(250)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target U8NullableField

				target.Set(250)

				So(target.Get(),ShouldEqual,250)
			})
			// []

			Convey("panics when no value is set", func() {
				var target U8NullableField

				So(func() {target.Get()}, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target U8NullableField

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target U8NullableField

				target.Set(250)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"250")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestU16NullableField(t *testing.T) {
	Convey("U16NullableField", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target U16NullableField

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target U16NullableField

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(62_000)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target U16NullableField

					target.Set(62_000)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target U16NullableField

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target U16NullableField

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(62_000)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target U16NullableField

					target.Set(62_000)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target U16NullableField

				target.Set(62_000)

				So(target.Get(),ShouldEqual,62_000)
			})
			// []

			Convey("panics when no value is set", func() {
				var target U16NullableField

				So(func() {target.Get()}, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target U16NullableField

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target U16NullableField

				target.Set(62_000)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"62000")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestU32NullableField(t *testing.T) {
	Convey("U32NullableField", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target U32NullableField

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target U32NullableField

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(10_880_000)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target U32NullableField

					target.Set(10_880_000)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target U32NullableField

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target U32NullableField

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(10_880_000)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target U32NullableField

					target.Set(10_880_000)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target U32NullableField

				target.Set(10_880_000)

				So(target.Get(),ShouldEqual,10_880_000)
			})
			// []

			Convey("panics when no value is set", func() {
				var target U32NullableField

				So(func() {target.Get()}, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target U32NullableField

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target U32NullableField

				target.Set(10_880_000)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"10880000")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestU64NullableField(t *testing.T) {
	Convey("U64NullableField", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target U64NullableField

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target U64NullableField

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(999_999_999_999)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target U64NullableField

					target.Set(999_999_999_999)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target U64NullableField

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target U64NullableField

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(999_999_999_999)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target U64NullableField

					target.Set(999_999_999_999)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target U64NullableField

				target.Set(999_999_999_999)

				So(target.Get(),ShouldEqual,999_999_999_999)
			})
			// []

			Convey("panics when no value is set", func() {
				var target U64NullableField

				So(func() {target.Get()}, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target U64NullableField

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target U64NullableField

				target.Set(999_999_999_999)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"999999999999")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestF32NullableField(t *testing.T) {
	Convey("F32NullableField", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target F32NullableField

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target F32NullableField

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(1.234512)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target F32NullableField

					target.Set(1.234512)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target F32NullableField

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target F32NullableField

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(1.234512)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target F32NullableField

					target.Set(1.234512)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target F32NullableField

				target.Set(1.234512)

				So(target.Get(),ShouldAlmostEqual,1.234512, 0.00001)
			})
			// [0.00001]

			Convey("panics when no value is set", func() {
				var target F32NullableField

				So(func() {target.Get()}, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target F32NullableField

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target F32NullableField

				target.Set(1.234512)

				a, b := target.MarshalJSON()


				So(func() float32 {v, e := strconv.ParseFloat(string(a), 32); if e != nil {panic(e)}; return float32(v)}(),ShouldAlmostEqual,1.234512, 0.00001)
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestF64NullableField(t *testing.T) {
	Convey("F64NullableField", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target F64NullableField

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target F64NullableField

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(12345.123456789)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target F64NullableField

					target.Set(12345.123456789)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target F64NullableField

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target F64NullableField

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(12345.123456789)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target F64NullableField

					target.Set(12345.123456789)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target F64NullableField

				target.Set(12345.123456789)

				So(target.Get(),ShouldAlmostEqual,12345.123456789)
			})
			// []

			Convey("panics when no value is set", func() {
				var target F64NullableField

				So(func() {target.Get()}, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target F64NullableField

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target F64NullableField

				target.Set(12345.1234567891)

				a, b := target.MarshalJSON()


				So(func() float64 {v, e := strconv.ParseFloat(string(a), 64); if e != nil {panic(e)}; return v}(),ShouldAlmostEqual,12345.1234567891)
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestBoolNullableField(t *testing.T) {
	Convey("BoolNullableField", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target BoolNullableField

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target BoolNullableField

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(true)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target BoolNullableField

					target.Set(true)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target BoolNullableField

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target BoolNullableField

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(true)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target BoolNullableField

					target.Set(true)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target BoolNullableField

				target.Set(true)

				So(target.Get(),ShouldEqual,true)
			})
			// []

			Convey("panics when no value is set", func() {
				var target BoolNullableField

				So(func() {target.Get()}, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target BoolNullableField

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target BoolNullableField

				target.Set(true)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"true")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestStrNullableField(t *testing.T) {
	Convey("StrNullableField", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target StrNullableField

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target StrNullableField

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set("hello world")
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target StrNullableField

					target.Set("hello world")

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target StrNullableField

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target StrNullableField

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set("hello world")
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target StrNullableField

					target.Set("hello world")

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target StrNullableField

				target.Set("hello world")

				So(target.Get(),ShouldEqual,"hello world")
			})
			// []

			Convey("panics when no value is set", func() {
				var target StrNullableField

				So(func() {target.Get()}, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target StrNullableField

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target StrNullableField

				target.Set("hello world")

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"\"hello world\"")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestSnowflakeNullableField(t *testing.T) {
	Convey("SnowflakeNullableField", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target SnowflakeNullableField

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target SnowflakeNullableField

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(func() Snowflake {out := Snowflake{}; out.SetRawValue(123_456_789_123); return out}())
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target SnowflakeNullableField

					target.Set(func() Snowflake {out := Snowflake{}; out.SetRawValue(123_456_789_123); return out}())

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target SnowflakeNullableField

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target SnowflakeNullableField

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(func() Snowflake {out := Snowflake{}; out.SetRawValue(123_456_789_123); return out}())
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target SnowflakeNullableField

					target.Set(func() Snowflake {out := Snowflake{}; out.SetRawValue(123_456_789_123); return out}())

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target SnowflakeNullableField

				target.Set(func() Snowflake {out := Snowflake{}; out.SetRawValue(123_456_789_123); return out}())

				So(target.Get().RawValue(),ShouldEqual,123456789123)
			})
			// []

			Convey("panics when no value is set", func() {
				var target SnowflakeNullableField

				So(func() {target.Get()}, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target SnowflakeNullableField

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target SnowflakeNullableField

				target.Set(func() Snowflake {out := Snowflake{}; out.SetRawValue(123_456_789_123); return out}())

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"123456789123")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestTimeNullableField(t *testing.T) {
	Convey("TimeNullableField", t, func() {

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target TimeNullableField

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target TimeNullableField

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(time.Unix(0, 0).UTC())
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target TimeNullableField

					target.Set(time.Unix(0, 0).UTC())

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target TimeNullableField

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target TimeNullableField

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(time.Unix(0, 0).UTC())
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target TimeNullableField

					target.Set(time.Unix(0, 0).UTC())

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target TimeNullableField

				target.Set(time.Unix(0, 0).UTC())

				So(target.Get(),ShouldResemble,time.Unix(0, 0).UTC())
			})
			// []

			Convey("panics when no value is set", func() {
				var target TimeNullableField

				So(func() {target.Get()}, ShouldPanicWith, ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target TimeNullableField

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target TimeNullableField

				target.Set(time.Unix(0, 0).UTC())

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"\"1970-01-01T00:00:00Z\"")
				So(b, ShouldBeNil)
			})
		})
	})
}
