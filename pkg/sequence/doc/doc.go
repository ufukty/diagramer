// Package doc provides the complete AST-like structure for the
// document which is ready to lay out. It makes every detail
// implicitly mentioned by the user explicit.
package doc

import (
	"github.com/ufukty/diagramer/pkg/sequence/parse/ast"
)

type LifelineHead struct {
	*ast.Lifeline // if available
	Track         int
}

type Message struct {
	Ast       *ast.Message
	TrackFrom int
	TrackTo   int
}

type Note struct {
	TrackFrom, TrackTo int
}

// MARK: Single block

type Break struct {
	Description string
	Statements  []Stmt
}

type Loop struct {
	Description string
	Statements  []Stmt
}

// MARK: Multiple blocks

type CriticalRegionBlock struct {
	Description string
	Statements  []Stmt
}

type CriticalRegion struct {
	Blocks []CriticalRegionBlock
}

type ParallelBlock struct {
	Description string
	Statements  []Stmt
}

type Parallel struct {
	Blocks []ParallelBlock
}

type AltBlock struct {
	Description string
	Statements  []Stmt
}

type Alt struct {
	Blocks []AltBlock
}

type Stmt interface {
	_stmt()
}

func (Alt) _stmt()            {}
func (Break) _stmt()          {}
func (CriticalRegion) _stmt() {}
func (Loop) _stmt()           {}
func (Message) _stmt()        {}
func (Parallel) _stmt()       {}

type Diagram struct {
	Ast       *ast.Diagram
	Lifelines []*LifelineHead
	Stmts     []Stmt
}
