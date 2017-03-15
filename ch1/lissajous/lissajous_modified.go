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
	"os"
	"strconv"
	"time"
)

// Packages not needed by version in book.

var palette = []color.Color{
	color.RGBA{0xFF, 0, 0, 0xFF},       // red
	color.RGBA{0xFF, 0xA5, 0, 0xFF},    // orange
	color.RGBA{0xFF, 0xFF, 0, 0xFF},    // yellow
	color.RGBA{0, 0xFF, 0, 0xFF},       // green
	color.RGBA{0, 0, 0xFF, 0xFF},       // blue
	color.RGBA{0x4B, 0, 0x82, 0xFF},    // indigo
	color.RGBA{0xEE, 0x82, 0xEE, 0xFF}, // violet
}

/*
 * Exercise 1.12: Modify the Lissajous server to read parameter values from the URL.
 * For example, you might arrange it so that a URL like http://localhost:8000/?cycles=20
 * sets the number of cycles to 20 instead of the default 5.
 * Use the strconv.Atoi function to convert the string parameter into an integer.
 * You can see its documentation with go doc strconv.Atoi.
 */
func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			params, err := strconv.Atoi(r.FormValue("cycles"))
			if err == nil {
				lissajousModified(w, params)
			}
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajousModified(os.Stdout, 0)
}

func lissajousModified(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		index := uint8(rand.Intn(len(palette)))
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
