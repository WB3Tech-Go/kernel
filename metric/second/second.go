package second

import (
	"github.com/WB3Tech-Go/kernel/measure"
	"github.com/WB3Tech-Go/kernel/measure/tally"
)

type Second struct {
	measure.Measure
	tally.Tally
}

func New(count int) *Second {
	return &Second{
		*measure.New("second", "s"),
		*tally.New(count),
	}
}

func (h *Second) Equals(second Second) bool {
	return h.Count() == second.Count()
}