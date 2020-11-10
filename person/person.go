package person

import "github.com/WB3Tech-Go/kernel/strings"

type Person struct {
	fistName string
	lastName string
}

func New(firstName, lastName string) *Person {
	return &Person{
		fistName: strings.CleanString(firstName),
		lastName: strings.CleanString(lastName),
	}
}

func (p *Person) FirstName() string  {
	return p.fistName
}

func (p *Person) LastName() string {
	return p.lastName
}

func (p *Person) FullName() string {
	return p.fistName + " " + p.lastName
}

func (p *Person) Equals(person Person) bool {

	first := p.fistName == person.FirstName()
	last := p.lastName == person.LastName()

	return first && last
}