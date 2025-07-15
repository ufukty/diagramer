package sequence

import (
	"fmt"
	"image"
	"os"

	"github.com/ufukty/diagramer/pkg/sequence/doc"
	"github.com/ufukty/diagramer/pkg/sequence/internal/text"
	"github.com/ufukty/diagramer/pkg/sequence/lay"
	"github.com/ufukty/diagramer/pkg/sequence/measure"
	"github.com/ufukty/diagramer/pkg/sequence/parse"
)

func FromFile(src string) error {
	file, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("opening: %w", err)
	}
	defer file.Close() //nolint:errcheck

	ast, err := parse.FromReader(file)
	if err != nil {
		return fmt.Errorf("parsing: %w", err)
	}
	explicit := doc.FromAst(ast)
	font, err := text.LoadFont()
	if err != nil {
		return fmt.Errorf("loading font: %w", err)
	}
	measured := measure.FromDoc(explicit, font)
	laid, bounding, err := lay.Out(measured)
	if err != nil {
		return fmt.Errorf("layout: %w", err)
	}
	canvas := image.NewNRGBA64(bounding.Bounds())

	return nil
}
