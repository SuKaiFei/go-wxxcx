package service

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"

	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	. "github.com/smartystreets/goconvey/convey"
)

func TestServiceGetVoiceList(t *testing.T) {
	Convey("GetVoiceList", t, func() {
		var (
			ctx = context.Background()
			in  = &v1.GetVoiceListRequest{
				Appid: "wxd02e65b36dd61734",
				Type:  "cxk",
			}
		)
		Convey("When everything goes positive", func() {
			p1, err := tSVC.voiceSvc.vuc.GetList(ctx, in.GetAppid())
			t.Logf("%+v", p1)
			Convey("Then err should be nil.p1 should not be nil.", func() {
				So(err, ShouldBeNil)
				So(p1, ShouldNotBeNil)
			})
		})
	})
}

func TestServiceGetVoiceDefault(t *testing.T) {
	Convey("GetVoiceDefault", t, func() {
		var (
			ctx = context.Background()
			in  = &v1.GetVoiceDefaultRequest{
				Appid: "wxf0628bd9092bd9dd",
			}
		)
		Convey("When everything goes positive", func() {
			p1, err := tSVC.voiceSvc.GetVoiceDefault(ctx, in)
			t.Logf("%+v", p1)
			Convey("Then err should be nil.p1 should not be nil.", func() {
				So(err, ShouldBeNil)
				So(p1, ShouldNotBeNil)
			})
		})
	})
}

type MongoVoice struct {
	ID     string `json:"_id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Sort   int    `json:"sort"`
	Status int    `json:"status"`
	Share  struct {
		Title    string `json:"title"`
		ImageURL string `json:"imageUrl"`
	} `json:"share"`
	Works biz.VoiceWorks `json:"works"`
}

func TestServiceAddVoice(t *testing.T) {
	var (
		ctx = context.Background()
		mv  *MongoVoice
		in  *biz.Voice
	)
	bytes, err := os.ReadFile("/Users/sukaifei/WeChatProjects/media/时代马戏团/db/时代马戏团.json")
	if err != nil {
		t.Fatal(err)
	}
	if err = json.Unmarshal(bytes, &mv); err != nil {
		return
	}

	in = &biz.Voice{
		Appid:         "wxc41448e80cf2bd82",
		Code:          mv.Code,
		Name:          mv.Name,
		Default:       true,
		Type:          biz.VoiceTypeInner,
		MpAppid:       "wx575f5d87fb66e69a",
		Sort:          1,
		ShareTitle:    mv.Share.Title,
		ShareImageUrl: mv.Share.ImageURL,
		Works:         mv.Works,
	}

	err = tSVC.voiceSvc.vuc.Add(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
}
