package mile

import (
	"github.com/WB3Tech-Go/kernel/measure"
	"github.com/WB3Tech-Go/kernel/measure/amount"
)

type Mile struct {
	measure.Measure
	amount.Amount
}

func New(quantity float64) *Mile {
	return &Mile{
		*measure.New("mile", "mi"),
		*amount.New(quantity),
	}
}