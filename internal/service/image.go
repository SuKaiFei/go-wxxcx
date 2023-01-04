package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"strings"
	"time"
)

type ImageService struct {
	v1.UnimplementedImageServer
	uc      *biz.ImageUseCase
	appConf *conf.Application
}

func NewImageService(uc *biz.ImageUseCase, appConf *conf.Application) *ImageService {
	return &ImageService{uc: uc, appConf: appConf}
}

func (s *ImageService) UploadImage(ctx context.Context, req *v1.UploadImageRequest) (
	reply *v1.UploadImageReply,
	err error,
) {
	if len(req.Openid) == 0 {
		return nil, errors.New(400, "", "wrong user identity")
	}
	return s.uc.UploadImage(ctx, req)
}

func (s *ImageService) UploadImageOld(ctx context.Context, req *v1.UploadImageRequest) (
	reply *v1.UploadImageReply,
	err error,
) {
	fileDir := fmt.Sprintf("/image/ikunzf/upload/%s", time.Now().Format("20060102"))
	mErr := os.MkdirAll(s.appConf.GetStaticPath()+fileDir, os.ModePerm)
	if mErr != nil {
		log.Warnw(
			"msg", "os.MkdirAll error",
			"error", mErr.Error(),
		)
	}
	filename := fmt.Sprintf("%s.%s", uuid.NewString(), strings.Split(req.GetFilename(), ".")[1])
	const sourceFileFlag = "source_"
	filePath := fmt.Sprintf("%s/%s%s", fileDir, sourceFileFlag, filename)
	f, err := os.OpenFile(s.appConf.GetStaticPath()+filePath, os.O_WRONLY|os.O_CREATE, 0o666)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	_, err = io.Copy(f, bytes.NewReader(req.GetFile()))
	if err != nil {
		return nil, err
	}

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)
	request.Header.SetMethod(fasthttp.MethodPost)
	request.SetRequestURI("https://jhgo.xahhp.com/api/image_process")
	request.SetHost("jhgo.xahhp.com")
	request.Header.SetUserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 11_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E217 MicroMessenger/6.8.0(0x16080000) NetType/WIFI Language/en Branch/Br_trunk MiniProgramEnv/Mac")
	request.Header.SetReferer("https://servicewechat.com/wxdb8936810cd717cc/372/page-frame.html")
	requestBody := new(bytes.Buffer)
	writer := multipart.NewWriter(requestBody)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			"file", filename))
	h.Set("Content-Type", http.DetectContentType(req.GetFile()))
	part, err := writer.CreatePart(h)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, bytes.NewReader(req.GetFile()))
	if err != nil {
		return nil, err
	}
	request.Header.SetContentType(writer.FormDataContentType())

	_ = writer.WriteField("name", filename)
	_ = writer.WriteField("a", "gpsimg")
	_ = writer.WriteField("para", fmt.Sprintf(`{"noauto":"","usekoutu":true,"bgimg":"https://img.alicdn.com/imgextra/i1/2201168165444/O1CN01gfyCWT1q5OyiKKrXM_!!2201168165444.jpg","bgimginfo":"230|30|370","pageid":"22809","randstr":"1%s"}`, strings.ReplaceAll(uuid.NewString(), "-", "")))
	writer.Close()
	request.SetBody(requestBody.Bytes())
	err = fasthttp.DoTimeout(request, response, 10*time.Second)
	if err != nil {
		return nil, err
	}

	req.File = nil

	var i *TidyImage
	_ = json.Unmarshal(response.Body(), &i)
	if i != nil && i.Detail.Personcut.Psimg != "" {
		_, body, _ := fasthttp.GetTimeout(nil, i.Detail.Personcut.Psimg, 5*time.Second)
		if len(body) > 0 {
			filePath := fmt.Sprintf("%s/%s", fileDir, filename)
			f, err := os.OpenFile(s.appConf.GetStaticPath()+filePath, os.O_WRONLY|os.O_CREATE, 0o666)
			if err != nil {
				return nil, err
			}
			defer f.Close()
			_, err = io.Copy(f, bytes.NewReader(body))
			if err != nil {
				return nil, err
			}
			return &v1.UploadImageReply{Path: fmt.Sprintf("https://api.wxxcx.top/static%s", filePath)}, nil
		}
	}

	path := fmt.Sprintf("https://api.wxxcx.top/static%s", filePath)
	return &v1.UploadImageReply{Path: path}, nil
}

type TidyImage struct {
	Code   int `json:"code"`
	Detail struct {
		Origin struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			Path   string `json:"path"`
		} `json:"origin"`
		Personcut struct {
			Psimg    string `json:"psimg"`
			Shareimg string `json:"shareimg"`
		} `json:"personcut"`
		DetectFace       interface{} `json:"DetectFace"`
		Headimg          interface{} `json:"headimg"`
		Faceswap         interface{} `json:"faceswap"`
		FaceswapShareImg string      `json:"faceswap_share_img"`
	} `json:"detail"`
	Msg string `json:"msg"`
}
