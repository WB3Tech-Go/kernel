package kilometer

import (
	"testing"
)

func TestCreateNewKilometer(t *testing.T) {

	kilom := New(20.22)

	if kilom.Name() != "kilometer" {
		t.Error("The metric name should be 'kilometer'.")
	}
	if kilom.Abbreviation() != "km" {
		t.Error("The metric abbreviation should be 'km'")
	}
	if kilom.Quantity() != 20.22 {
		t.Error("The kilometer amount should be 20.22.")
	}

}
