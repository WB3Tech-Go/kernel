package repetition

import "github.com/WB3Tech-Go/kernel/valueObject/measure"

type Repetition struct {
	measure.Measure
	count int
}

func New(count int) *Repetition {
	return &Repetition{
		*measure.New("repetition", "rep"),
		count,
	}
}

func (r *Repetition) Count() int {
	return r.count
}
