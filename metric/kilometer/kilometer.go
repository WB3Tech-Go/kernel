package kilometer

import (
	"github.com/WB3Tech-Go/kernel/measure"
	"github.com/WB3Tech-Go/kernel/measure/amount"
)

type Kilometer struct {
	measure.Measure
	amount.Amount
}

func New(quantity float64) *Kilometer {
	return &Kilometer{
		*measure.New("kilometer", "km"),
		*amount.New(quantity),
	}
}