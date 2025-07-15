// Package measure contains the types that are measured for their
// dimensions and ready to getting laid out onto the canvas.
package measure

import (
	"github.com/ufukty/diagramer/pkg/sequence/doc"
)

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
