package user_test

import (
	"testing"

	. "github.com/foxcapades/lib-go-discord/pkg/discord/user"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUserImpl_MarshalJSON(t *testing.T) {
	Convey("User.MarshalJSON", t, func() {
		Convey("returns an error when", func() {
			Convey("validation is enabled and", func() {
				Convey("the `id` field value is 0", func() {
					target := NewUser(true).
						SetUsername("hello").
						SetDiscriminator("1234")

					a, b := target.MarshalJSON()

					So(a, ShouldBeNil)
					So(b, ShouldEqual, ErrNoId)
				})
			})
		})
	})
}
