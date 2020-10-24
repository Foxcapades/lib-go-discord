package discord_test

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"strconv"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNullableInt8
	Convey("NullableInt8

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target discord.NullableInt8

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target discord.NullableInt8

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(120)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target discord.NullableInt8

					target.Set(120)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target discord.NullableInt8

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target discord.NullableInt8

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(120)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target discord.NullableInt8

					target.Set(120)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target discord.NullableInt8

				target.Set(120)

				So(target.Get(),ShouldEqual,120)
			})
			// []

			Convey("panics when no value is set", func() {
				var target discord.NullableInt8

				So(func() {target.Get()}, ShouldPanicWith, discord.ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target discord.NullableInt8

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target discord.NullableInt8

				target.Set(120)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"120")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableInt16
	Convey("NullableInt16

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target discord.NullableInt16

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target discord.NullableInt16

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(32_215)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target discord.NullableInt16

					target.Set(32_215)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target discord.NullableInt16

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target discord.NullableInt16

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(32_215)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target discord.NullableInt16

					target.Set(32_215)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target discord.NullableInt16

				target.Set(32_215)

				So(target.Get(),ShouldEqual,32_215)
			})
			// []

			Convey("panics when no value is set", func() {
				var target discord.NullableInt16

				So(func() {target.Get()}, ShouldPanicWith, discord.ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target discord.NullableInt16

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target discord.NullableInt16

				target.Set(32_215)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"32215")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableInt32
	Convey("NullableInt32

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target discord.NullableInt32

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target discord.NullableInt32

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(3_331_115)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target discord.NullableInt32

					target.Set(3_331_115)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target discord.NullableInt32

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target discord.NullableInt32

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(3_331_115)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target discord.NullableInt32

					target.Set(3_331_115)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target discord.NullableInt32

				target.Set(3_331_115)

				So(target.Get(),ShouldEqual,3_331_115)
			})
			// []

			Convey("panics when no value is set", func() {
				var target discord.NullableInt32

				So(func() {target.Get()}, ShouldPanicWith, discord.ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target discord.NullableInt32

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target discord.NullableInt32

				target.Set(3_331_115)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"3331115")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableInt64
	Convey("NullableInt64

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target discord.NullableInt64

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target discord.NullableInt64

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(123_456_789_123)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target discord.NullableInt64

					target.Set(123_456_789_123)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target discord.NullableInt64

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target discord.NullableInt64

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(123_456_789_123)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target discord.NullableInt64

					target.Set(123_456_789_123)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target discord.NullableInt64

				target.Set(123_456_789_123)

				So(target.Get(),ShouldEqual,123_456_789_123)
			})
			// []

			Convey("panics when no value is set", func() {
				var target discord.NullableInt64

				So(func() {target.Get()}, ShouldPanicWith, discord.ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target discord.NullableInt64

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target discord.NullableInt64

				target.Set(123_456_789_123)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"123456789123")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableUint8
	Convey("NullableUint8

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target discord.NullableUint8

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target discord.NullableUint8

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(250)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target discord.NullableUint8

					target.Set(250)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target discord.NullableUint8

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target discord.NullableUint8

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(250)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target discord.NullableUint8

					target.Set(250)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target discord.NullableUint8

				target.Set(250)

				So(target.Get(),ShouldEqual,250)
			})
			// []

			Convey("panics when no value is set", func() {
				var target discord.NullableUint8

				So(func() {target.Get()}, ShouldPanicWith, discord.ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target discord.NullableUint8

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target discord.NullableUint8

				target.Set(250)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"250")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableUint16
	Convey("NullableUint16

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target discord.NullableUint16

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target discord.NullableUint16

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(62_000)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target discord.NullableUint16

					target.Set(62_000)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target discord.NullableUint16

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target discord.NullableUint16

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(62_000)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target discord.NullableUint16

					target.Set(62_000)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target discord.NullableUint16

				target.Set(62_000)

				So(target.Get(),ShouldEqual,62_000)
			})
			// []

			Convey("panics when no value is set", func() {
				var target discord.NullableUint16

				So(func() {target.Get()}, ShouldPanicWith, discord.ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target discord.NullableUint16

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target discord.NullableUint16

				target.Set(62_000)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"62000")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableUint32
	Convey("NullableUint32

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target discord.NullableUint32

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target discord.NullableUint32

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(10_880_000)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target discord.NullableUint32

					target.Set(10_880_000)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target discord.NullableUint32

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target discord.NullableUint32

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(10_880_000)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target discord.NullableUint32

					target.Set(10_880_000)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target discord.NullableUint32

				target.Set(10_880_000)

				So(target.Get(),ShouldEqual,10_880_000)
			})
			// []

			Convey("panics when no value is set", func() {
				var target discord.NullableUint32

				So(func() {target.Get()}, ShouldPanicWith, discord.ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target discord.NullableUint32

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target discord.NullableUint32

				target.Set(10_880_000)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"10880000")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableUint64
	Convey("NullableUint64

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target discord.NullableUint64

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target discord.NullableUint64

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(999_999_999_999)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target discord.NullableUint64

					target.Set(999_999_999_999)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target discord.NullableUint64

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target discord.NullableUint64

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(999_999_999_999)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target discord.NullableUint64

					target.Set(999_999_999_999)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target discord.NullableUint64

				target.Set(999_999_999_999)

				So(target.Get(),ShouldEqual,999_999_999_999)
			})
			// []

			Convey("panics when no value is set", func() {
				var target discord.NullableUint64

				So(func() {target.Get()}, ShouldPanicWith, discord.ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target discord.NullableUint64

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target discord.NullableUint64

				target.Set(999_999_999_999)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"999999999999")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableFloat32
	Convey("NullableFloat32

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target discord.NullableFloat32

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target discord.NullableFloat32

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(1.234512)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target discord.NullableFloat32

					target.Set(1.234512)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target discord.NullableFloat32

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target discord.NullableFloat32

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(1.234512)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target discord.NullableFloat32

					target.Set(1.234512)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target discord.NullableFloat32

				target.Set(1.234512)

				So(target.Get(),ShouldAlmostEqual,1.234512, 0.00001)
			})
			// [0.00001]

			Convey("panics when no value is set", func() {
				var target discord.NullableFloat32

				So(func() {target.Get()}, ShouldPanicWith, discord.ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target discord.NullableFloat32

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target discord.NullableFloat32

				target.Set(1.234512)

				a, b := target.MarshalJSON()


				So(func() float32 {v, e := strconv.ParseFloat(string(a), 32); if e != nil {panic(e)}; return float32(v)}(),ShouldAlmostEqual,1.234512, 0.00001)
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableFloat64
	Convey("NullableFloat64

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target discord.NullableFloat64

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target discord.NullableFloat64

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(12345.123456789)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target discord.NullableFloat64

					target.Set(12345.123456789)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target discord.NullableFloat64

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target discord.NullableFloat64

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(12345.123456789)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target discord.NullableFloat64

					target.Set(12345.123456789)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target discord.NullableFloat64

				target.Set(12345.123456789)

				So(target.Get(),ShouldAlmostEqual,12345.123456789)
			})
			// []

			Convey("panics when no value is set", func() {
				var target discord.NullableFloat64

				So(func() {target.Get()}, ShouldPanicWith, discord.ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target discord.NullableFloat64

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target discord.NullableFloat64

				target.Set(12345.1234567891)

				a, b := target.MarshalJSON()


				So(func() float64 {v, e := strconv.ParseFloat(string(a), 64); if e != nil {panic(e)}; return v}(),ShouldAlmostEqual,12345.1234567891)
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableBool
	Convey("NullableBool

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target discord.NullableBool

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target discord.NullableBool

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(true)
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target discord.NullableBool

					target.Set(true)

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target discord.NullableBool

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target discord.NullableBool

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(true)
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target discord.NullableBool

					target.Set(true)

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target discord.NullableBool

				target.Set(true)

				So(target.Get(),ShouldEqual,true)
			})
			// []

			Convey("panics when no value is set", func() {
				var target discord.NullableBool

				So(func() {target.Get()}, ShouldPanicWith, discord.ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target discord.NullableBool

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target discord.NullableBool

				target.Set(true)

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"true")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableString
	Convey("NullableString

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target discord.NullableString

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target discord.NullableString

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set("hello world")
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target discord.NullableString

					target.Set("hello world")

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target discord.NullableString

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target discord.NullableString

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set("hello world")
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target discord.NullableString

					target.Set("hello world")

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target discord.NullableString

				target.Set("hello world")

				So(target.Get(),ShouldEqual,"hello world")
			})
			// []

			Convey("panics when no value is set", func() {
				var target discord.NullableString

				So(func() {target.Get()}, ShouldPanicWith, discord.ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target discord.NullableString

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target discord.NullableString

				target.Set("hello world")

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"\"hello world\"")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableSnowflake
	Convey("NullableSnowflake

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target discord.NullableSnowflake

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target discord.NullableSnowflake

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(func() discord.Snowflake {out := discord.Snowflake{}; out.SetRawValue(123_456_789_123); return out}())
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target discord.NullableSnowflake

					target.Set(func() discord.Snowflake {out := discord.Snowflake{}; out.SetRawValue(123_456_789_123); return out}())

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target discord.NullableSnowflake

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target discord.NullableSnowflake

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(func() discord.Snowflake {out := discord.Snowflake{}; out.SetRawValue(123_456_789_123); return out}())
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target discord.NullableSnowflake

					target.Set(func() discord.Snowflake {out := discord.Snowflake{}; out.SetRawValue(123_456_789_123); return out}())

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target discord.NullableSnowflake

				target.Set(func() discord.Snowflake {out := discord.Snowflake{}; out.SetRawValue(123_456_789_123); return out}())

				So(target.Get().RawValue(),ShouldEqual,123456789123)
			})
			// []

			Convey("panics when no value is set", func() {
				var target discord.NullableSnowflake

				So(func() {target.Get()}, ShouldPanicWith, discord.ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target discord.NullableSnowflake

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target discord.NullableSnowflake

				target.Set(func() discord.Snowflake {out := discord.Snowflake{}; out.SetRawValue(123_456_789_123); return out}())

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"123456789123")
				So(b, ShouldBeNil)
			})
		})
	})
}

func TestNullableTime
	Convey("NullableTime

		Convey(".IsNull()", func() {
			Convey("returns true", func() {
				Convey("when first constructed", func() {
					var target discord.NullableTime

					So(target.IsNull(), ShouldBeTrue)
				})

				Convey("when set to null", func() {
					var target discord.NullableTime

					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)

					target.Set(time.Unix(0, 0).UTC())
					target.SetNull()

					So(target.IsNull(), ShouldBeTrue)
				})
			})

			Convey("returns false", func() {
				Convey("when a value is set", func() {
					var target discord.NullableTime

					target.Set(time.Unix(0, 0).UTC())

					So(target.IsNull(), ShouldBeFalse)
				})
			})
		})

		Convey(".IsNotNull()", func() {
			Convey("returns false", func() {
				Convey("when first constructed", func() {
					var target discord.NullableTime

					So(target.IsNotNull(), ShouldBeFalse)
				})

				Convey("when set to null", func() {
					var target discord.NullableTime

					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)

					target.Set(time.Unix(0, 0).UTC())
					target.SetNull()

					So(target.IsNotNull(), ShouldBeFalse)
				})
			})

			Convey("returns true", func() {
				Convey("when a value is set", func() {
					var target discord.NullableTime

					target.Set(time.Unix(0, 0).UTC())

					So(target.IsNotNull(), ShouldBeTrue)
				})
			})
		})

		Convey(".Get()", func() {
			Convey("returns the set value", func() {
				var target discord.NullableTime

				target.Set(time.Unix(0, 0).UTC())

				So(target.Get(),ShouldResemble,time.Unix(0, 0).UTC())
			})
			// []

			Convey("panics when no value is set", func() {
				var target discord.NullableTime

				So(func() {target.Get()}, ShouldPanicWith, discord.ErrNullField)
			})
		})

		Convey(".MarshalJSON()", func() {
			Convey("returns null when the field is null", func() {
				var target discord.NullableTime

				a, b := target.MarshalJSON()

				So(a, ShouldResemble, []byte{'n', 'u', 'l', 'l'})
				So(b, ShouldBeNil)
			})

			Convey("returns the serialized value when the field is not null", func() {
				var target discord.NullableTime

				target.Set(time.Unix(0, 0).UTC())

				a, b := target.MarshalJSON()


				So(string(a),ShouldEqual,"\"1970-01-01T00:00:00Z\"")
				So(b, ShouldBeNil)
			})
		})
	})
}
