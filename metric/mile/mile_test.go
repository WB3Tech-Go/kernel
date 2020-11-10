package mile

import (
	"testing"
)

func TestNewMile(t *testing.T) {
	mi := New(1.00)

	if mi.Name() != "mile" {
		t.Error("The metric name should be 'mile'.")
	}
	if mi.Abbreviation() != "mi" {
		t.Error("The metric abbreviation should be 'mi'")
	}
	if mi.Quantity() != 1.00 {
		t.Error("The kilometer amount should be 1.00.")
	}
}
