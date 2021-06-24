package pkg

import (
	"bufio"
	"bytes"
	"fmt"
	"image/color"
	"io"
	"math"
	"strconv"
	"strings"
)

type GimpPalette color.Palette

type GPL = GimpPalette

const (
	commentCharacter = "#"
	line1            = "GIMP Palette\r\n"
	line2            = "# Name: %s\r\n"
	line3            = "#\r\n"
	fmtComponent     = "  %v"
	fmtLine          = "%s %s %s\r\n"
	fmtErr           = "error encoding DAT to p format, %v"
)

/*
GIMP Palette
# Name: Bears
#
8   8   8	grey3
68  44  44
80   8  12
72  56  56
104  84  68
116  96  80
84  56  44
140 104  88
*/

func Decode(r io.Reader) (GPL, error) {
	lineScan := bufio.NewScanner(r)

	lines := make([]string, 0)
	for lineScan.Scan() {
		lines = append(lines, lineScan.Text())
	}

	gpl := make(GPL, 0)

	for lineIdx := range lines {
		wordScan := bufio.NewScanner(bytes.NewBufferString(lines[lineIdx]))
		wordScan.Split(bufio.ScanWords)

		words := make([]string, 0)
		for wordScan.Scan() {
			words = append(words, wordScan.Text())
		}

		if len(words) < 1 {
			continue
		}

		if strings.Contains(words[0], commentCharacter) {
			continue
		}

		if _, err := strconv.ParseInt(words[0], 10, 32); err != nil {
			continue
		}

		if len(words) < 3 {
			continue
		}

		r, err := strconv.ParseInt(words[0], 10, 32)
		if err != nil {
			return nil, err
		}

		g, err := strconv.ParseInt(words[1], 10, 32)
		if err != nil {
			return nil, err
		}

		b, err := strconv.ParseInt(words[2], 10, 32)
		if err != nil {
			return nil, err
		}

		c := color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: math.MaxUint8,
		}

		gpl = append(gpl, c)
	}

	return gpl, nil
}

func (p GPL) Encode(name string, w io.Writer) error {
	const numHeaderLines = 3

	numColors := len(p)
	numLines := numColors + numHeaderLines
	lines := make([]string, numLines)

	lines[0] = line1
	lines[1] = fmt.Sprintf(line2, name)
	lines[2] = line3

	strComponent := func(n int) string {
		s := fmt.Sprintf(fmtComponent, n)
		return s[len(s)-3:]
	}

	for idx := range p {
		r, g, b, _ := p[idx].RGBA()
		rs := strComponent(int(uint8(r)))
		gs := strComponent(int(uint8(g)))
		bs := strComponent(int(uint8(b)))

		lines[numHeaderLines+idx] = fmt.Sprintf(fmtLine, rs, gs, bs)
	}

	for idx := range lines {
		if _, err := w.Write([]byte(lines[idx])); err != nil {
			return fmt.Errorf(fmtErr, err)
		}
	}

	return nil
}

func FromPalette(p color.Palette) GPL {
	return GPL(p)
}
