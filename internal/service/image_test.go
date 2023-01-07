package service

import (
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

func TestNewQRCodeReader(t *testing.T) {
	file, err := os.Open("/Users/sukaifei/Downloads/IMG_3924.JPG")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	rgba := grayingImage(img)
	//if err := png.Encode(create, rgba); err != nil {
	//	t.Fatal(err)
	//}
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		t.Fatal(err)
	}
	hints := map[gozxing.DecodeHintType]interface{}{
		gozxing.DecodeHintType_TRY_HARDER: true,
		gozxing.DecodeHintType_POSSIBLE_FORMATS: []gozxing.BarcodeFormat{
			gozxing.BarcodeFormat_QR_CODE},
	}
	reader := qrcode.NewQRCodeReader()
	result, err := reader.Decode(bmp, hints)
	t.Log("qrcode.NewQRCodeReader1", result, err)

	bmp, err = gozxing.NewBinaryBitmapFromImage(rgba)
	if err != nil {
		t.Fatal(err)
	}
	result, err = reader.Decode(bmp, hints)
	t.Log("qrcode.NewQRCodeReader2", result, err)
	t.Log("时间消耗：", time.Now().Sub(now).String())
	create, err := os.Create("/Users/sukaifei/Downloads/IMG_39241.JPG")
	if err != nil {
		t.Fatal(err)
	}
	defer create.Close()
	if err := jpeg.Encode(create, rgba, nil); err != nil {
		t.Fatal(err)
	}
}

func grayingImage(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			colorRgb := m.At(x, y)
			_, g, _, a := colorRgb.RGBA()
			newG := 255 - uint8(g>>8)
			newA := uint8(a >> 8)
			// 将每个点的设置为灰度值
			newRgba.SetRGBA(x, y, color.RGBA{R: newG, G: newG, B: newG, A: newA})
		}
	}

	return newRgba
}
