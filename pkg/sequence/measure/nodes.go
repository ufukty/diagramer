// Package measure contains the types that are measured for their
// dimensions and ready to getting laid out onto the canvas.
package measure

import (
	"github.com/ufukty/diagramer/pkg/sequence/doc"
	"golang.org/x/image/math/fixed"
)

type size struct {
	Width, Height fixed.Int26_6
}

type _range struct {
	Min, Max size
}

type dimensional struct {
	Content, Asked _range
}

func clamp(a, b fixed.Int26_6) fixed.Int26_6 {
	return min(a, b)
}

func (d *dimensional) Dimensions() size {
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

type Participant struct {
	Node *doc.Lifeline
}

type Actor struct {
	Node *doc.Lifeline
}

type Message struct {
	Source, Dest *doc.Lifeline
	Message      *doc.Message
}

type Diagram struct {
	Lifelines []Lifeline
	Messages  []*Message
}
