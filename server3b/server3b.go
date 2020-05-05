// server 3
// If you want to run it:
//   sudo lsof -i:8000 ##make sure nobody else is listening on port 8000
//   go run server3.go &
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"	
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
//	"os"
)



var palette = []color.Color{color.White, color.Black}
var mu sync.Mutex
var count int
var f = fmt.Fprintf

func main() {
	//Run lissajous and turn off the echo.
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    //	lissajous(w)
	//})

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the http request
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	
	u, err := url.Parse(*r.URL)
	if err != nil {
		panic(err)
	}
	f(w, "URL is %s", u.URL)
	f(w, u.RawQuery)
	f(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		f(w, "Header[%q] = %q\n", k, v)
	}
	f(w, "Host = %q\n", r.Host)
	f(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		f(w, "Form[%q] = %q\n", k, v)
	}
	
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	f(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer) {
	const (
		cycles	= 5 		// number of complete x oscillator revolutions
		res		= 0.001		// angular resolution
		size	= 100		// image canvas covers [-size..+size]
		nframes = 64		// number of animation frames
		delay   = 8			// delay between frames in 10 ms units
	)

	mu.Lock()
	count++
	mu.Unlock()
	
	freq	:= rand.Float64() * 3.0 	// relative frequency of y oscillator
	anim 	:= gif.GIF{LoopCount: nframes}
	phase 	:= 0.0
	for i:= 0; i < nframes; i++ {
		rect	:= image.Rect(0, 0, 2*size+1, 2*size+1)
		img 	:= image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase		+=	0.1
		anim.Delay	=	append(anim.Delay, delay)
		anim.Image  =	append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Note: ignoring encoding errors
}