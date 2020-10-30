package guild_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	. "github.com/foxcapades/lib-go-discord/v0/internal/types/guild"
)

func TestNewWidget(t *testing.T) {
	Convey("guild.NewWidget()", t, func() {
		Convey("does not return nil", func() {
			So(NewWidget(), ShouldNotBeNil)
		})
	})
}

