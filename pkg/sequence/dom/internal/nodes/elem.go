package nodes

type Elem struct {
	TagName   string
	Id        string
	ClassList []string
	Style     map[string]string
	Children  []Htmler
}

func NewElement(tag string) *Elem {
	return &Elem{TagName: tag}
}

func (e *Elem) SetClassList(list ...string) *Elem {
	e.ClassList = list
	return e
}

func (e *Elem) SetStyle(style map[string]string) *Elem {
	e.Style = style
	return e
}

func (e *Elem) SetChildren(children ...Htmler) *Elem {
	e.Children = children
	return e
}
