package biz

import (
	"context"
	"fmt"
	errors2 "github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"gorm.io/gorm"
	"time"
)

func (uc *CommunityUseCase) taskSendUserTitle() {
	ctx := context.Background()
	c := cron.New(cron.WithSeconds())
	// 每天上午1:01触发
	EntryID, err := c.AddFunc("0 15 0 * * ?", func() {
		uc.log.Infow("msg", "sendUserTitleTemplateMessage start")
		openids, err := uc.repo.GetSendUserTitle(ctx)
		if err != nil {
			uc.log.Errorw("msg", "GetSendUserTitle", "err", err)
			return
		}
		for _, openid := range openids {
			if err := uc.sendUserTitleTemplateMessage(ctx, openid); err != nil {
				uc.log.Errorw("msg", "sendUserTitleTemplateMessage", "openid", openid, "err", err)
				continue
			}
			uc.log.Infow("msg", "sendUserTitleTemplateMessage", "openid", openid)
		}
		uc.log.Infow("msg", "sendUserTitleTemplateMessage end")
	})
	if err != nil {
		panic(err)
	}
	uc.log.Infow("msg", "taskSendUserTitle cronAddFunc", "EntryID", EntryID, "err", err)

	c.Start()
}

func (uc *CommunityUseCase) sendUserTitleTemplateMessage(ctx context.Context, openid string) error {
	first := "你昨天作品上了热门，免费送个专属头衔以资鼓励"
	wxUser, err := uc.wechatUc.GetUser(ctx, AppidCommunity, openid)
	if err != nil {
		return errors2.WithStack(err)
	}
	if wxUser.Unionid == "" {
		return errors2.New("unionid is empty")
	}

	const value = "火鸡"
	userTitle, err := uc.repo.GetUserTitleByValue(ctx, openid, value)
	const addTime = 48 * time.Hour
	if err != nil {
		if errors2.Cause(err) == gorm.ErrRecordNotFound {
			userTitle = &CommunityUserTitle{
				Openid:         openid,
				Value:          value,
				Class:          ".badge-orange",
				ValidityPeriod: time.Now().Add(addTime),
			}
			_, err := uc.repo.AddUserTitle(ctx, userTitle)
			if err != nil {
				return errors2.WithStack(err)
			}
		} else {
			return errors2.WithStack(err)
		}
	} else {
		if time.Now().After(userTitle.ValidityPeriod) {
			userTitle.ValidityPeriod = time.Now()
		}
		userTitle.ValidityPeriod = userTitle.ValidityPeriod.Add(addTime)
		if err := uc.repo.UpdateUserTitle(ctx, userTitle.ID, userTitle); err != nil {
			return errors2.WithStack(err)
		}
	}
	if err := uc.repo.UpdateUserUserTitle(ctx, userTitle); err != nil {
		return errors2.WithStack(err)
	}

	const templateID = "B5SBjFuSUcJcGivgIzOyLMQMDLl33Gf_16rhC1XxRtk"
	var msg CommunityNoticeMessage

	data := make(map[string]*message.TemplateDataItem)
	data["first"] = &message.TemplateDataItem{Value: first}
	data["keyword1"] = &message.TemplateDataItem{Value: "专属头衔"}
	data["keyword2"] = &message.TemplateDataItem{Value: fmt.Sprintf("加48小时，还剩%d小时", int(userTitle.ValidityPeriod.Sub(time.Now()).Hours()))}
	data["remark"] = &message.TemplateDataItem{Value: "点击卡片设置专属头衔"}
	msg.TemplateMessage = message.TemplateMessage{
		TemplateID: templateID,
		Data:       data,
	}
	msg.MiniProgram.AppID = AppidCommunity
	msg.MiniProgram.PagePath = fmt.Sprintf("pages/mine/edit")

	errMsg := "ok"
	msgID, err := uc.wechatUc.SendTemplateMsg(ctx, wxUser.Unionid, &msg.TemplateMessage)
	if err != nil {
		errMsg = err.Error()
	}

	m := &CommunityNoticeHistory{
		Openid:     wxUser.Openid,
		Unionid:    wxUser.Unionid,
		ReqOpenid:  "system",
		ReqUnionid: "system",
		TemplateID: templateID,
		MsgID:      uint64(msgID),
		Detail:     msg,
		ErrMsg:     errMsg,
	}
	if _, err := uc.repo.AddNoticeHistory(ctx, m); err != nil {
		return errors2.WithStack(err)
	}

	return nil
}
