package text

import (
	"maps"
	"slices"
	"testing"

	"golang.org/x/image/math/fixed"
)

var FONTPATH = "/System/Library/Fonts/Supplemental/Arial.ttf" // macOS

func TestFont_TextWidth_HardCoded(t *testing.T) {
	font, err := LoadFont(FONTPATH, 14)
	if err != nil {
		t.Fatalf("prep: %v", err)
	}
	defer font.Close()

	tcs := map[string]fixed.Int26_6{
		"Hello Arial 42": fixed.I(82),
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			got := font.TextWidth(input)
			expected := tcs[input]
			if got != expected {
				t.Errorf("expected %v got %v", expected, got)
			}
		})
	}
}

func TestFont_TextWidth_Relative(t *testing.T) {
	font, err := LoadFont(FONTPATH, 14)
	if err != nil {
		t.Fatalf("prep: %v", err)
	}
	defer font.Close()

	tcs := []string{
		"Hello Arial",
		`<!DOCTYPEhtml><htmllang="en"><head><metacharset="UTF-8"><metahttp-equiv="X-UA-Compatible"content="IE=edge">`,
		"Lorem ipsum dolor sit amet. Consectur adipiscing elit.",
	}

	for _, input := range tcs {
		t.Run(input, func(t *testing.T) {
			halfinput := input[:len(input)/2]
			width, halfwidth := font.TextWidth(input), font.TextWidth(halfinput)
			if width <= halfwidth {
				t.Errorf("expected %q (%dpx) to be shorter than %q (%dpx)", halfinput, halfwidth, input, width)
			}
		})
	}
}
