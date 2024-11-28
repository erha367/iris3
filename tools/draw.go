package tools

import (
	"github.com/chai2010/webp"
	"image/png"

	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func DrawLines(Path string) {
	sep := `/`
	var (
		img image.Image
		err error
	)
	fileArr := strings.Split(Path, sep)
	fileName := fileArr[len(fileArr)-1]
	// 在图片上绘制一条x色线段
	// 打开图片文件
	file, err := os.Open(Path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if strings.HasSuffix(Path, ".webp") {
		img, err = webp.Decode(file)
	} else if strings.HasSuffix(Path, ".png") {
		img, err = png.Decode(file)
		log.Println(`o png err`, err)
	} else {
		// 解码 JPEG 图片
		img, err = jpeg.Decode(file)
	}
	if err != nil {
		panic(err)
	}
	//计算宽高
	bounds := img.Bounds()
	x := bounds.Dx()
	y := bounds.Dy()
	log.Println(`图片宽高为(px)：`, x, y)
	line := image.NewRGBA(img.Bounds())
	toLineXY := GetXY(x, y)
	for _, v := range toLineXY {
		drawLine(line, v.Color, v.Start, v.End)
	}
	log.Println("drawLine finished.")
	// 将线段合并到原始图片上
	result := image.NewRGBA(img.Bounds())
	draw.Draw(result, img.Bounds(), img, image.ZP, draw.Src)
	draw.Draw(result, line.Bounds(), line, image.ZP, draw.Over)
	// 保存结果图片(时间戳)
	now := time.Now().Unix()
	nowStr := strconv.Itoa(int(now))
	fName := nowStr + `_` + fileName
	outFile, err := os.Create(fName)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	jpeg.Encode(outFile, result, &jpeg.Options{Quality: 100})
	log.Println("success !", fName)
}

// 绘制一条线段
func drawLine(img *image.RGBA, c color.RGBA, start, end image.Point) {
	dx := abs(end.X - start.X)
	dy := abs(end.Y - start.Y)
	sx := -1
	if start.X < end.X {
		sx = 1
	}
	sy := -1
	if start.Y < end.Y {
		sy = 1
	}
	err := dx - dy

	for {
		img.Set(start.X, start.Y, c)
		if start == end {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			start.X += sx
		}
		if e2 < dx {
			err += dx
			start.Y += sy
		}
	}
}

// 返回 x 的绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
