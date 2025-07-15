package doc

import (
	"github.com/ufukty/diagramer/pkg/sequence/parse/ast"
)

func has[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

func FromAst(src *ast.Diagram) *Diagram {
	dst := &Diagram{}

	occupied := tracks{}
	lookup := map[string]*Lifeline{} // by actor/participant name
	for _, ll := range src.Lifelines {
		lookup[ll.Name] = &Lifeline{
			Lifeline:    ll,
			Track:       len(occupied),
			StepCreate:  0,
			StepDestroy: len(src.Lifelines) + 1, // doc end
		}
		occupied = append(occupied, true)
	}

	for _, m := range src.Messages {
		if !has(lookup, m.From) {

		}
		if !has(lookup, m.To) {

		}
	}

	return dst
}
