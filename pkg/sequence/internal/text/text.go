package text

import (
	"golang.org/x/image/font"
)

func TextWidth(text string, face font.Face) int {
	return (&font.Drawer{Face: face}).MeasureString(text).Round()
}
