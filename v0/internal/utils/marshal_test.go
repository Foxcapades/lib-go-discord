package utils_test

import (
	"github.com/foxcapades/lib-go-discord/v0/internal/utils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUnmarshalString(t *testing.T) {
	Convey("UnmarshalString", t, func() {
		Convey("returns an error when", func() {
			Convey("the input value is empty", func() {
				a, b := utils.UnmarshalString([]byte{})
				So(a, ShouldEqual, "")
				So(b, ShouldEqual, utils.ErrTooShort)
			})

			Convey("the input value has only 1 character", func() {
				a, b := utils.UnmarshalString([]byte{'"'})
				So(a, ShouldEqual, "")
				So(b, ShouldEqual, utils.ErrTooShort)
			})

			Convey("the input value has no closing double quote", func() {
				a, b := utils.UnmarshalString([]byte{'"', 'H', 'i'})
				So(a, ShouldEqual, "")
				So(b, ShouldEqual, utils.ErrBadFmt)
			})

			Convey("the input value has no opening double quote", func() {
				a, b := utils.UnmarshalString([]byte{'H', 'i', '"'})
				So(a, ShouldEqual, "")
				So(b, ShouldEqual, utils.ErrBadFmt)
			})
		})

		Convey("returns the unwrapped string", func() {
			Convey("when given a valid JSON string", func() {
				a, b := utils.UnmarshalString([]byte(`"hi"`))
				So(a, ShouldEqual, "hi")
				So(b, ShouldBeNil)
			})
		})
	})
}
