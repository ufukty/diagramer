package doc

import "slices"

// reuse tracks after previous lifeline destroyed
type tracks []bool

// occupies the first available track passes the after
// after should be -1 for the first track
func (t *tracks) occupy(after int) int {
	if after != -1 {
		if i := slices.Index((*t)[after:], false); i != -1 {
			(*t)[i] = true
			return i
		}
	}
	*t = append(*t, true)
	return len(*t) - 1
}

func (t *tracks) destroy(track int) {
	(*t)[track] = false
}
