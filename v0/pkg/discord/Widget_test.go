package discord_test

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestWidgetImpl_UnmarshalJSON(t *testing.T) {
	Convey("Widget.UnmarshalJSON", t, func() {
		target1 := discord.NewWidget()
		data1 := `{"enabled": true, "channel_id": null}`

		err := target1.UnmarshalJSON([]byte(data1))
		So(err, ShouldBeNil)
		So(target1.Enabled(), ShouldBeTrue)
		So(target1.ChannelIDIsNull(), ShouldBeTrue)

		target2 := discord.NewWidget()
		data2 := `{"enabled": false, "channel_id": "123456"}`

		err = target2.UnmarshalJSON([]byte(data2))
		So(err, ShouldBeNil)
		So(target2.Enabled(), ShouldBeFalse)
		So(target2.ChannelIDIsNull(), ShouldBeFalse)
		So(target2.ChannelID().RawValue(), ShouldEqual, 123456)
	})
}

