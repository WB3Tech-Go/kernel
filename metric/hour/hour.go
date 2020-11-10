package hour

import (
	"github.com/WB3Tech-Go/kernel/measure"
	"github.com/WB3Tech-Go/kernel/measure/tally"
)

type Hour struct {
	measure.Measure
	tally.Tally
}

func New(count int) *Hour {
	return &Hour{
		*measure.New("hour", "hr"),
		*tally.New(count),
	}
}

func (h *Hour) Equals(hour Hour) bool {
	return h.Count() == hour.Count()
}