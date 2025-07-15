package lay

import (
	"github.com/ufukty/diagramer/pkg/sequence/doc"
	"golang.org/x/image/math/fixed"
)

type location struct {
	X, Y fixed.Int26_6
}

type Participant struct {
	Node     *doc.Lifeline
	Location location
}

type Actor struct {
	Node     *doc.Lifeline
	Location location
}

type Message struct {
	Source, Dest *doc.Lifeline
	Message      *doc.Message
	Location     location
}

type Diagram struct {
	Lifelines []any
	Messages  []*Message
}
