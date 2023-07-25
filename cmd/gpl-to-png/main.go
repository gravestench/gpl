package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/gravestench/gpl/pkg"
)

type options struct {
	gplPath *string
	pngPath *string
}

func parseOptions(o *options) (terminate bool) {
	o.gplPath = flag.String("pal", "", "input pal file (required)")
	o.pngPath = flag.String("png", "", "path to png file (optional)")

	flag.Parse()

	return *o.gplPath == ""
}

func main() {
	o := &options{}

	if parseOptions(o) {
		flag.Usage()
	}

	data, err := ioutil.ReadFile(*o.gplPath)
	if err != nil {
		const fmtErr = "could not read file, %v"
		fmt.Print(fmt.Errorf(fmtErr, err))

		return
	}

	pal, err := pkg.Decode(bytes.NewBuffer(data))
	if err != nil {
		return
	}

	img := makeImage(color.Palette(pal))

	if *o.pngPath != "" {
		writeImage(*o.pngPath, img)
	}
}

func makeImage(pal color.Palette) image.Image {
	numPixels := len(pal)

	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: numPixels, Y: 1},
	})

	for palIdx := range pal {
		r, g, b, a := pal[palIdx].RGBA()

		rgba := color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: uint8(a),
		}

		img.SetRGBA(palIdx, 0, rgba)
	}

	return img
}

func writeImage(outPath string, img image.Image) {
	f, err := os.Create(outPath)
	if err != nil {
		fmt.Println(err)
	}

	if err := png.Encode(f, img); err != nil {
		fmt.Println(err)

		err = f.Close()
		if err != nil {
			fmt.Println(err)
		}

		return
	}

	if err = f.Close(); err != nil {
		fmt.Println(err)
	}
}
