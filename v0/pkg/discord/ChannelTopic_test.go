package discord_test

import (
	"bytes"
	"testing"

	"github.com/francoispqt/gojay"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDecodeChannelTopic(t *testing.T) {
	Convey("DecodeChannelTopic", t, func() {
		Convey("Should return nil", func() {
			Convey("when the input json value is null", func() {
				dec := gojay.NewDecoder(bytes.NewBufferString("null"))

				a, b := DecodeChannelTopic(dec)
				So(a, ShouldBeNil)
				So(b, ShouldBeNil)
			})
		})

		Convey("Should return an error", func() {
			Convey("when the given input json value is not a string or null", func() {
				dec := gojay.NewDecoder(bytes.NewBufferString("1234"))

				a, b := DecodeChannelTopic(dec)
				So(a, ShouldBeNil)
				So(b, ShouldNotBeNil)
			})
		})

		Convey("Should return a ChannelTopic pointer", func() {
			Convey("when the given input json value is a string", func() {
				dec := gojay.NewDecoder(bytes.NewBufferString(`"testing"`))

				a, b := DecodeChannelTopic(dec)
				So(a, ShouldNotBeNil)
				So(b, ShouldBeNil)
				So(*a, ShouldEqual, "testing")
			})
		})
	})
}
