package service

import (
	"context"
	"testing"

	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServiceGetArticle(t *testing.T) {
	Convey("GetArticle", t, func() {
		var (
			ctx = context.Background()
			in  = &v1.GetArticleRequest{
				Appid: "wx575f5d87fb66e69a",
				Code:  "about",
			}
		)
		Convey("When everything goes positive", func() {
			p1, err := tSVC.articleSvc.GetArticle(ctx, in)
			t.Logf("%+v", p1)
			Convey("Then err should be nil.p1 should not be nil.", func() {
				So(err, ShouldBeNil)
				So(p1, ShouldNotBeNil)
			})
		})
	})
}

func TestServiceGetArticles(t *testing.T) {
	Convey("GetArticles", t, func() {
		var (
			ctx = context.Background()
			in  = &v1.GetArticlesRequest{
				Appid: "wx575f5d87fb66e69a",
				Code:  "yulu",
			}
		)
		Convey("When everything goes positive", func() {
			p1, err := tSVC.articleSvc.GetArticles(ctx, in)
			t.Logf("%+v", p1)
			Convey("Then err should be nil.p1 should not be nil.", func() {
				So(err, ShouldBeNil)
				So(p1, ShouldNotBeNil)
			})
		})
	})
}
