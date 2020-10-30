package com_test

import (
	"bytes"
	"github.com/foxcapades/lib-go-discord/v0/internal/types/com"
	"testing"

	"github.com/francoispqt/gojay"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDecodeSnowflake(t *testing.T) {
	Convey("DecodeSnowflake", t, func() {
		Convey("returns an error", func() {
			Convey("when the input value is not a string", func() {
				dec := gojay.NewDecoder(bytes.NewBufferString(`123`))

				a, b := com.DecodeSnowflake(dec)
				So(a, ShouldBeNil)
				So(b, ShouldNotBeNil)
			})

			Convey("when the input value is not a numeric string", func() {
				dec := gojay.NewDecoder(bytes.NewBufferString(`"hello"`))

				a, b := com.DecodeSnowflake(dec)
				So(a, ShouldBeNil)
				So(b, ShouldNotBeNil)
			})
		})

		Convey("returns a new Snowflake", func() {
			Convey("when given a numeric string value", func() {
				dec := gojay.NewDecoder(bytes.NewBufferString(`"123"`))

				a, b := com.DecodeSnowflake(dec)
				So(a, ShouldNotBeNil)
				So(b, ShouldBeNil)
				So(a.RawValue(), ShouldEqual, 123)
			})
		})

		Convey("returns a nil value", func() {
			Convey("when given a json null input value", func() {
				dec := gojay.NewDecoder(bytes.NewBufferString(`null`))

				a, b := com.DecodeSnowflake(dec)
				So(a, ShouldBeNil)
				So(b, ShouldBeNil)
			})
		})
	})
}
