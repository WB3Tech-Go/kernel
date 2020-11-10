package amount

type Amount struct {
	quantity float64
}

func New(quantity float64) *Amount {
	return &Amount{quantity: quantity}
}

func (a *Amount) Quantity() float64 {
	return a.quantity
}

func (a *Amount) Equals(amount Amount) bool {
	return a.Quantity() == amount.Quantity()
}
