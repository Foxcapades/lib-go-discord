package user_test

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/user"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDiscriminator_Validate(t *testing.T) {
	Convey("Discriminator.Validate", t, func() {
		Convey("returns an error", func() {
			Convey("when the value is higher than 9999", func() {
				target := user.Discriminator(10_000)

				So(target.Validate(), ShouldEqual, user.ErrDiscriminatorLength)
			})

			Convey("when the value is lower than 1000", func() {
				target := user.Discriminator(999)

				So(target.Validate(), ShouldEqual, user.ErrDiscriminatorLength)
			})
		})

		Convey("returns nil", func() {
			Convey("when the value is in the range [1000, 9999]", func() {
				for i := 1000; i < 10_000; i++ {
					target := user.Discriminator(i)

					So(target.Validate(), ShouldBeNil)
				}
			})
		})
	})
}

func TestDiscriminator_IsValid(t *testing.T) {
	Convey("Discriminator.IsValid", t, func() {
		Convey("returns false", func() {
			Convey("when the value is higher than 9999", func() {
				target := user.Discriminator(10_000)

				So(target.IsValid(), ShouldBeFalse)
			})

			Convey("when the value is lower than 1000", func() {
				target := user.Discriminator(999)

				So(target.IsValid(), ShouldBeFalse)
			})
		})

		Convey("returns nil", func() {
			Convey("when the value is in the range [1000, 9999]", func() {
				for i := 1000; i < 10_000; i++ {
					target := user.Discriminator(i)

					So(target.IsValid(), ShouldBeTrue)
				}
			})
		})
	})
}

func TestDiscriminator_MarshalJSON(t *testing.T) {
	Convey("Discriminator.MarshalJSON", t, func() {
		Convey("returns the uint16 value as a json string", func() {
			target := user.Discriminator(1234)

			a, b := target.MarshalJSON()

			So(b, ShouldBeNil)
			So(a, ShouldResemble, []byte(`"1234"`))
		})
	})
}

func TestDiscriminator_UnmarshalJSON(t *testing.T) {
	Convey("Discriminator.UnmarshalJSON", t, func() {
		Convey("returns an error", func() {
			Convey("when the input value is not a string", func() {
				var target user.Discriminator

				So(target.UnmarshalJSON([]byte("1234")), ShouldNotBeNil)
			})

			Convey("when the input value string is not numeric", func() {
				var target user.Discriminator

				So(target.UnmarshalJSON([]byte("hello")), ShouldNotBeNil)
			})
		})
	})
}
