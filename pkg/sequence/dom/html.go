package dom

import (
	"fmt"

	"github.com/ufukty/diagramer/pkg/sequence/dom/internal/html"
)

func (n ActorHead) ToHtml() string {
	return html.
		NewElement("div").
		SetClassList("actor").
		SetStyle(map[string]string{
			"grid-template-rows":    fmt.Sprintf("%d / %d", n.StepCreate, n.StepDestroy),
			"grid-template-columns": fmt.Sprintf("%d / %d", n.Track, n.Track+1),
		}).
		SetChildren(
			html.NewText(n.Alias),
		).
		ToHtml()
}

func (n ParticipantHead) ToHtml() string {
	return html.
		NewElement("div").
		SetClassList("participant").
		SetStyle(map[string]string{
			"grid-template-rows":    fmt.Sprintf("%d / %d", n.StepCreate, n.StepDestroy),
			"grid-template-columns": fmt.Sprintf("%d / %d", n.Track, n.Track+1),
		}).
		SetChildren(
			html.NewText(n.Alias),
		).
		ToHtml()
}

func (n Message) ToHtml() string {
	// return toHtml("div", nil, )
	// Ast       *ast.Message
	// TrackFrom int
	// TrackTo   int
	// Step      int
	return html.NewElement("div").
		ToHtml()
}

func (n Diagram) ToHtml() string {
	// return toHtml("div", nil, )
	// Ast       *ast.Diagram
	// Lifelines []*Lifeline
	// Messages  []*Message
	return html.NewElement("div").
		ToHtml()
}
