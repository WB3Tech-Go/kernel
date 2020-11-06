package z_tests

import (
	"github.com/WB3Tech-Go/kernel/valueObject/metric/minute"
	"testing"
)

func TestNewMinute(t *testing.T) {
	min := minute.New(3)

	if min.Name() != "minute" { t.Error("The name must be 'minute'.")}
	if min.Abbreviation() != "min" { t.Error("The abbreviation must be 'min'.")}
	if min.Count() != 3 { t.Error("The minute must have a count of 3.")}
}

func TestSameMinuteCountMustEqual(t *testing.T) {
	min1 := minute.New(3)
	min2 := minute.New(3)

	if min1.Equals(*min2) != true { t.Error("Two minutes with the same quantity should equal.")}
}

func TestDifferentMinuteCountMustNotEqual(t *testing.T) {
	min1 := minute.New(3)
	min2 := minute.New(7)

	if min1.Equals(*min2) != false { t.Error("Two minutes with the different quantities should not equal.")}
}
