package lexer

import (
	"regexp"
)

var (
	regexLifeline = regexp.MustCompile(`(participant|actor)\s+(\w+)(?:\s+as\s+(.+))?`)
	regexMessage  = regexp.MustCompile(`([^\s-]+)\s*(?:->>[-+]?)\s*([^\s:]*)(?::\s*(.+))?`)
)
