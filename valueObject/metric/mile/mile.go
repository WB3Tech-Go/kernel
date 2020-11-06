package mile

import "github.com/WB3Tech-Go/kernel/valueObject/measure"

type Mile struct {
	measure.Measure
	quantity float64
}

func New(quantity float64) *Mile {
	return &Mile{
		*measure.New("mile", "mi"),
		quantity,
	}
}

func (q *Mile) Quantity() float64 {
	return q.quantity
}