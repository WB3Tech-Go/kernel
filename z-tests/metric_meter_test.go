package z_tests

import (
	"github.com/WB3Tech-Go/kernel/valueObject/metric/meter"
	"testing"
)

func TestNewMetric(t *testing.T) {
	mtr := meter.New(200)

	if mtr.Quantity() != 200 { t.Error("The meter quantity must be 200.")}
}

func TestSameMetersMustEqual(t *testing.T) {
	mtr1 := meter.New(200)
	mtr2 := meter.New(200)

	if mtr1.Equals(*mtr2) != true { t.Error("Two meter objects with the same quantities must equal each other.")}
}

func TestDifferingMetersMustNotEqual(t *testing.T) {
	mtr1 := meter.New(200)
	mtr2 := meter.New(500)

	if mtr1.Equals(*mtr2) != false { t.Error("Two meter objects with the differing quantities must not equal each other.")}
}