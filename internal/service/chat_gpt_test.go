package service

import (
	"context"
	"testing"

	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServiceGetChatGptCompletions(t *testing.T) {
	Convey("GetChatGptCompletions", t, func() {
		var (
			ctx = context.Background()
			in  = &v1.GetChatGptCompletionsRequest{
				Appid: "wx575f5d87fb66e69a",
			}
		)
		Convey("When everything goes positive", func() {
			p1, err := tSVC.chatGptSvc.GetChatGptCompletions(ctx, in)
			t.Logf("%+v", p1)
			Convey("Then err should be nil.p1 should not be nil.", func() {
				So(err, ShouldBeNil)
				So(p1, ShouldNotBeNil)
			})
		})
	})
}
