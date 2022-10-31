package service

import (
	"context"
	"testing"

	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServiceGetBqbIndex(t *testing.T) {
	Convey("GetBqbIndex", t, func() {
		var (
			ctx = context.Background()
			in  = &v1.GetBqbIndexRequest{
				Appid: "wxd02e65b36dd61734",
			}
		)
		Convey("When everything goes positive", func() {
			p1, err := tSVC.bqbSvc.GetBqbIndex(ctx, in)
			t.Logf("%+v", p1)
			Convey("Then err should be nil.p1 should not be nil.", func() {
				So(err, ShouldBeNil)
				So(p1, ShouldNotBeNil)
			})
		})
	})
}

func TestServiceGetBqbList(t *testing.T) {
	Convey("GetBqbList", t, func() {
		var (
			ctx = context.Background()
			in  = &v1.GetBqbListRequest{
				Appid: "wxd02e65b36dd61734",
				Type:  "cxk",
			}
		)
		Convey("When everything goes positive", func() {
			p1, err := tSVC.bqbSvc.GetBqbList(ctx, in)
			t.Logf("%+v", p1)
			Convey("Then err should be nil.p1 should not be nil.", func() {
				So(err, ShouldBeNil)
				So(p1, ShouldNotBeNil)
			})
		})
	})
}
