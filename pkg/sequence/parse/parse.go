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
	for _, stmt := range l.Lines {
		latest := stack[len(stack)-1]

		switch stmt := stmt.(type) {

		case *lexer.Activate:
			panic("not implemented")
		case *lexer.Alt:
			panic("not implemented")
		case *lexer.And:
			panic("not implemented")
		case *lexer.Box:
			panic("not implemented")
		case *lexer.Break:
			panic("not implemented")
		case *lexer.Create:
			panic("not implemented")
		case *lexer.Critical:
			panic("not implemented")
		case *lexer.Deactivate:
			panic("not implemented")
		case *lexer.Destroy:
			panic("not implemented")
		case *lexer.Else:
			panic("not implemented")
		case *lexer.End:
			panic("not implemented")
		case *lexer.LifelineDecl:
			ll := &ast.Lifeline{
				Type:  stmt.Type,
				Alias: stmt.Alias,
				Name:  stmt.Name,
			}
			lls[ll.Name] = ll
			latest.AppendStmt(ll)

		case *lexer.Loop:
			panic("not implemented")
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

		case *lexer.Note:
			panic("not implemented")
		case *lexer.Option:
			panic("not implemented")
		case *lexer.Parallel:
			panic("not implemented")
		case *lexer.WideNote:
			panic("not implemented")

		default:
			fmt.Printf("skipping #v\n")
		}
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("found %d errors:\n%s", len(errs), strings.Join(errs, "\n"))
	}

	return diagram, nil
}
