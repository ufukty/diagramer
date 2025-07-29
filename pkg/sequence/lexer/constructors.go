package lexer

import (
	"regexp"
	"strings"
)

var (
	regexActivate   = regexp.MustCompile(`activate\s+(\w+)`)
	regexAlt        = regexp.MustCompile(`alt\s+(.+)`)
	regexAnd        = regexp.MustCompile(`and\s+(.+)`)
	regexBox        = regexp.MustCompile(`box(\s+#[0-9A-Fa-f]{6})?\s+(.*)`)
	regexBreak      = regexp.MustCompile(`break\s+(.+)`)
	regexCritical   = regexp.MustCompile(`critical\s+(.+)`)
	regexDeactivate = regexp.MustCompile(`deactivate\s+(\w+)`)
	regexDestroy    = regexp.MustCompile(`destroy\s+(\w+)`)
	regexElse       = regexp.MustCompile(``)
	regexEnd        = regexp.MustCompile(``)
	regexLifeline   = regexp.MustCompile(`(participant|actor)\s+(\w+)(?:\s+as\s+(.+))?`)
	regexLoop       = regexp.MustCompile(``)
	regexMessage    = regexp.MustCompile(`([^\s-]+)\s*(?:->>[-+]?)\s*([^\s:]*)(?::\s*(.+))?`)
	regexNote       = regexp.MustCompile(``)
	regexOption     = regexp.MustCompile(``)
	regexParallel   = regexp.MustCompile(``)
	regexWideNote   = regexp.MustCompile(``)
)

func (Activate) construct(line string) Line {
	a := &Activate{}
	if ms := regexActivate.FindStringSubmatch(line); len(ms) > 0 {
		a.Lifeline = ms[1]
	}
	return a
}

func (Alt) construct(line string) Line {
	a := &Alt{}
	if ms := regexAlt.FindStringSubmatch(line); len(ms) > 0 {
		a.Description = ms[1]
	}
	return a
}

func (And) construct(line string) Line {
	a := &And{}
	if ms := regexAnd.FindStringSubmatch(line); len(ms) > 0 {
		a.Action = ms[1]
	}
	return a
}

func (Box) construct(line string) Line {
	a := &Box{}
	if ms := regexAnd.FindStringSubmatch(line); len(ms)-1 == 2 {
		a.Color = ms[1]
		a.Title = ms[2]
	} else if len(ms)-1 == 1 {
		a.Title = ms[1]
	}
	return a
}

func (Break) construct(line string) Line {
	a := &Break{}
	if ms := regexBreak.FindStringSubmatch(line); len(ms) > 0 {
		a.Description = ms[1]
	}
	return a
}

var lifelineDecl = &LifelineDecl{}

func (Create) construct(line string) Line {
	a := &Create{}
	if ld := lifelineDecl.construct(strings.TrimSpace(strings.TrimPrefix(line, "create"))); ld != nil {
		if ld, ok := ld.(*LifelineDecl); ok {
			a.LifelineDecl = *ld
		}
	}
	return a
}

func (Critical) construct(line string) Line {
	a := &Critical{}
	if ms := regexCritical.FindStringSubmatch(line); len(ms) > 0 {
	}
	return a
}

func (Deactivate) construct(line string) Line {
	a := &Deactivate{}
	if ms := regexDeactivate.FindStringSubmatch(line); len(ms) > 0 {
	}
	return a
}

func (Destroy) construct(line string) Line {
	a := &Destroy{}
	if ms := regexDestroy.FindStringSubmatch(line); len(ms) > 0 {
	}
	return a
}

func (Else) construct(line string) Line {
	a := &Else{}
	if ms := regexElse.FindStringSubmatch(line); len(ms) > 0 {
	}
	return a
}

func (End) construct(line string) Line {
	a := &End{}
	if ms := regexEnd.FindStringSubmatch(line); len(ms) > 0 {
	}
	return a
}

func (LifelineDecl) construct(line string) Line {
	match := regexLifeline.FindStringSubmatch(line)
	if len(match) < 3 {
		return nil
	}
	p := &LifelineDecl{
		Type:  match[1],
		Alias: "",
		Name:  match[2],
	}
	if len(match) > 3 {
		p.Alias = match[3]
	}
	return p
}

func (Loop) construct(line string) Line {
	a := &Loop{}
	if ms := regexLoop.FindStringSubmatch(line); len(ms) > 0 {
	}
	return a
}

func (Message) construct(line string) Line {
	match := regexMessage.FindStringSubmatch(line)
	if len(match) < 4 {
		return nil
	}
	m := &Message{
		From:    match[1],
		To:      match[2],
		Content: "",
	}
	if len(match) > 3 {
		m.Content = match[3]
	}
	return m
}

func (Note) construct(line string) Line {
	a := &Note{}
	if ms := regexNote.FindStringSubmatch(line); len(ms) > 0 {
	}
	return a
}

func (Option) construct(line string) Line {
	a := &Option{}
	if ms := regexOption.FindStringSubmatch(line); len(ms) > 0 {
	}
	return a
}

func (Parallel) construct(line string) Line {
	a := &Parallel{}
	if ms := regexParallel.FindStringSubmatch(line); len(ms) > 0 {
	}
	return a
}

func (WideNote) construct(line string) Line {
	a := &WideNote{}
	if ms := regexWideNote.FindStringSubmatch(line); len(ms) > 0 {
	}
	return a
}
