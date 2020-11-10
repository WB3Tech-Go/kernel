package measure

import (
	"testing"
)

func TestCreteNewMeasure(t *testing.T) {

	mes := New("mile", "mi")

	if mes.Name() != "mile" {
		t.Error("The measure name must be 'mile'.")
	}
	if mes.Abbreviation() != "mi" {
		t.Error("The measure abbreviation must be 'mi'.")
	}

}

func TestSameMeasureEqualEachOther(t *testing.T) {

	mes1 := New("mile", "mi")
	mes2 := New("mile", "mi")

	if mes1.Equals(*mes2) != true {
		t.Error("Measures with the same name and abbreviation must equal.")
	}

}

func TestDifferingMeasureNameMustNoBeEqual(t *testing.T) {

	mes1 := New("mile", "mi")
	mes2 := New("mile-diff", "mi")

	if mes1.Equals(*mes2) != false {
		t.Error("Measures with differing names must not equal.")
	}

}

func TestDifferingMeasureAbbreviationMustNoBeEqual(t *testing.T) {

	mes1 := New("mile", "mi")
	mes2 := New("mile", "mi-diff")

	if mes1.Equals(*mes2) != false {
		t.Error("Measures with differing abbreviations must not equal.")
	}

}
