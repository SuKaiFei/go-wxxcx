package biz

import (
	"context"
	"fmt"
	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	cosSts "github.com/SuKaiFei/go-wxxcx/util/cos"
	"github.com/go-kratos/kratos/v2/log"
	jsoniter "github.com/json-iterator/go"
	errors2 "github.com/pkg/errors"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type CosUseCase struct {
	log       *log.Helper
	cosConf   *conf.Application_Cos
	stsClient *cosSts.Client
	client    *cos.Client
	bucketURL string
}

func NewCosUseCase(appConf *conf.Application, logger log.Logger) *CosUseCase {
	cosConf := appConf.GetCos()
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", cosConf.Bucket, cosConf.Region))
	ciu, _ := url.Parse(fmt.Sprintf("https://%s.ci.%s.myqcloud.com", cosConf.Bucket, cosConf.Region))
	// 用于Get Service 查询，默认全地域 service.cos.myqcloud.com
	su, _ := url.Parse("https://service.cos.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su, CIURL: ciu}

	return &CosUseCase{
		cosConf:   cosConf,
		log:       log.NewHelper(logger),
		bucketURL: u.String(),
		stsClient: cosSts.NewClient(cosConf.SecretId, cosConf.SecretKey, nil),
		client: cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  cosConf.SecretId,
				SecretKey: cosConf.SecretKey,
			},
		}),
	}
}

func (uc *CosUseCase) Upload(ctx context.Context, key string, file io.Reader) (string, error) {
	put, err := uc.client.Object.Put(ctx, key, file, nil)
	if err != nil {
		return "", errors2.WithStack(err)
	}
	if put.StatusCode != http.StatusOK {
		return "", errors2.New("上传异常")
	}

	return uc.GetPresignedURL(ctx, key)
}

func (uc *CosUseCase) GetPresignedURL(ctx context.Context, key string) (string, error) {
	if strings.Contains(key, uc.bucketURL) {
		urlEnd := strings.Index(key, "?")
		if urlEnd <= 0 {
			urlEnd = len(key)
		}
		key = key[len(uc.bucketURL)+1 : urlEnd]
	}
	if key[:4] == "http" {
		return key, nil
	}
	presignedURL, err := uc.client.Object.GetPresignedURL(ctx, http.MethodGet, key, uc.cosConf.SecretId, uc.cosConf.SecretKey, time.Minute, nil)
	if err != nil {
		return "", errors2.WithStack(err)
	}

	return presignedURL.String(), nil
}

func (uc *CosUseCase) GetPresignedURLByPhoto(ctx context.Context, photos []*v1.Photo) ([]*v1.Photo, error) {
	for _, photo := range photos {
		if strings.Contains(photo.Url, uc.bucketURL) {
			presignedURL, err := uc.GetPresignedURL(ctx, photo.Url)
			if err != nil {
				return nil, errors2.WithStack(err)
			}
			photo.Url = presignedURL
		}
	}

	return photos, nil
}

func (uc *CosUseCase) BatchImageAuditing(ctx context.Context, images []string) (bool, error) {
	options := make([]cos.ImageAuditingInputOptions, len(images))
	for i, image := range images {
		options[i] = cos.ImageAuditingInputOptions{
			DataId: strconv.Itoa(i),
		}
		image = strings.ReplaceAll(image, uc.bucketURL, "")
		if image[:4] == "http" {
			options[i].Url = image
		} else {
			end := strings.Index(image, "?")
			if end <= 0 {
				end = len(image)
			}
			image = image[:end]
			options[i].Object = image
		}
	}
	opt := &cos.BatchImageAuditingOptions{
		Input: options,
	}
	res, _, err := uc.client.CI.BatchImageAuditing(ctx, opt)
	if err != nil {
		return false, errors2.WithStack(err)
	}

	errCount := 0
	for _, result := range res.JobsDetail {
		if result.Label != "Normal" {
			errCount++
			r, _ := jsoniter.MarshalToString(result)
			uc.log.Warnw(
				"msg", "BatchImageAuditing",
				"result", r,
			)
		}
	}

	return errCount == 0, nil
}

func (uc *CosUseCase) GetTempCredentials() (res *cosSts.CredentialResult, err error) {
	// 策略概述 https://cloud.tencent.com/document/product/436/18023
	opt := &cosSts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          uc.cosConf.Region,
		Policy: &cosSts.CredentialPolicy{
			Statement: []cosSts.CredentialPolicyStatement{
				{
					// 密钥的权限列表。简单上传和分片需要以下的权限，其他权限列表请看 https://cloud.tencent.com/document/product/436/31923
					Action: []string{
						// 简单上传
						"name/cos:PostObject",
						"name/cos:PutObject",
						"name/cos:PutObject",
						"name/ci:CreateAuditingPictureJob",
						"name/ci:DescribeAuditingPictureFiles",
					},
					Effect: "allow",
					Resource: []string{
						// 这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						// 存储桶的命名格式为 BucketName-APPID，此处填写的 bucket 必须为此格式
						fmt.Sprintf("qcs::cos:%s:uid/%d:%s/article/*", uc.cosConf.Region, uc.cosConf.Appid, uc.cosConf.Bucket),
						fmt.Sprintf("qcs::cos:%s:uid/%d:%s/avatar/*", uc.cosConf.Region, uc.cosConf.Appid, uc.cosConf.Bucket),
						fmt.Sprintf("qcs::ci:%s:uid/%d:%s/article/*", uc.cosConf.Region, uc.cosConf.Appid, uc.cosConf.Bucket),
						fmt.Sprintf("qcs::ci:%s:uid/%d:%s/avatar/*", uc.cosConf.Region, uc.cosConf.Appid, uc.cosConf.Bucket),
					},
				},
			},
		},
	}

	return uc.stsClient.GetCredential(opt)
}
