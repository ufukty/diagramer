package measure

type Lifeline interface {
	lifeline()
}

func (*Participant) lifeline() {}
func (*Actor) lifeline()       {}
