package service

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServiceGetMusics(t *testing.T) {
	Convey("GetMusics", t, func() {
		var (
			ctx = context.Background()
			in  = &v1.GetMusicListRequest{
				Appid: "wx3b3021e25e5a3ce1",
				Code:  "index",
			}
		)
		Convey("When everything goes positive", func() {
			p1, err := tSVC.musicSvc.GetMusicList(ctx, in)
			t.Logf("%+v", p1)
			Convey("Then err should be nil.p1 should not be nil.", func() {
				So(err, ShouldBeNil)
				So(p1, ShouldNotBeNil)
			})
		})
	})
}

type MongoMusic struct {
	ID         string  `json:"_id"`
	Singer     string  `json:"singer"`
	Name       string  `json:"name"`
	Image      string  `json:"image"`
	ShareImage string  `json:"shareImage"`
	Duration   float64 `json:"duration"`
	Url        string  `json:"url"`
}

func TestServiceAddMusics(t *testing.T) {
	fi, err := os.Open("/Users/sukaifei/WeChatProjects/media/KK坤乐/database_export-ZqUdnx5g8dnU.json")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer func() { _ = fi.Close() }()

	br := bufio.NewReader(fi)
	const appid = "wx3b3021e25e5a3ce1"
	const code = "index"
	i := 0
	for {
		i++
		var mv *MongoMusic
		a, err := br.ReadString('\n')
		if err != nil {
			t.Fatal(err.Error())
			return
		}
		if err = json.Unmarshal([]byte(a), &mv); err != nil {
			t.Fatal(err.Error())
			return
		}
		insertSQL := fmt.Sprintf("INSERT INTO `wxxcx`.`music` "+
			"(`created_at`, `updated_at`, `appid`, `code`, `image_path`, `name`, `singer`, `url`, `duration`, `share_image_url`, `sort`) VALUES "+
			"(NOW(),NOW(),'%s','%s','%s','%s','%s','%s',%d,'%s',%d);", appid, code, mv.Image, mv.Name, mv.Singer, mv.Url, int(mv.Duration), mv.ShareImage, i)
		fmt.Println(insertSQL)
	}
}
