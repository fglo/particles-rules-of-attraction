package gui

import (
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func loadFont(path string, size float64) (font.Face, error) {
	fontData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}
