package lexer

import "strings"

func (Activate) check(line string) bool {
	panic("not implemented")
}

func (Alt) check(line string) bool {
	panic("not implemented")
}

func (And) check(line string) bool {
	panic("not implemented")
}

func (Box) check(line string) bool {
	panic("not implemented")
}

func (Break) check(line string) bool {
	panic("not implemented")
}

func (Create) check(line string) bool {
	panic("not implemented")
}

func (Critical) check(line string) bool {
	panic("not implemented")
}

func (Deactivate) check(line string) bool {
	panic("not implemented")
}

func (Destroy) check(line string) bool {
	panic("not implemented")
}

func (Else) check(line string) bool {
	panic("not implemented")
}

func (End) check(line string) bool {
	panic("not implemented")
}

func (LifelineDecl) check(line string) bool {
	return strings.HasPrefix(line, "participant") || strings.HasPrefix(line, "actor")
}

func (Loop) check(line string) bool {
	panic("not implemented")
}

func (Message) check(line string) bool {
	return strings.Contains(line, "->>")
}

func (Note) check(line string) bool {
	panic("not implemented")
}

func (Option) check(line string) bool {
	panic("not implemented")
}

func (Parallel) check(line string) bool {
	panic("not implemented")
}

func (WideNote) check(line string) bool {
	panic("not implemented")
}
