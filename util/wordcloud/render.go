package wordcloud

import (
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"strconv"
)

type WordCloudRender struct {
	MaxFontSize float64
	MinFontSize float64
	FontPath    string
	OutlineImg  image.Image
	MeasureDc   *gg.Context
	DrawDc      *gg.Context
	TextList    []string
	Angles      []int
	Colors      []*color.RGBA
	OutImgPath  image.Image
	worldMap    *WorldMap
}

func NewWordCloudRender(maxFontSize, minFontSize float64, fontPath string,
	img image.Image, textList []string,
	angles []int, colors []*color.RGBA) (*WordCloudRender, error) {

	render := &WordCloudRender{
		MaxFontSize: maxFontSize,
		MinFontSize: minFontSize,
		FontPath:    fontPath,
		OutlineImg:  img,
		TextList:    textList,
		Angles:      angles,
		Colors:      colors,
	}
	worldMap := TwoByBitmap(img)
	render.worldMap = worldMap
	drawDc := gg.NewContext(worldMap.RealImageWidth, worldMap.RealImageHeight)
	drawDc.SetRGB(1, 1, 1)
	drawDc.Clear()
	drawDc.SetRGB(0, 0, 0)
	render.DrawDc = drawDc
	if err := drawDc.LoadFontFace(fontPath, render.MaxFontSize); err != nil {
		return nil, err
	}

	render.ResetMeasureDc(render.MaxFontSize)
	return render, nil
}

func (wcr *WordCloudRender) Render() image.Image {
	fontSize := wcr.MaxFontSize
	currentTextIdx := 0
	colorIdx := 0
	checkRet := &CheckResult{}

	gridCache := make(map[string]*Grid)

	var itemGrid *Grid
	bigestSizeCnt := 0
	for {
		msg := wcr.TextList[currentTextIdx]
		key := strconv.Itoa(int(fontSize)) + msg
		if _, ok := gridCache[key]; ok {
			itemGrid = gridCache[key]
		} else {
			itemGrid = &Grid{}
			w, h, xscale, yscale := GetTextBound(wcr, fontSize, msg)
			itemGrid.XScale = int(xscale)
			itemGrid.YScale = int(yscale)
			if int(w)%2 != 0 {
				w += XUNIT
			}
			if int(h)%2 != 0 {
				h += YUNIT
			}
			positions, w1, h1 := TwoByBlock(int(w), int(h))
			itemGrid.Width = int(w1)
			itemGrid.Height = int(h1)
			itemGrid.positions = positions
			gridCache[key] = itemGrid
		}

		isFound := wcr.collisionCheck(
			0, wcr.worldMap, itemGrid, checkRet, wcr.Angles)
		if isFound {
			currentTextIdx++
			currentTextIdx = currentTextIdx % len(wcr.TextList)
			color := wcr.Colors[colorIdx]
			colorIdx++
			colorIdx = colorIdx % len(wcr.Colors)
			wcr.DrawDc.SetRGB(float64(color.R), float64(color.G), float64(color.B))
			DrawText(wcr.DrawDc, msg, float64(checkRet.Xpos+itemGrid.XScale/2),
				float64(checkRet.Ypos+itemGrid.YScale/2), Angle2Pi(float64(checkRet.Angle)))
			if fontSize == wcr.MaxFontSize {
				bigestSizeCnt++
				if bigestSizeCnt > len(wcr.TextList) {
					fontSize = 40
					wcr.UpdateFontSize(fontSize)
				}
			}
		} else {
			fontSize -= 3
			if fontSize < wcr.MinFontSize {
				break
			}
			wcr.UpdateFontSize(fontSize)
		}
	}
	return wcr.DrawDc.Image()
}

func (wcr *WordCloudRender) UpdateFontSize(fontSize float64) {
	face, err := gg.LoadFontFace(wcr.FontPath, fontSize)
	if err != nil {
		panic(err)
	}
	wcr.DrawDc.SetFontFace(face)
	wcr.MeasureDc.SetFontFace(face)
}

func (wcr *WordCloudRender) ResetMeasureDc(fontSize float64) {
	measureDc := gg.NewContext(wcr.worldMap.RealImageWidth, wcr.worldMap.RealImageHeight)
	measureDc.SetRGBA(0, 0, 0, 0)
	measureDc.Clear()
	wcr.MeasureDc = measureDc
	if err := measureDc.LoadFontFace(wcr.FontPath, fontSize); err != nil {
		panic(err)
	}
}

func (wcr *WordCloudRender) collisionCheck(lastCheckAngle float64, worldMap *WorldMap,
	itemGrid *Grid, ret *CheckResult, tryAngles []int) bool {

	centerX := worldMap.Width / 2
	centerY := worldMap.Height / 2
	isFound := true
	xDistanceToCenter := 0
	yDistanceToCenter := 0
	tempXpos := 0
	tempYpos := 0

	angleMark := 0
	currentAngleIdx := 0
	for angle := lastCheckAngle; angle <= DEGREE_360; angle += 1 {
		currentAngleIdx = 0
		angleMark = tryAngles[currentAngleIdx]
		currentAngleIdx++
		Rotate(itemGrid, float64(angleMark), centerX, centerY)
		xDiff := CosT(angle) * 1
		yDiff := SinT(angle) * 1
		tempXpos = 0
		tempYpos = 0
		xLeiji := xDiff
		yLeiji := yDiff
		xDistanceToCenter = 0
		yDistanceToCenter = 0
		result := IS_NOT_FIT
		for {
			result = IS_NOT_FIT
			if xDistanceToCenter != tempXpos || yDistanceToCenter != tempYpos {
				tempXpos = xDistanceToCenter
				tempYpos = yDistanceToCenter
				result = itemGrid.IsFit(xDistanceToCenter, yDistanceToCenter, worldMap.Width, worldMap.Height, worldMap.CollisionMap)
				if result == OUT_INDEX {
					if currentAngleIdx < len(tryAngles) {
						angleMark = tryAngles[currentAngleIdx]
						currentAngleIdx++
						Rotate(itemGrid, float64(angleMark), centerX, centerY)
						xLeiji = xDiff
						yLeiji = yDiff
						tempXpos = 0
						tempYpos = 0
						xDistanceToCenter = 0
						yDistanceToCenter = 0
					} else {
						ret.Angle = 0
						isFound = false
						break
					}
				} else if result == IS_FIT {
					isFound = true
					itemGrid.Fill(worldMap.Width, worldMap.Height, worldMap.CollisionMap)
					ret.Angle = angleMark
					ret.Xpos = (xDistanceToCenter + centerX) * XUNIT
					ret.Ypos = (yDistanceToCenter + centerY) * YUNIT
					ret.LastCheckAngle = int(angle)
					break
				}
			}
			xLeiji += xDiff
			yLeiji += yDiff
			xDistanceToCenter = int(CeilT(xLeiji))
			yDistanceToCenter = int(CeilT(yLeiji))
		}
		if angle >= DEGREE_360 {
			ret.Angle = 0
			isFound = false
			break
		}
		if result == IS_FIT {
			break
		}
	}
	return isFound
}
