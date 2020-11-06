package z_tests

import (
	"github.com/WB3Tech-Go/kernel/valueObject/metric/second"
	"testing"
)

func TestNewSecond(t *testing.T) {
	s := second.New(3)

	if s.Name() != "second" { t.Error("The name must be 'second'.")}
	if s.Abbreviation() != "s" { t.Error("The abbreviation must be 's'.")}
	if s.Count() != 3 { t.Error("The second must have a count of 3.")}
}

func TestSameSecondCountMustEqual(t *testing.T) {
	s1 := second.New(3)
	s2 := second.New(3)

	if s1.Equals(*s2) != true { t.Error("Two seconds with the same quantity should equal.")}
}

func TestDifferentSecondCountMustNotEqual(t *testing.T) {
	s1 := second.New(3)
	s2 := second.New(7)

	if s1.Equals(*s2) != false { t.Error("Two seconds with the different quantities should not equal.")}
}

