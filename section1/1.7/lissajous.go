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
	"strconv"
	"time"
)

type lissaconf struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

var palette = []color.Color{
	color.Black,
	color.RGBA{
		R: 0,
		G: 0xFF,
		B: 0,
		A: 0xFF,
	}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UnixNano())
	lconf := lissaconf{
		cycles:  5,
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		for i, c := range params {
			switch i {
			case "cycles":
				lconf.cycles, _ = strconv.Atoi(c[0])
			case "res":
				lconf.res, _ = strconv.ParseFloat(c[0], 64)
			case "size":
				lconf.size, _ = strconv.Atoi(c[0])
			case "nframes":
				lconf.nframes, _ = strconv.Atoi(c[0])
			case "delay":
				lconf.delay, _ = strconv.Atoi(c[0])
			}
		}
		lissajous(w, lconf)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, l lissaconf) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: l.nframes}
	phase := 0.0
	for i := 0; i < l.nframes; i++ {
		rect := image.Rect(0, 0, 2*l.size+1, 2*l.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(l.cycles)*2*math.Pi; t += l.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(l.size+int(x*float64(l.size)+0.5), l.size+int(y*float64(l.size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, l.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
