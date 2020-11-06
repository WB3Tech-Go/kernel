package kilometer

import "github.com/WB3Tech-Go/kernel/valueObject/measure"

type Kilometer struct {
	measure.Measure
	quantity float64
}

func New(quantity float64) *Kilometer {
	return &Kilometer{
		*measure.New("kilometer", "km"),
		quantity,
	}
}

func (k *Kilometer) Quantity() float64 {
	return k.quantity
}
