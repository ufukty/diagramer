package html

import (
	"fmt"
	"strings"
)

type Htmler interface {
	ToHtml() string
}

func joinmap(kv map[string]string, kvsep, sep string) string {
	lines := []string{}
	for k, v := range kv {
		lines = append(lines, fmt.Sprintf("%s%s%s", k, kvsep, v))
	}
	return strings.Join(lines, sep)
}

func htmlers(hs []Htmler) []string {
	ss := []string{}
	for _, htmler := range hs {
		ss = append(ss, htmler.ToHtml())
	}
	return ss
}

func (n *Elem) ToHtml() string {
	return fmt.Sprintf(`<div class="lifeline" style=%q>%s</div>`,
		joinmap(n.Style, ": ", "; "),
		strings.Join(htmlers(n.Children), " "),
	)
}
