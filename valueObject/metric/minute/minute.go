package minute

import (
	"github.com/WB3Tech-Go/kernel/valueObject/measure"
	"github.com/WB3Tech-Go/kernel/valueObject/measure/tally"
)

type Minute struct {
	measure.Measure
	tally.Tally
}

func New(count int) *Minute {
	return &Minute{
		*measure.New("minute", "min"),
		*tally.New(count),
	}
}

func (h *Minute) Equals(minute Minute) bool {
	return h.Count() == minute.Count()
}