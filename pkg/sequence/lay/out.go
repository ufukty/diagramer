// Package lay lays out the measured nodes onto the canvas
// applying styling rules
package lay

import (
	"image"

	"github.com/ufukty/diagramer/pkg/sequence/measure"
)

func bounding(l *Diagram) image.Rectangle {
	return image.Rect(0, 0, 1, 1)
}

func Out(m *measure.Diagram) (*Diagram, image.Rectangle, error) {
	l := &Diagram{}
	// for _, l := range m.Lifelines {

	// }
	return l, bounding(l), nil
}
