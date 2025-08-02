package ast

type Lifeline struct {
	Type  string
	Alias string
	Name  string
}

type Message struct {
	From, To *Lifeline
	Content  string
}

type Note struct {
	From, To *Lifeline
	Content  string
}

// MARK: Single block

type Break struct {
	Description string
	Stmts       []Stmt
}

type Loop struct {
	Description string
	Stmts       []Stmt
}

// MARK: Multiple blocks

type CriticalRegionBlock struct {
	Description string
	Stmts       []Stmt
}

type CriticalRegion struct {
	Blocks []CriticalRegionBlock
}

type ParallelBlock struct {
	Description string
	Stmts       []Stmt
}

type Parallel struct {
	Blocks []ParallelBlock
}

type AltBlock struct {
	Description string
	Stmts       []Stmt
}

type Alt struct {
	Blocks []AltBlock
}

// MARK: Diagram

type DiagramOpts struct {
	AutoNumber bool
}

type Diagram struct {
	Stmts []Stmt
	Opts  DiagramOpts
}
