// Package ast defines the types that is just enough to represent
// document with every detail explicitly specified by the user.
//
// AST nodes doesn't represents implied yet undeclared details such
// as lifelines mentioned in links but not declared at the top.
//
// For more complete document representation see the doc package.
package ast

type Node interface {
	node()
}

// Lifeline represents a diagram participant or actor
type Lifeline struct {
	Type  string
	Alias string
	Name  string
}

// Message represents a message between lifelines
type Message struct {
	From    string
	To      string
	Content string
}

type Diagram struct {
	Lifelines  []*Lifeline
	Messages   []*Message
	AutoNumber bool
}

func (*Lifeline) node() {}
func (*Message) node()  {}
func (*Diagram) node()  {}
