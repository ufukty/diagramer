package lexer

import "strings"

func startsWith(line, word string) bool {
	return strings.Split(line, " ")[0] == word
}

func (Activate) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "activate")
}

func (Alt) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "alt")
}

func (And) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "and")
}

func (Box) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "box")
}

func (Break) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "break")
}

func (Create) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "create")
}

func (Critical) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "critical")
}

func (Deactivate) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "deactivate")
}

func (Destroy) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "destroy")
}

func (Else) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "else")
}

func (End) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "end")
}

func (LifelineDecl) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "participant") || startsWith(strings.TrimSpace(line), "actor")
}

func (Loop) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "loop")
}

func (Message) check(line string) bool {
	return strings.Contains(line, "->>")
}

func (Note) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "note") && !strings.Contains(strings.Split(line, ":")[0], ",")
}

func (Option) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "opt")
}

func (Parallel) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "par")
}

func (WideNote) check(line string) bool {
	return startsWith(strings.TrimSpace(line), "note") && strings.Contains(strings.Split(line, ":")[0], ",")
}
