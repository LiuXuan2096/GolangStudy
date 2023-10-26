package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	const size = 300                                   // 图片大小
	pic := image.NewGray(image.Rect(0, 0, size, size)) // 根据大小创建灰度图
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// 把图片中每个像素填充为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size  // 让sin的值的范围在0-2π之间
		y := size/2 - math.Sin(s)*(size/2)    // sin的幅度为一半的像素。向下偏移一半像素并翻转
		pic.SetGray(x, int(y), color.Gray{0}) // 用黑色绘制sin轨迹
	}

	file, err := os.Create("D:\\GoProject\\GolangStudy\\6_流程控制\\sin.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, pic)
	file.Close()
}
