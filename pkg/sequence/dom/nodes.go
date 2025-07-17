package dom

import (
	"github.com/ufukty/diagramer/pkg/sequence/doc"
	"github.com/ufukty/diagramer/pkg/sequence/parse/ast"
)

type ParticipantHead struct {
	*doc.Lifeline
	Track       int
	StepCreate  int
	StepDestroy int
}

type ActorHead struct {
	*doc.Lifeline
	Track       int
	StepCreate  int
	StepDestroy int
}

type LifelineLine struct {
	*doc.Lifeline
	StepCreate  int
	StepDestroy int
}

type Message struct {
	Ast       *ast.Message
	TrackFrom int
	TrackTo   int
	Step      int
}

type Diagram struct {
	Ast          *ast.Diagram
	Participants []*ParticipantHead
	Actors       []*ActorHead
	Lines        []*LifelineLine
	Messages     []*Message
}
