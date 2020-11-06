package z_tests

import (
	"github.com/WB3Tech-Go/kernel/valueObject/metric/mile"
	"testing"
)

func TestNewMile(t *testing.T) {
	mi := mile.New(1.00)

	if mi.Name() != "mile" { t.Error("The metric name should be 'mile'.")}
	if mi.Abbreviation() != "mi" { t.Error("The metric abbreviation should be 'mi'")}
	if mi.Quantity() != 1.00 { t.Error("The kilometer amount should be 1.00.")}
}

