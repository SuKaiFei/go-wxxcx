package biz

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/json-iterator/go"
	errors2 "github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"time"
)

type ChatGPTRepo interface {
	Get(context.Context, string, string) (*ChatGPT, error)
	Add(context.Context, *ChatGPT) error
	CountByOpenid(context.Context, string) (int64, error)
}

type ChatGPTUseCase struct {
	repo          ChatGPTRepo
	log           *log.Helper
	chatGptApikey string
}

func NewChatGPTUseCase(repo ChatGPTRepo, appConf *conf.Application, logger log.Logger) *ChatGPTUseCase {
	return &ChatGPTUseCase{
		repo:          repo,
		chatGptApikey: appConf.GetChatGptApikey(),
		log:           log.NewHelper(logger),
	}
}

func (uc *ChatGPTUseCase) isLimit(ctx context.Context, openid string) (bool, error) {
	if openid == "oeAvF5bWne0rjIGB0MuPrGumQoBQ" {
		return false, nil
	}

	c, err := uc.repo.CountByOpenid(ctx, openid)
	if err != nil {
		return false, errors2.WithStack(err)
	}
	return c >= 5, nil
}

func (uc *ChatGPTUseCase) Completions(ctx context.Context, appid, openid, prompt string) (string, error) {
	isLimit, err := uc.isLimit(ctx, openid)
	if err != nil {
		return "", errors2.WithStack(err)
	}
	if isLimit {
		uc.log.Warnf("completions limit openid(%s)", openid)
		return "你的今日次数已用完\n点击右上角三个点，可以免费分享给好友使用", nil
	}

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.SetRequestURI("https://api.openai.com/v1/completions")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", uc.chatGptApikey))
	request.Header.SetContentType("application/json")
	request.Header.SetMethod(fasthttp.MethodPost)

	bodyObj := make(map[string]interface{})
	bodyObj["model"] = "text-davinci-003"
	bodyObj["prompt"] = prompt
	bodyObj["max_tokens"] = 2048
	bodyObj["user"] = fmt.Sprintf("%s_%s", appid, openid)

	body, err := jsoniter.Marshal(bodyObj)
	if err != nil {
		return "", errors2.WithStack(err)
	}

	request.SetBody(body)
	err = fasthttp.DoTimeout(request, response, 40*time.Second)
	if err != nil {
		return "", errors2.WithStack(err)
	}

	var resp *CompletionsResp
	if err = json.Unmarshal(response.Body(), &resp); err != nil {
		return "", errors2.WithStack(err)
	}
	if resp.Error.Message != "" {
		return "", errors2.New(resp.Error.Message)
	}

	uc.log.Infof("completions req(%+v) resp(%s)", bodyObj, response.Body())

	go uc.AddCompletionToDB(appid, openid, prompt, resp)

	return resp.Choices[0].Text, nil
}

func (uc *ChatGPTUseCase) AddCompletionToDB(appid, openid, prompt string, completionsResp *CompletionsResp) {
	hash := md5.New()
	hash.Write([]byte(prompt))

	c := &ChatGPT{
		Appid:  appid,
		Openid: openid,
		Code:   fmt.Sprintf("%x", hash.Sum(nil)),
		Prompt: prompt,
		Result: completionsResp.Choices[0].Text,
	}
	if err := uc.repo.Add(context.Background(), c); err != nil {
		uc.log.Errorf("uc.repo.Add(%+v)", err)
	}
}
