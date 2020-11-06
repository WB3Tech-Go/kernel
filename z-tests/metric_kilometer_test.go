package z_tests

import (
	"github.com/WB3Tech-Go/kernel/valueObject/metric/kilometer"
	"testing"
)

func TestCreateNewKilometer(t *testing.T) {

	kilom := kilometer.New(20.22)

	if kilom.Name() != "kilometer" { t.Error("The metric name should be 'kilometer'.")}
	if kilom.Abbreviation() != "km" { t.Error("The metric abbreviation should be 'km'")}
	if kilom.Quantity() != 20.22 { t.Error("The kilometer quantity should be 20.22.")}

}
