package gpl

import (
	"image/color"
	"io"

	"github.com/gravestench/gpl/pkg"
)

// these aliases are here so you can import from repo root

type (
	GPL         = pkg.GPL
	GimpPalette = pkg.GimpPalette
)

func Decode(r io.Reader) (GPL, error) {
	return pkg.Decode(r)
}

func FromPalette(p color.Palette) GPL {
	return pkg.FromPalette(p)
}
