package biz

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	pb "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/json-iterator/go"
	errors2 "github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
	"strconv"
	"time"
)

const (
	defaultImagePath = "https://mmbiz.qpic.cn/mmbiz_jpg/Mhr8pCDXpQqoWjx3avyHfIMn9OJc93Pobae5GkdGX5E93SOBrichibgCNKxEn3JyCqeBBNmYkA5SHricEtyfMhs8A/0?wx_fmt=jpeg"
)

type ChatGPTRepo interface {
	Get(context.Context, uint) (*ChatGPT, error)
	Add(context.Context, *ChatGPT) (uint, error)
	CountByOpenid(context.Context, string) (int64, error)
	Search(context.Context, string, string) (*ChatGPT, error)

	GetTodayQuota(context.Context, string) (*ChatGPTQuota, error)
	AddQuotaUseCount(context.Context, string) error
	AddQuotaUnusedCount(context.Context, string) error
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
	//if openid == "oeAvF5bWne0rjIGB0MuPrGumQoBQ" {
	//	return false, nil
	//}

	c, err := uc.GetTodaySendCount(ctx, openid)
	if err != nil {
		return false, errors2.WithStack(err)
	}
	return c < 1, nil
}

func (uc *ChatGPTUseCase) searchByDB(ctx context.Context, prompt string) (*ChatGPT, error) {
	hash := md5.New()
	hash.Write([]byte(prompt))

	m, err := uc.repo.Search(ctx, fmt.Sprintf("%x", hash.Sum(nil)), prompt)
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return m, nil
}

func (uc *ChatGPTUseCase) GetChatGptHistory(ctx context.Context, openid string, id uint64) (*ChatGPT, error) {
	m, err := uc.repo.Get(ctx, uint(id))
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	if m.Type == ChatGPTTypeQuote {
		pid, _ := strconv.Atoi(m.Result)
		if pid > 0 {
			pm, err := uc.repo.Get(ctx, uint(pid))
			if err != nil {
				return nil, errors2.WithStack(err)
			}
			m.Result = pm.Result
		}
	}

	uc.log.Infof("GetChatGptHistory invited %s->%s", m.Openid, openid)

	return m, nil
}

func (uc *ChatGPTUseCase) GetTodaySendCount(ctx context.Context, openid string) (uint64, error) {
	m, err := uc.repo.GetTodayQuota(ctx, openid)
	if err != nil {
		return 0, errors2.WithStack(err)
	}
	return m.UnusedCount - m.UseCount, nil
}

func (uc *ChatGPTUseCase) AddQuotaUnusedCount(ctx context.Context, openid string) error {
	err := uc.repo.AddQuotaUnusedCount(ctx, openid)
	if err != nil {
		return errors2.WithStack(err)
	}
	return nil
}

func (uc *ChatGPTUseCase) Completions(ctx context.Context, appid, openid, prompt string) (*pb.GetChatGptCompletionsReply, error) {
	reply := &pb.GetChatGptCompletionsReply{
		ImagePath: defaultImagePath,
		Id:        0,
	}

	isLimit, err := uc.isLimit(ctx, openid)
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	if isLimit {
		uc.log.Warnf("completions limit openid(%s)", openid)
		reply.Result = "你的今日次数已用完\n点击右上角三个点，可以免费分享给好友使用"
		return reply, nil
	}

	m, sErr := uc.searchByDB(ctx, prompt)
	if sErr != nil && errors2.Cause(sErr) != gorm.ErrRecordNotFound {
		return nil, errors2.WithStack(sErr)
	}
	if m != nil {
		id, err := uc.AddCompletionToDB(ctx, appid, openid, prompt, strconv.Itoa(int(m.ID)), ChatGPTTypeQuote)
		if err != nil {
			return nil, errors2.WithStack(err)
		}
		reply.Result = m.Result
		reply.Id = uint64(id)
		return reply, nil
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
		return nil, errors2.WithStack(err)
	}

	request.SetBody(body)
	err = fasthttp.DoTimeout(request, response, 40*time.Second)
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	var resp *CompletionsResp
	if err = json.Unmarshal(response.Body(), &resp); err != nil {
		return nil, errors2.WithStack(err)
	}
	if resp.Error.Message != "" {
		return nil, errors2.New(resp.Error.Message)
	}

	uc.log.Infof("completions req(%+v) resp(%s)", bodyObj, response.Body())

	reply.Result = resp.Choices[0].Text
	id, err := uc.AddCompletionToDB(ctx, appid, openid, prompt, reply.Result, ChatGPTTypeReply)
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	reply.Id = uint64(id)

	return reply, nil
}

func (uc *ChatGPTUseCase) AddCompletionToDB(ctx context.Context, appid, openid, prompt, result string, typ ChatGPTType) (uint, error) {
	hash := md5.New()
	hash.Write([]byte(prompt))

	c := &ChatGPT{
		Appid:  appid,
		Openid: openid,
		Code:   fmt.Sprintf("%x", hash.Sum(nil)),
		Prompt: prompt,
		Result: result,
		Type:   typ,
	}
	if err := uc.repo.AddQuotaUseCount(ctx, openid); err != nil {
		return 0, errors2.WithStack(err)
	}

	return uc.repo.Add(ctx, c)
}
