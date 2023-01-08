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
	"image/color"
	"image/jpeg"
	"image/png"
	"time"
)

type ImageUseCase struct {
	log       *log.Helper
	cosUC     *CosUseCase
	wechatUC  *WechatUseCase
	whitelist map[string]struct{}
}

func NewImageUseCase(logger log.Logger, cosUC *CosUseCase, wechatUC *WechatUseCase) *ImageUseCase {
	whitelist := map[string]struct{}{
		"":                             {},
		"oLbjy5U8441c3vmPZdaXcZb4s5h8": {},
		"oLbjy5QvnIhStDV6Ot2KBiOicorI": {},
		"oLbjy5Q3Bb3fu2zFCJXYN0UnqNEs": {},
	}
	return &ImageUseCase{whitelist: whitelist, cosUC: cosUC, wechatUC: wechatUC, log: log.NewHelper(logger)}
}

func (uc *ImageUseCase) UploadImage(ctx context.Context, req *v1.UploadImageRequest) (*v1.UploadImageReply, error) {
	if len(req.File) == 0 {
		return nil, errors.New(400, "", "file is empty")
	}

	img, filetype, _ := image.Decode(bytes.NewReader(req.File))
	if img != nil {
		size := len(req.File)

		if size >= 10*1024*1024 {
			return nil, errors.New(400, "", "图片太大")
		}
		var (
			point               = img.Bounds().Size()
			maxWidth, maxHeight uint
			maxSize             = 1280
			//maxWidth, maxHeight, imgSize uint
			//isWidth                      bool
		)

		//point := img.Bounds().Size()
		//if point.X > point.Y {
		//	isWidth = true
		//	imgSize = uint(point.X)
		//} else {
		//	imgSize = uint(point.Y)
		//}
		//var maxSize uint = 1280
		//if imgSize > maxSize {
		//	if isWidth {
		//		maxWidth = maxSize
		//	} else {
		//		maxHeight = maxSize
		//	}
		//}

		if point.X > maxSize {
			maxWidth = uint(maxSize)
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
		content := uc.GetQRCodeContent(img)
		if len(content) > 0 {
			return nil, errors2.New("禁止二维码图片")
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

func (uc *ImageUseCase) GetQRCodeContent(img image.Image) string {
	if img != nil {
		bmp, err := gozxing.NewBinaryBitmapFromImage(img)
		if err != nil {
			return ""
		}
		reader := qrcode.NewQRCodeReader()
		result, _ := reader.Decode(bmp, nil)
		if result != nil && len(result.GetText()) > 0 {
			return result.GetText()
		}

		grayingImage := uc.ToBlackAndWhiteImage(img)
		bmp, err = gozxing.NewBinaryBitmapFromImage(grayingImage)
		if err != nil {
			return ""
		}
		result, _ = reader.Decode(bmp, nil)
		if result != nil && len(result.GetText()) > 0 {
			return result.GetText()
		}

		grayingImage = uc.ReverseImageColor(img)
		bmp, err = gozxing.NewBinaryBitmapFromImage(grayingImage)
		if err != nil {
			return ""
		}
		result, _ = reader.Decode(bmp, nil)
		if result != nil && len(result.GetText()) > 0 {
			return result.GetText()
		}
	}

	return ""
}

func (uc *ImageUseCase) ToBlackAndWhiteImage(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			colorRgb := m.At(x, y)
			newRgba.Set(x, y, monoModel(colorRgb))
		}
	}

	return newRgba
}

func (uc *ImageUseCase) ReverseImageColor(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			colorRgb := m.At(x, y)
			_, g, _, a := colorRgb.RGBA()
			newG := 255 - uint8(g>>8)
			// 将每个点的设置为灰度值
			newRgba.SetRGBA(x, y, color.RGBA{
				R: newG,
				G: newG,
				B: newG,
				A: uint8(a >> 8),
			})
		}
	}

	return newRgba
}

type Pixel bool

const (
	Black Pixel = true
	White Pixel = false
)

// RGBA returns the RGBA values for the receiver.
func (c Pixel) RGBA() (r, g, b, a uint32) {
	if c == Black {
		return 0, 0, 0, 0xffff
	}
	return 0xffff, 0xffff, 0xffff, 0xffff
}

func monoModel(c color.Color) color.Color {
	if _, ok := c.(Pixel); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := (299*r + 587*g + 114*b + 500) / 500
	return Pixel(uint16(y) < 0x8000)
}
