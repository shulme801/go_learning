// Lissajous generates GIF animations of random Lissajous figures.package lissajous

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//var palette = []color.Color{color.White, color.Black, color.RGBA{0x7C, 0xFC, 0x00, 0xff}}
var palette = []color.Color{color.Black, color.RGBA{0x7C, 0xFC, 0x00, 0xff}}
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	greenIndex = 2 // our green color
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles	= 5 		// number of complete x oscillator revolutions
		res		= 0.001		// angular resolution
		size	= 100		// image canvas covers [-size..+size]
		nframes = 64		// number of animation frames
		delay   = 8			// delay between frames in 10 ms units
	)
	freq	:= rand.Float64() * 3.0 	// relative frequency of y oscillator
	anim 	:= gif.GIF{LoopCount: nframes}
	phase 	:= 0.0
	for i:= 0; i < nframes; i++ {
		rect	:= image.Rect(0, 0, 2*size+1, 2*size+1)
		myPalette := palette[blackIndex:greenIndex]
		img 	:= image.NewPaletted(rect, myPalette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase		+=	0.1
		anim.Delay	=	append(anim.Delay, delay)
		anim.Image  =	append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Note: ignoring encoding errors
}