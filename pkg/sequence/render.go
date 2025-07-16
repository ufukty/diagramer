package sequence

import (
	"fmt"
	"image"
	"io"

	"github.com/ufukty/diagramer/pkg/sequence/doc"
	"github.com/ufukty/diagramer/pkg/sequence/dom"
	"github.com/ufukty/diagramer/pkg/sequence/internal/text"
	"github.com/ufukty/diagramer/pkg/sequence/lay"
	"github.com/ufukty/diagramer/pkg/sequence/measure"
	"github.com/ufukty/diagramer/pkg/sequence/parse"
)

func Render(dst io.Writer, src io.Reader) error {
	ast, err := parse.FromReader(src)
	if err != nil {
		return fmt.Errorf("parsing: %w", err)
	}
	explicit := doc.FromAst(ast)
	dom := dom.Build(explicit)
	font, err := text.LoadFont("Arial", 14)
	if err != nil {
		return fmt.Errorf("loading font: %w", err)
	}
	measured := measure.Measure(dom, font)
	laid, bounding, err := lay.Out(dom)
	if err != nil {
		return fmt.Errorf("layout: %w", err)
	}
	canvas := image.NewNRGBA64(bounding.Bounds())
	return nil
}
