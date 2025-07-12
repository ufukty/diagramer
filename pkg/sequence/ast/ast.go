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
	Lifelines  map[string]*Lifeline
	Messages   []*Message
	AutoNumber bool
}

func (*Lifeline) node() {}
func (*Message) node()  {}
func (*Diagram) node()  {}
