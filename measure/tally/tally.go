package tally

type Tally struct {
	count int
}

func New(count int) *Tally {
	return &Tally{count: count}
}

func (t *Tally) Count() int {
	return t.count
}

func (t *Tally) Equals(tally Tally) bool {
	return t.Count() == tally.Count()
}
