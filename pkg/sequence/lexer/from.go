// Package lexer tokenizes the lines into nodes that has limited hierarchy.
package lexer

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func FromReader(src io.Reader) (*Diagram, error) {
	diagram := &Diagram{
		Lines: []Line{},
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
			if ll := mLifelineDecl(line); ll != nil {
				diagram.Lines = append(diagram.Lines, ll)
			}

		case strings.Contains(line, "->>"):
			if m := mMessage(line); m != nil {
				diagram.Lines = append(diagram.Lines, m)
			}

		default:
			fmt.Printf("skipping line #%d: %s\n", i, line)
		}
	}

	return diagram, nil
}
