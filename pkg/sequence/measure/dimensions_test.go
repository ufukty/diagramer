package measure

import (
	"testing"

	"golang.org/x/image/math/fixed"
)

func wh(width, height int) size {
	return size{Width: fixed.I(width), Height: fixed.I(height)}
}

func r(min, max size) _range {
	return _range{Min: min, Max: max}
}

func TestDimensionsCases(t *testing.T) {
	cases := []struct {
		name  string
		input dimensional
		want  size
	}{
		{
			name:  "Content and Asked overlap nicely",
			input: dimensional{Content: r(wh(100, 50), wh(300, 200)), Asked: r(wh(150, 100), wh(250, 180))},
			want:  wh(150, 100),
		},
		{
			name:  "Asked min larger than Content max",
			input: dimensional{Content: r(wh(100, 100), wh(300, 300)), Asked: r(wh(400, 400), wh(500, 500))},
			want:  wh(400, 400), // clamps to asked min
		},
		{
			name:  "Content constraints dominate",
			input: dimensional{Content: r(wh(200, 200), wh(400, 400)), Asked: r(wh(100, 100), wh(500, 500))},
			want:  wh(200, 200),
		},
		{
			name:  "Content and Asked exact match",
			input: dimensional{Content: r(wh(150, 150), wh(150, 150)), Asked: r(wh(150, 150), wh(150, 150))},
			want:  wh(150, 150),
		},
		{
			name:  "Content wider than Asked max",
			input: dimensional{Content: r(wh(100, 100), wh(500, 500)), Asked: r(wh(50, 50), wh(300, 300))},
			want:  wh(100, 100),
		},
		{
			name:  "Asked tighter than Content but fits",
			input: dimensional{Content: r(wh(100, 100), wh(500, 500)), Asked: r(wh(120, 130), wh(200, 250))},
			want:  wh(120, 130),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.input.Dimensions()
			if got != c.want {
				t.Errorf("got %+v, want %+v", got, c.want)
			}
		})
	}
}
