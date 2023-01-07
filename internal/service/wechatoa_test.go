package service

import (
	"encoding/json"
	"testing"
)

var (
	appid      = "wx9ef62ba2e3525812"                          // 鸡你太美小助手
	appidJyh   = "wx575f5d87fb66e69a"                          // 鸡音盒
	appidZJ    = "wx3b3021e25e5a3ce1"                          // 指尖坤巴
	appidGame  = "wx4cc11e69406205ff"                          // 鸡音盒
	tagIDBoss  = 100                                           // 鸡你太美小助手TagID
	templateID = "3EHnd170OYZ7ucWq1cOdRdeGGzrXD0u73QkTYHZHc4o" // 成绩更新提醒
)

var (
	openidA   = "ol8Ng53_-RSSiCdr4JXdKqz-gJho" // 🦊
	openidB   = "ol8Ng5_zLCtOEVKRSEk5AQiNvkPg" // jyh
	openidC   = "ol8Ng5xppYhcFrtss8zeEjZep2SE" // emo
	openidAll = []string{openidA, openidB, openidC}
)

func TestGetUserList(t *testing.T) {
	userinfoList, err := tSVC.wechatOcSvc.GetUserList(appid)
	if err != nil {
		t.Fatal(err)
	}
	for _, info := range userinfoList {
		marshal, _ := json.Marshal(info)
		t.Logf("%s", marshal)
	}
}

func TestSend(t *testing.T) {
	openIDs := openidAll
	for _, s := range openidAll {
		tSVC.wechatOcSvc.SendAsync(s)
	}
	t.Logf("推送完成(%d)\n", len(openIDs))
}

func TestSendLocalhost(t *testing.T) {
	for i, openID := range openidAll {
		msgID, err := tSVC.wechatOcSvc.Send(openID)
		t.Logf("index(%d) msgID(%d) error(%+v)\n", i, msgID, err)
	}

	t.Log("推送完成")
}

func TestSendAll(t *testing.T) {
	//_, client := tSVC.wechatOcSvc.GetApp(appid)
	//openIDs, err := client.GetUser().ListAllUserOpenIDs()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//for _, openID := range openIDs {
	//	tSVC.wechatOcSvc.SendAsync(openID)
	//}
	//
	//t.Logf("推送完成(%d)\n", len(openIDs))
}
