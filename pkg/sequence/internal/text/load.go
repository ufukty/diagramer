package text

import (
	"fmt"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const DPI = 72

func LoadFont(path string, size float64) (font.Face, error) {
	fontBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading font file: %w", err)
	}

	fnt, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("parsing font: %w", err)
	}

	face, err := opentype.NewFace(fnt, &opentype.FaceOptions{
		Size:    size,
		DPI:     DPI,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, fmt.Errorf("creating font face: %w", err)
	}

	return face, nil
}
