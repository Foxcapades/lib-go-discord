package discord_test

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUserImpl_MarshalJSON(t *testing.T) {
	Convey("User.MarshalJSON", t, func() {
		Convey("returns an error when", func() {
			Convey("validation is enabled and", func() {
				Convey("the `id` field value is 0", func() {
					target := discord.NewUser(true).
						SetUsername("hello").
						SetDiscriminator(1234)

					a, b := target.MarshalJSON()

					So(a, ShouldBeNil)
					So(b, ShouldEqual, discord.ErrNoId)
				})

				Convey("the `username` field value is empty", func() {
					target := discord.NewUser(true).
						SetID(dlib.Snowflake{}).
						SetDiscriminator(1234)

					a, b := target.MarshalJSON()

					So(a, ShouldBeNil)
					So(b, ShouldEqual, discord.ErrNoId)
				})

				Convey("the `discriminator` field value is 0", func() {
					target := discord.NewUser(true).
						SetID(dlib.Snowflake{}).
						SetDiscriminator(1234)

					a, b := target.MarshalJSON()

					So(a, ShouldBeNil)
					So(b, ShouldEqual, discord.ErrNoId)
				})
			})
		})
	})
}