package meter

import (
	"github.com/WB3Tech-Go/kernel/valueObject/measure"
	"github.com/WB3Tech-Go/kernel/valueObject/measure/amount"
)

type Meter struct {
	measure.Measure
	amount.Amount
}

func New(quantity float64) *Meter {
	return &Meter{
		*measure.New("meter", "m"),
		*amount.New(quantity),
	}
}

func (m *Meter) Equals(meter Meter) bool {
	return m.Quantity() == meter.Quantity()
}
