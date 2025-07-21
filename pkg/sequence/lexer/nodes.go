package lexer

type LifelineDecl struct {
	Type  string
	Alias string
	Name  string
}

type Message struct {
	From, To string
	Content  string
}

type Note struct {
	From, To string
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

// MARK: Diagram

type DiagramOpts struct {
	AutoNumber bool
}

type Diagram struct {
	Stmts []Stmt
	Opts  DiagramOpts
}
