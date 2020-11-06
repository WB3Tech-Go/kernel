package z_tests

import (
	"github.com/WB3Tech-Go/kernel/valueObject/metric/hour"
	"testing"
)

func TestNewHour(t *testing.T) {
	hr := hour.New(3)

	if hr.Name() != "hour" { t.Error("The name must be 'hour'.")}
	if hr.Abbreviation() != "hr" { t.Error("The abbreviation must be 'hr'.")}
	if hr.Count() != 3 { t.Error("The hour must have a count of 3.")}
}

func TestSameHourCountMustEqual(t *testing.T) {
	hr1 := hour.New(3)
	hr2 := hour.New(3)

	if hr1.Equals(*hr2) != true { t.Error("Two hours with the same quantity should equal.")}
}

func TestDifferentHourCountMustNotEqual(t *testing.T) {
	hr1 := hour.New(3)
	hr2 := hour.New(7)

	if hr1.Equals(*hr2) != false { t.Error("Two hours with the different quantities should not equal.")}
}
