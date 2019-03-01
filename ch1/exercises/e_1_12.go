package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
)

/**
	运行：go build ch1/lissajous.go
 		 ./lissajous > img.gif
 */

// color.Color{...} 复合声明，生成的是一个slice切片
var pallete = []color.Color{
	color.White,
	color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff}, 	// red
	color.RGBA{0x00, 0xff, 0x00, 0xff}, 	// green
	color.RGBA{0x00, 0x00, 0xff, 0xff}, 	// blue
	}


// 常量声明和变量声明一般出现在包级别，常量在整个包中都是可以共享的
const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2
)

func main() {
	// 生成可执行文件
	// lissajois(os.Stdout)

	// 浏览器展示gif动画
	http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request) {
		if err := r.ParseForm(); err != nil {
			cyclesRes := r.Form.Get("cycles")
			cycles := 0
			if len(cyclesRes) == 0 {
				cycles = 5
			}
			cycles = int(cyclesRes[0])
			lissajois(w, cycles)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajois(out io.Writer, cycles int) {
	// 常量定义在函数体内部，则只能在函数体内使用
	const (
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)

	freq := rand.Float64() * 3.0

	// gif.GIF{...} 复合声明，生成的是一个struct结构体
	// 内部变量LoopCount被设置为nframes，而其他字段被设置成为各自类型默认的零值
	anim := gif.GIF{LoopCount:nframes}
	phase := 0.0

	// 生成64帧
	for i:=0; i<nframes; i++ {
		// 生成新图片：201*201矩形，黑白两种颜色
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Cos(t*freq + phase)
			idx := rand.Intn(5)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(idx)+1)
		}
		phase += 0.1
		// 80ms延迟
		anim.Delay = append(anim.Delay, delay)
		// 向gif图片末尾追加一帧
		anim.Image = append(anim.Image, img)
	}
	// 编码成gif图片，并将结果写入标准输出流
	gif.EncodeAll(out, &anim)
}