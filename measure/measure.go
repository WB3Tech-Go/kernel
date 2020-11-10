package measure

import "github.com/WB3Tech-Go/kernel/strings"

type Measure struct {
	name string
	abbreviation string
}

func New(name, abbrev string) *Measure {
	return &Measure{
		strings.CleanString(name),
		strings.CleanString(abbrev),
	}
}

func (m *Measure) Name() string {
	return m.name
}

func (m *Measure) Abbreviation() string {
	return m.abbreviation
}

func (m *Measure) Equals(measure Measure) bool {
	sameName := m.name == measure.Name()
	sameAbbrev := m.abbreviation == measure.Abbreviation()
	return sameName && sameAbbrev
}