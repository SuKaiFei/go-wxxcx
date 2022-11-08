package service

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
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
				Appid: "wx575f5d87fb66e69a",
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

func TestServiceGetVoiceByID(t *testing.T) {
	Convey("GetVoiceById", t, func() {
		var (
			ctx = context.Background()
			in  = &v1.GetVoiceByIdRequest{
				Appid: "wx575f5d87fb66e69a",
				Id:    0,
			}
		)
		Convey("When everything goes positive", func() {
			p1, err := tSVC.voiceSvc.GetVoiceById(ctx, in)
			t.Logf("%+v", p1)
			Convey("Then err should be nil.p1 should not be nil.", func() {
				So(err, ShouldBeNil)
				So(p1, ShouldNotBeNil)
			})
		})
	})
}

type MongoVoice struct {
	ID     string  `json:"_id"`
	Code   string  `json:"code"`
	Name   string  `json:"name"`
	Sort   float64 `json:"sort"`
	Status float64 `json:"status"`
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

func TestServiceAddVoices(t *testing.T) {
	var (
		ctx = context.Background()
		in  *biz.Voice
	)
	fi, err := os.Open("/Users/sukaifei/WeChatProjects/media/jyh-media/db/导出20221107/music_navigation.json")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer func() { _ = fi.Close() }()

	br := bufio.NewReader(fi)
	for {
		var mv *MongoVoice
		a, err := br.ReadString('\n')
		if err != nil {
			t.Fatal(err.Error())
			return
		}
		if err = json.Unmarshal([]byte(a), &mv); err != nil {
			t.Fatal(err.Error())
			return
		}

		for i, work := range mv.Works {
			for j, voiceWork := range work {
				if voiceWork.TpUrl != "" {
					continue
				}
				if voiceWork.Name == "夺回秋雅" {
					t.Log(1)
				}
				if !strings.Contains(voiceWork.Id, "https://") {
					mv.Works[i][j].TpUrl = fmt.Sprintf("https://api.wxxcx.top/static/music/%s/%s", mv.Code, voiceWork.Id)
				} else {
					substr := "/" + mv.Code + "/"
					mv.Works[i][j].Id = voiceWork.Id[strings.LastIndex(voiceWork.Id, substr)+len(substr):]
					mv.Works[i][j].TpUrl = fmt.Sprintf("https://api.wxxcx.top/static/music/%s/%s", mv.Code, mv.Works[i][j].Id)
				}
			}
		}

		in = &biz.Voice{
			Appid:         "wx575f5d87fb66e69a",
			Code:          mv.Code,
			Name:          mv.Name,
			Default:       true,
			Type:          biz.VoiceTypeInner,
			Sort:          1,
			ShareTitle:    mv.Share.Title,
			ShareImageUrl: mv.Share.ImageURL,
			Works:         mv.Works,
		}

		//marshal, _ := json.Marshal(in)
		//t.Logf("%s", marshal)
		//continue
		err = tSVC.voiceSvc.vuc.Add(ctx, in)
		if err != nil {
			t.Logf("%+v", in)
			t.Fatal(err)
		}
	}
}
