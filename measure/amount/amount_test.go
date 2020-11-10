package amount

import (
	"testing"
)

func TestNewQuantity(t *testing.T) {
	amt := New(3.5)

	if amt.Quantity() != 3.5 { t.Error("The amount is supposed to be 3.5.")}
}

func TestSameQuantityShouldEqual(t *testing.T) {
	amt1 := New(3.5)
	amt2 := New(3.5)

	if amt1.Equals(*amt2) != true { t.Error("Amounts with the same amount should be equal.")}
}

func TestDifferingQuantityShouldNotEqual(t *testing.T) {
	amt1 := New(5)
	amt2 := New(3.5)

	if amt1.Equals(*amt2) != false { t.Error("Amounts with the different quantities should not be equal.")}
}
