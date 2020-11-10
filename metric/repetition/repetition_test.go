package repetition

import (
	"testing"
)

func TestNewRepetition(t *testing.T) {
	rep := New(8)

	if rep.Name() != "repetition" {
		t.Error("The metric name should be 'repetition'.")
	}
	if rep.Abbreviation() != "rep" {
		t.Error("The metric abbreviation should be 'mi'")
	}
	if rep.Count() != 8 {
		t.Error("The repetition count should be 8.")
	}

}
