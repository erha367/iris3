package tools

import (
	"image"
	"image/color"
)

type Lines struct {
	Start, End image.Point
	Color      color.RGBA
}

var (
	RedLine  = color.RGBA{R: 255, A: 255}
	BlueLine = color.RGBA{R: 127, G: 255, B: 212, A: 255}
)

// GetXY 获取要划线的坐标点集合
func GetXY(x, y int) []Lines {
	var (
		res   []Lines
		line  Lines
		scale int
	)
	if x > y {
		scale = y / 8
	} else {
		scale = x / 8
	}
	//处理横线
	for i := 0; scale*i <= y; i++ {
		if i%2 == 0 {
			line.Color = RedLine
		} else {
			line.Color = BlueLine
		}
		line.Start = image.Point{X: 0, Y: scale * i}
		line.End = image.Point{X: x, Y: scale * i}
		res = append(res, line)
	}
	//处理竖线
	for i := 0; scale*i <= x; i++ {
		if i%2 == 0 {
			line.Color = RedLine
		} else {
			line.Color = BlueLine
		}
		line.Start = image.Point{X: scale * i, Y: 0}
		line.End = image.Point{X: scale * i, Y: y}
		res = append(res, line)
	}
	return res
}
