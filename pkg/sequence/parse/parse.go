package parse

import (
	"fmt"
	"strings"

	"github.com/ufukty/diagramer/pkg/sequence/ast"
	"github.com/ufukty/diagramer/pkg/sequence/lexer"
)

func Parse(l *lexer.Diagram) (*ast.Diagram, error) {
	diagram := &ast.Diagram{
		Stmts: []ast.Stmt{},
		Opts:  ast.DiagramOpts{},
	}

	errs := []string{}
	lls := map[string]*ast.Lifeline{} // name => node
	stack := []ast.ScopeDefining{}
	for _, stmt := range l.Stmts {
		latest := stack[len(stack)-1]

		switch stmt := stmt.(type) {
		case *lexer.LifelineDecl:
			ll := &ast.Lifeline{
				Type:  stmt.Type,
				Alias: stmt.Alias,
				Name:  stmt.Name,
			}
			lls[ll.Name] = ll
			latest.AppendStmt(ll)

		case *lexer.Message:
			from, ok := lls[stmt.From]
			if !ok {
				errs = append(errs, fmt.Sprintf("the sender %s is not previously declared", from))
			}
			to, ok := lls[stmt.To]
			if !ok {
				errs = append(errs, fmt.Sprintf("the receiver %s is not previously declared", to))
			}
			latest.AppendStmt(&ast.Message{
				From:    from,
				To:      to,
				Content: stmt.Content,
			})

		default:
			fmt.Printf("skipping #v\n")
		}
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("found %s errors:\n", len(errs), strings.Join(errs, "\n"))
	}

	return diagram, nil
}
