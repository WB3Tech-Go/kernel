package z_tests

import (
	"github.com/WB3Tech-Go/kernel/valueObject/measure"
	"testing"
)

func CreteNewMeasure(t *testing.T) {

	mes := measure.New("mile", "mi")

	if mes.Name() != "mile" { t.Error("The measure name must be 'mile'.")}
	if mes.Abbreviation() != "mi" { t.Error("The measure abbreviation must be 'mi'.")}

}

func SameMeasureEqualEachOther(t *testing.T) {

	mes1 := measure.New("mile", "mi")
	mes2 := measure.New("mile", "mi")

	if mes1.Equals(*mes2) != true { t.Error("Measures with the same name and abbreviation must equal.")}

}

func DifferingMeasureNameMustNoBeEqual(t *testing.T) {

	mes1 := measure.New("mile", "mi")
	mes2 := measure.New("mile-diff", "mi")

	if mes1.Equals(*mes2) != false { t.Error("Measures with differing names must not equal.")}

}

func DifferingMeasureAbbreviationMustNoBeEqual(t *testing.T) {

	mes1 := measure.New("mile", "mi")
	mes2 := measure.New("mile", "mi-diff")

	if mes1.Equals(*mes2) != false { t.Error("Measures with differing abbreviations must not equal.")}

}
