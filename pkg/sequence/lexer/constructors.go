package lexer

func (Activate) construct(line string) Line {
	panic("not implemented")
}

func (Alt) construct(line string) Line {
	panic("not implemented")
}

func (And) construct(line string) Line {
	panic("not implemented")
}

func (Box) construct(line string) Line {
	panic("not implemented")
}

func (Break) construct(line string) Line {
	panic("not implemented")
}

func (Create) construct(line string) Line {
	panic("not implemented")
}

func (Critical) construct(line string) Line {
	panic("not implemented")
}

func (Deactivate) construct(line string) Line {
	panic("not implemented")
}

func (Destroy) construct(line string) Line {
	panic("not implemented")
}

func (Else) construct(line string) Line {
	panic("not implemented")
}

func (End) construct(line string) Line {
	panic("not implemented")
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
	panic("not implemented")
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
	panic("not implemented")
}

func (Option) construct(line string) Line {
	panic("not implemented")
}

func (Parallel) construct(line string) Line {
	panic("not implemented")
}

func (WideNote) construct(line string) Line {
	panic("not implemented")
}
