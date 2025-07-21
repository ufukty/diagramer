package ast

type Lifeline struct {
	Type  string
	Alias string
	Name  string
}

type Message struct {
	TrackFrom int
	TrackTo   int
	Content   string
}

type Note struct {
	From, To *Lifeline
	Content  string
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

type Diagram struct {
	Stmts []Stmt
}
