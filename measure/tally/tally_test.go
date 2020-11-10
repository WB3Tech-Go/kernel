package tally

import (
	"testing"
)

func TestNewTally(t *testing.T) {
	tal := New(5)

	if tal.Count() != 5 { t.Error("The tally count is supposed to be 5.")}
}

func TestSameTallyShouldEqual(t *testing.T) {
	tal1 := New(5)
	tal2 := New(5)

	if tal1.Equals(*tal2) != true { t.Error("Tallies with the same count should be equal.")}
}

func TestDifferingTallyShouldNotEqual(t *testing.T) {
	tal1 := New(5)
	tal2 := New(10)

	if tal1.Equals(*tal2) != false { t.Error("Tallies with the different counts should not be equal.")}
}
