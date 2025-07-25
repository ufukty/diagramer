package doc

type Gridable interface {
	GetContacts() Contacts
}

type Vertical struct {
	Top, Center, Bottom []Gridable
}

type Horizontal struct {
	Left, Center, Right []Gridable
}

type Contacts struct {
	Vertical   Vertical
	Horizontal Horizontal
}

// TODO: remove this hack when Go starts to allow
// generalization based on common fields
func (c Contacts) GetContacts() Contacts {
	return c
}

type GridItemSpec struct {
	// how many columns / rows it needs to span across to allow its
	// contacts to be able to attach from its center when the viewport
	// is large enough
	NoColumns, NoRows int
	// eg. {"auto", "0", "auto"}
	TemplateColumns, TemplateRows []string
}

// TODO: change to use LCM-like commonalization instead of [max]
func specUp(g Gridable) GridItemSpec {
	columns := 1
	if contacts := g.GetContacts().Horizontal.Center; len(contacts) > 0 {
		max := -1
		for _, contact := range contacts {
			if a := specUp(contact); a.NoColumns > max {
				max = a.NoColumns
			}
		}
		if max > -1 {
			// balance further dividends on both left and right
			// for equal cutting contact point in large-enough viewports
			columns = 2 * max
		}
	}

	rows := 1
	if contacts := g.GetContacts().Vertical.Center; len(contacts) > 0 {
		max := -1
		for _, contact := range contacts {
			if a := specUp(contact); a.NoRows > max {
				max = a.NoRows
			}
		}
		if max > -1 {
			// balance further dividends on both left and right
			// for equal cutting contact point in large-enough viewports
			rows = 2 * max
		}
	}

	return GridItemSpec{
		NoColumns: columns,
		NoRows:    rows,
	}
}
