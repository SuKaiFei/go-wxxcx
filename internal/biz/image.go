package biz

import (
	"bytes"
	"context"
	"fmt"
	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/nfnt/resize"
	errors2 "github.com/pkg/errors"
	"image"
	"image/jpeg"
	"image/png"
	"time"
)

type ImageUseCase struct {
	log       *log.Helper
	cosUC     *CosUseCase
	whitelist map[string]struct{}
}

func NewImageUseCase(logger log.Logger, cosUC *CosUseCase) *ImageUseCase {
	whitelist := map[string]struct{}{
		"":                             {},
		"oLbjy5U8441c3vmPZdaXcZb4s5h8": {},
		"oLbjy5QvnIhStDV6Ot2KBiOicorI": {},
		"oLbjy5Q3Bb3fu2zFCJXYN0UnqNEs": {},
	}
	return &ImageUseCase{whitelist: whitelist, cosUC: cosUC, log: log.NewHelper(logger)}
}

func (uc *ImageUseCase) UploadImage(ctx context.Context, req *v1.UploadImageRequest) (*v1.UploadImageReply, error) {
	if len(req.File) == 0 {
		return nil, errors.New(400, "", "file is empty")
	}

	img, filetype, _ := image.Decode(bytes.NewReader(req.File))
	if img != nil {
		size := len(req.File)

		if size >= 10*1024*1024 {
			return nil, errors.New(400, "", "图片多大")
		}
		var (
			maxWidth, maxHeight, imgSize uint
			isWidth                      bool
		)

		point := img.Bounds().Size()
		if point.X > point.Y {
			isWidth = true
			imgSize = uint(point.X)
		} else {
			imgSize = uint(point.Y)
		}
		if imgSize > 1080 {
			if isWidth {
				maxWidth = 1080
			} else {
				maxHeight = 1080
			}
		}

		b := new(bytes.Buffer)
		imgResize := resize.Resize(maxWidth, maxHeight, img, resize.NearestNeighbor)

		if filetype == "png" {
			if errImg := png.Encode(b, imgResize); errImg == nil {
				req.File = b.Bytes()
			}
		} else {
			if errImg := jpeg.Encode(b, imgResize, nil); errImg == nil {
				req.File = b.Bytes()
			}
		}

		uc.log.Infow(
			"msg", "imageResize",
			"压缩前", size/1024*1024,
			"压缩后", len(req.File)/1024*1024,
		)
	}
	if _, found := uc.whitelist[req.Openid]; !found {
		if img != nil {
			bmp, err := gozxing.NewBinaryBitmapFromImage(img)
			if err != nil {
				return nil, err
			}

			result, _ := qrcode.NewQRCodeReader().Decode(bmp, nil)
			if result != nil && len(result.GetText()) > 0 {
				return nil, errors2.New("禁止二维码图片")
			}
		}
	}
	now := time.Now()
	if req.Code == "" {
		req.Code = "article"
	}
	key := fmt.Sprintf(`%s/%0.4d%0.2d%0.2d/%s.%s`, req.Code, now.Year(), now.Month(), now.Day(), uuid.NewString(), filetype)
	upload, err := uc.cosUC.Upload(ctx, key, bytes.NewReader(req.File))
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	return &v1.UploadImageReply{Path: upload}, nil
}
