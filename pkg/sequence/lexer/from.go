// Package lexer builds a node tree representing the user doc.
// The result differs from an AST with its lack to symbol based
// input validation, and reference between symbols. It produces
// an AST solely based on literal values.
package lexer

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func FromReader(src io.Reader) (*Diagram, error) {
	diagram := &Diagram{
		Stmts: []Stmt{},
		Opts:  DiagramOpts{},
	}

	scanner := bufio.NewScanner(src)
	for i := 0; scanner.Scan(); i++ {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "%%") {
			continue
		}

		switch {
		case line == "sequenceDiagram":
			continue

		case line == "autoNumber":
			diagram.Opts.AutoNumber = true

		case strings.HasPrefix(line, "participant") || strings.HasPrefix(line, "actor"):
			if ll := lexLifeline(line); ll != nil {
				diagram.Stmts = append(diagram.Stmts, ll)
			}

		case strings.Contains(line, "->>"):
			if m := lexMessage(line); m != nil {
				diagram.Stmts = append(diagram.Stmts, m)
			}

		default:
			fmt.Printf("skipping line #%d: %s\n", i, line)
		}
	}

	return diagram, nil
}
