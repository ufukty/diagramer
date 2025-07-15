package parse

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/ufukty/diagramer/pkg/sequence/parse/ast"
	"github.com/ufukty/diagramer/pkg/sequence/parse/match"
)

func FromReader(src io.Reader) (*ast.Diagram, error) {
	diagram := &ast.Diagram{
		Lifelines:  []*ast.Lifeline{},
		Messages:   []*ast.Message{},
		AutoNumber: false,
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
			diagram.AutoNumber = true

		case strings.HasPrefix(line, "participant") || strings.HasPrefix(line, "actor"):
			if p := match.Lifeline(line); p != nil {
				diagram.Lifelines = append(diagram.Lifelines, p)
			}

		case strings.Contains(line, "->>"):
			if m := match.Message(line); m != nil {
				diagram.Messages = append(diagram.Messages, m)
			}

		default:
			fmt.Printf("skipping line #%d: %s\n", i, line)
		}
	}

	return diagram, nil
}
