package biz

import (
	"bytes"
	"context"
	v1 "github.com/SuKaiFei/go-wxxcx/api/wxxcx/v1"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/SuKaiFei/go-wxxcx/util/wordcloud"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nfnt/resize"
	errors2 "github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"math/rand"
	"time"
)

type WordcloudUseCase struct {
	log      *log.Helper
	imageUC  *ImageUseCase
	wechatUC *WechatUseCase

	minFontSize, maxFontSize float64
	angles                   []int
	colors                   []*color.RGBA
	fontPath                 string
}

func NewWordcloudUseCase(logger log.Logger, imageUC *ImageUseCase, wechatUC *WechatUseCase, appConf *conf.Application) *WordcloudUseCase {
	angles := []int{0, 15, -15, 90}
	colors := []*color.RGBA{
		{0x0, 0x60, 0x30, 0xff},
		{0x60, 0x0, 0x0, 0xff},
		{255, 255, 255, 1},
	}
	return &WordcloudUseCase{
		maxFontSize: 60,
		minFontSize: 8,
		imageUC:     imageUC,
		wechatUC:    wechatUC,
		log:         log.NewHelper(logger),
		angles:      angles,
		colors:      colors,
		fontPath:    appConf.WordcloudFontPath,
	}
}

func (uc *WordcloudUseCase) GenerateImage(ctx context.Context, imagePath string, words []string) (*v1.UploadImageReply, error) {
	if len(words) == 0 {
		return nil, errors.New(400, "", "请输入词语")
	}

	words = uc.randomStrings(words)

	_, body, err := fasthttp.GetTimeout(nil, imagePath, 5*time.Second)
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	inImg, imgType, err := image.Decode(bytes.NewReader(body))
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	outImageBytes := new(bytes.Buffer)
	if imgType == "gif" {
		gifObj, err := gif.DecodeAll(bytes.NewReader(body))
		if err != nil {
			return nil, errors2.WithStack(err)
		}

		outGif := new(gif.GIF)
		for i, gitImage := range gifObj.Image {
			tmp, err := uc.generate(words, gitImage)
			if err != nil {
				return nil, errors2.WithStack(err)
			}
			newGifImage := image.NewPaletted(tmp.Bounds(), gitImage.Palette)
			draw.Draw(newGifImage, tmp.Bounds(), tmp, tmp.Bounds().Size(), draw.Over)
			if outGif.Config.Width < tmp.Bounds().Size().X {
				outGif.Config.Width = tmp.Bounds().Size().X
			}
			if outGif.Config.Height < tmp.Bounds().Size().Y {
				outGif.Config.Height = tmp.Bounds().Size().Y
			}

			outGif.Image = append(outGif.Image, newGifImage)
			outGif.Delay = append(outGif.Delay, gifObj.Delay[i])
		}
		if err := gif.EncodeAll(outImageBytes, outGif); err != nil {
			return nil, errors2.WithStack(err)
		}
	} else {
		outImg, err := uc.generate(words, inImg)
		if err != nil {
			return nil, errors2.WithStack(err)
		}
		if imgType == "png" {
			if err := png.Encode(outImageBytes, outImg); err != nil {
				return nil, errors2.WithStack(err)
			}
		} else {
			if err := jpeg.Encode(outImageBytes, outImg, nil); err != nil {
				return nil, errors2.WithStack(err)
			}
		}
	}

	req := &v1.UploadImageRequest{
		File: outImageBytes.Bytes(),
		Code: "wordcloud",
	}
	return uc.imageUC.UploadImage(ctx, req)
}

func (uc *WordcloudUseCase) generate(words []string, inImg image.Image) (image.Image, error) {
	startedAt := time.Now()
	defer func() {
		uc.log.Infof("generate消耗:%s", time.Now().Sub(startedAt).String())
	}()
	var resizeImg image.Image
	point := inImg.Bounds().Size()
	var (
		width   uint = 0
		height  uint = 0
		minSize uint = 400
	)

	if point.X < int(minSize) {
		width = minSize
	}
	if point.Y < int(minSize) {
		height = minSize
	}
	if width == minSize && height == minSize {
		if point.X > point.Y {
			width = 0
		} else {
			height = 0
		}
	}

	if width == 0 && height == 0 {
		resizeImg = inImg
	} else {
		resizeImg = resize.Resize(width, height, inImg, resize.NearestNeighbor)
	}
	render, err := wordcloud.NewWordCloudRender(uc.maxFontSize, uc.minFontSize, uc.fontPath, resizeImg, words, uc.angles, uc.colors)
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	return render.Render(), nil
}

func (uc *WordcloudUseCase) randomStrings(words []string) []string {
	for i := len(words) - 1; i > 0; i-- {
		rNum := rand.Intn(i + 1)
		words[i], words[rNum] = words[rNum], words[i]
	}
	return words
}
