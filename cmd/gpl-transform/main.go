package main

import (
	"bytes"
	"flag"
	"fmt"
	"image/color"
	"io/ioutil"
	"os"

	"github.com/gravestench/gpl/pkg"
)

type options struct {
	srcGpl  *string
	dstGpl  *string
	outPath *string
}

func parseOptions(o *options) (terminate bool) {
	o.srcGpl = flag.String("src", "", "input pal file (required)")
	o.dstGpl = flag.String("dst", "", "input pal file (required)")
	o.outPath = flag.String("out", "", "path to output gpl file (optional)")

	flag.Parse()

	return *o.srcGpl == ""
}

func main() {
	o := &options{}

	if parseOptions(o) {
		flag.Usage()
	}

	src := decodeGpl(*o.srcGpl)
	dst := decodeGpl(*o.dstGpl)

	newPalette := make(color.Palette, len(src))

	for srcIdx := range src {
		newPalette[srcIdx] = dst.Convert(src[srcIdx])
	}

	if *o.outPath != "" {
		writePalette(*o.outPath, newPalette)
	}
}

func decodeGpl(path string) color.Palette {
	srcData, err := ioutil.ReadFile(path)

	if err != nil {
		const fmtErr = "could not read file, %v"
		fmt.Print(fmt.Errorf(fmtErr, err))

		return nil
	}

	pal, err := pkg.Decode(bytes.NewBuffer(srcData))
	if err != nil {
		return nil
	}

	return color.Palette(pal)
}

func writePalette(outPath string, pal color.Palette) {
	f, err := os.Create(outPath)
	if err != nil {
		fmt.Println(err)
	}

	if err := pkg.FromPalette(pal).Encode("", f); err != nil {
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
