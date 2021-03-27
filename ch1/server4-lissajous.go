// Server2 is a minimal "echo" and counter server
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int
var palette = []color.Color{
	color.Black,
	color.RGBA{0x17, 0x3F, 0x5F, 0xff},
	color.RGBA{0x20, 0x63, 0x9B, 0xff},
	color.RGBA{0x3C, 0xAE, 0xA3, 0xff},
	color.RGBA{0xF6, 0xD5, 0x5C, 0xff},
	color.RGBA{0xED, 0x55, 0x3B, 0xff},
}

const (
	blackIndex    = 0 // first color in palette
	darkBlueIndex = 1 // second color in palette
	blueIndex     = 2
	tealIndex     = 3
	yellowIndex   = 4
	orangeIndex   = 5
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles, err := strconv.Atoi(r.URL.Query().Get("cycles"))
		if err != nil {
			log.Fatal(err)
		}
		lissajous(w, cycles)
	})
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer, cycles int) {
	const (
		// cycles  = 5       // number of complete x oscillator revolutions
		res     = 0.00001 // angular resolution
		size    = 100     // image canvas covers [-size..+size]
		nframes = 64      // number of animation frames
		delay   = 8       // delay between frames in 10ms units
	)

	if cycles == 0 {
		cycles = 5
	}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		var colorIndex = uint8(1)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			if i%3 == 0 {
				colorIndex = uint8(rand.Intn(5-1+1) + 1)
			}
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		return
	}
}
