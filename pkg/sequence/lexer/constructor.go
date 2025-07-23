package lexer

// selective line constructor
type constructor interface {
	check(string) bool
	construct(string) Line
}

// internal purposes only
var precedence = []constructor{
	&Activate{},
	&Alt{},
	&And{},
	&Box{},
	&Break{},
	&Create{},
	&Critical{},
	&Deactivate{},
	&Destroy{},
	&Else{},
	&End{},
	&LifelineDecl{},
	&Loop{},
	&Message{},
	&Note{},
	&Option{},
	&Parallel{},
	&WideNote{},
}
