package text

import (
	"fmt"
	"testing"
)

func TestWrite(t *testing.T) {
	fontPath := "/System/Library/Fonts/Supplemental/Arial.ttf" // macOS example

	face, err := LoadFont(fontPath, 14)
	if err != nil {
		t.Fatalf("prep: %v", err)
	}
	defer face.Close()

	text := "Hello Arial 42"
	width := TextWidth(text, face)

	fmt.Printf("Measured width for '%s' in Arial: %d px\n", text, width)
}
