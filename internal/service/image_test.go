package service

import (
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"testing"
	"time"
)

func TestImageResize(t *testing.T) {
	file, err := os.Open("/Users/sukaifei/Downloads/IMG_3851.PNG")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	now := time.Now()
	img, s, err := image.Decode(file)
	if err != nil {
		t.Fatal(err)
	}
	//s := "a.png"
	t.Log("文件名：", s)

	fNearestNeighbor, err := os.Create("/Users/sukaifei/Downloads/imgNearestNeighbor." + s)
	if err != nil {
		t.Fatal(err)
	}
	defer fNearestNeighbor.Close()

	var (
		maxWidth, maxHeight, imgSize uint
		isWidth                      bool
	)

	if img.Bounds().Dx() > img.Bounds().Dy() {
		isWidth = true
		imgSize = uint(img.Bounds().Size().X)
	} else {
		imgSize = uint(img.Bounds().Size().Y)
	}
	if imgSize > 1080 {
		if isWidth {
			maxWidth = 1080
		} else {
			//maxHeight = 1080
		}
	}

	imgResize := resize.Resize(maxWidth, maxHeight, img, resize.Lanczos3)
	t.Log("时间消耗：", time.Now().Sub(now).String())

	if err := png.Encode(fNearestNeighbor, imgResize); err != nil {
		t.Fatal(err)
	}
	//Bilinear
	//Bicubic
	//MitchellNetravali
	//Lanczos2
	//Lanczos3
	tests := []string{
		"NearestNeighbor",
		"Bilinear",
		"Bicubic",
		"MitchellNetravali",
		"Lanczos2",
		"Lanczos3",
	}
	for i, test := range tests {
		imgResize := resize.Resize(maxWidth, maxHeight, img, resize.InterpolationFunction(i))

		fNearestNeighbor, err = os.Create("/Users/sukaifei/Downloads/image" + test + "." + s)
		if err := png.Encode(fNearestNeighbor, imgResize); err != nil {
			t.Fatal(err)
		}
	}
	t.Log("时间消耗：", time.Now().Sub(now).String())
}

func TestGet(t *testing.T) {
	file, err := os.Open("/Users/sukaifei/Downloads/WechatIMG311.jpeg")
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	content := new(biz.ImageUseCase).GetQRCodeContent(img)

	t.Log("content：", content)
	t.Log("时间消耗：", time.Now().Sub(now).String())
}

func TestNewQRCodeReader(t *testing.T) {
	file, err := os.Open("/Users/sukaifei/Downloads/WechatIMG311.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	rgba := grayingImage(img, 255)
	//if err := png.Encode(create, rgba); err != nil {
	//	t.Fatal(err)
	//}
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		t.Fatal(err)
	}
	reader := qrcode.NewQRCodeReader()
	result, err := reader.Decode(bmp, nil)
	t.Log("qrcode.NewQRCodeReader1", result, err)

	bmp, err = gozxing.NewBinaryBitmapFromImage(rgba)
	if err != nil {
		t.Fatal(err)
	}
	result, err = reader.Decode(bmp, nil)
	t.Log("qrcode.NewQRCodeReader2", result, err)
	create, err := os.Create("/Users/sukaifei/Downloads/123123.JPG")
	if err != nil {
		t.Fatal(err)
	}
	defer create.Close()
	if err := jpeg.Encode(create, rgba, nil); err != nil {
		t.Fatal(err)
	}
	rgba = grayingImage(img, 128)
	bmp, err = gozxing.NewBinaryBitmapFromImage(rgba)
	result, err = reader.Decode(bmp, nil)
	t.Log("qrcode.NewQRCodeReader3", result, err)
	t.Log("时间消耗：", time.Now().Sub(now).String())
	create, err = os.Create("/Users/sukaifei/Downloads/asdfsaf.JPG")
	if err != nil {
		t.Fatal(err)
	}
	defer create.Close()
	if err := jpeg.Encode(create, rgba, nil); err != nil {
		t.Fatal(err)
	}
}

func grayingImage(m image.Image, ng uint8) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			colorRgb := m.At(x, y)
			_, g, _, _ := colorRgb.RGBA()
			newG := uint8(g)
			if newG > 200 {
				newG = 1
			} else {
				newG = 0
			}
			//newA := uint8(a >> 8)
			// 将每个点的设置为灰度值
			newRgba.Set(x, y, monoModel(colorRgb))
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
