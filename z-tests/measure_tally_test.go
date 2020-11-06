package z_tests

import (
"github.com/WB3Tech-Go/kernel/valueObject/measure/tally"
	"testing"
)

func TestNewTally(t *testing.T) {
	tal := tally.New(5)

	if tal.Count() != 5 { t.Error("The tally count is supposed to be 5.")}
}

func TestSameTallyShouldEqual(t *testing.T) {
	tal1 := tally.New(5)
	tal2 := tally.New(5)

	if tal1.Equals(*tal2) != true { t.Error("Tallies with the same count should be equal.")}
}

func TestDifferingTallyShouldNotEqual(t *testing.T) {
	tal1 := tally.New(5)
	tal2 := tally.New(10)

	if tal1.Equals(*tal2) != false { t.Error("Tallies with the different counts should not be equal.")}
}
