package text

import (
	"fmt"
	"os"

	"github.com/flopp/go-findfont"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

const DPI = 72

type Font struct {
	face font.Face
}

func LoadFont(fontname string, size float64) (*Font, error) {
	path, err := findfont.Find(fontname)
	if err != nil {
		return nil, fmt.Errorf("finding font file: %w", err)
	}

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

	return &Font{face}, nil
}

func (f Font) Close() error {
	return f.face.Close()
}

func (f *Font) TextWidth(text string) fixed.Int26_6 {
	return (&font.Drawer{Face: f.face}).MeasureString(text)
}
