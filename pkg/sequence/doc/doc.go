// Package doc provides the complete AST-like structure for the
// document which is ready to lay out. It makes every detail
// implicitly mentioned by the user explicit.
package doc

import (
	"github.com/ufukty/diagramer/pkg/sequence/parser/ast"
)

type Lifeline struct {
	*ast.Lifeline // if available
	Track         int
	StepCreate    int
	StepDestroy   int // the steps the lifeline gets created and destroyed
}

type Message struct {
	Ast       *ast.Message
	TrackFrom int
	TrackTo   int
	Step      int
}

type Diagram struct {
	Ast       *ast.Diagram
	Lifelines []*Lifeline
	Messages  []*Message
}
