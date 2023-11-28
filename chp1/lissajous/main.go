package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var pallete []color.Color = []color.Color{color.White, color.RGBA{8, 76, 97, 100}, color.RGBA{219, 80, 74, 100}, color.RGBA{227, 181, 5, 100}}

const (
	whiteIndex = 0 // first color in pallete
	blackIndex = 1 // second color in pallete
)

func main() {
	lissajous(os.Stdout)

}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	// freq := 2.0099 love this one
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(3)+1))
		}
		phase += 1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	if err := gif.EncodeAll(out, &anim); err != nil {
		fmt.Printf("[ERROR]%v", err)
	}
}
