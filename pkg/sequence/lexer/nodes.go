package lexer

import "github.com/ufukty/diagramer/pkg/sequence/lexer/tokens"

// MARK: Lifelines and rel.

type LifelineDecl struct {
	Type  string
	Alias string
	Name  string
}

type Create struct {
	LifelineDecl
}

type Destroy struct {
	LifelineDecl
}

// Box expects [End]
type Box struct {
	Color string
	Title string
}

type Activate struct {
	Lifeline string
}

type Deactivate struct {
	Lifeline string
}

// MARK: Messages and rel.

type Message struct {
	From, To   string
	Content    string
	Activation tokens.Activation
}

type Note struct {
	Lifeline string
	Pos      tokens.NotePos
	Content  string
}

type WideNote struct {
	From, To string
	Content  string
}

// MARK: Single block

// Break expects [End]
type Break struct {
	Description string
}

// Loop expects [End]
type Loop struct {
	Description string
}

// MARK: Multiple blocks

// Critical accepts [Option], expects [End]
type Critical struct {
	Description string
}

// Option can be used inside [Critical]
type Option struct {
	Description string
}

// End is used after [Alt], [Box], [Break], [Critical], [Loop], [Parallel]
type End struct{}

type Parallel struct {
	Action string
}

// And is used after [Parallel]
type And struct {
	Action string
}

type Alt struct {
	Description string
}

// Else is used after [Alt], [Parallel]
type Else struct {
	Description string
}

// MARK: Diagram

type DiagramOpts struct {
	AutoNumber bool
}

type Diagram struct {
	Opts  DiagramOpts
	Lines []Line
}
