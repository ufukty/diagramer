package measure

import "golang.org/x/image/math/fixed"

type size struct {
	Width, Height fixed.Int26_6
}

type _range struct {
	Min, Max size
}

type dimensions struct {
	Content, Asked _range
}

func clamp(a, b fixed.Int26_6) fixed.Int26_6 {
	return min(a, b)
}

func (d *dimensions) Negotiate() size {
	return size{
		Width: clamp(
			max(d.Content.Min.Width, d.Asked.Min.Width),
			min(d.Content.Max.Width, d.Asked.Max.Width),
		),
		Height: clamp(
			max(d.Content.Min.Height, d.Asked.Min.Height),
			min(d.Content.Max.Height, d.Asked.Max.Height),
		),
	}
}
