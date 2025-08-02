package ast

type Stmt interface {
	_stmt()
}

func (*Alt) _stmt()            {}
func (*Break) _stmt()          {}
func (*CriticalRegion) _stmt() {}
func (*Lifeline) _stmt()       {}
func (*Loop) _stmt()           {}
func (*Message) _stmt()        {}
func (*Parallel) _stmt()       {}
