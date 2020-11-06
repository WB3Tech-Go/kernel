package repetition

import (
	"github.com/WB3Tech-Go/kernel/valueObject/measure"
	"github.com/WB3Tech-Go/kernel/valueObject/measure/tally"
)

type Repetition struct {
	measure.Measure
	tally.Tally
}

func New(count int) *Repetition {
	return &Repetition{
		*measure.New("repetition", "rep"),
		*tally.New(count),
	}
}
