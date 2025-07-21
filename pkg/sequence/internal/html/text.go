package html

type Text struct {
	Content string
}

func NewText(content string) *Text {
	return &Text{Content: content}
}

func (t *Text) ToHtml() string {
	return t.Content
}
