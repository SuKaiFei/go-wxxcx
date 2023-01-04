package service

import (
	"github.com/nfnt/resize"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"testing"
)

func TestImageResize(t *testing.T) {
	file, err := os.Open("/Users/sukaifei/Downloads/6a2dc9db-50e5-469e-a0b7-e1643aeb1d32.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

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
			maxHeight = 1080
		}
	}

	imgResize := resize.Resize(maxWidth, maxHeight, img, resize.NearestNeighbor)

	if err := png.Encode(fNearestNeighbor, imgResize); err != nil {
		t.Fatal(err)
	}
}
