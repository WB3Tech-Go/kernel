package valueObject

import "testing"

func CreateNewPersonObject(t *testing.T) {

	person := NewPerson("Bill", "Bensing")

	if person.FirstName() != "Bill" { t.Error("The first name must equal 'Bill'.")}
	if person.LastName() != "Bensing" { t.Error("The last name must equal 'Bensing'.")}
	if person.FullName() != "Bill Bensing" { t.Error("The full name must equal 'Bill Bensing'.")}

}

func TwoPersonWithSameNameMustEqual(t *testing.T)  {

	person1 := NewPerson("Bill", "Bensing")
	person2 := NewPerson("Bill", "Bensing")

	if !person1.Equals(*person2) { t.Error("Two person instances with the name 'Bill Bensing' must equal each other.")}

}

func TwoPersonWithDifferentNameMustNotEqual(t *testing.T)  {

	person1 := NewPerson("Bill", "Bensing")
	person2 := NewPerson("Bill-2", "Bensing-2")

	if !person1.Equals(*person2) { t.Error("Two person instances with different names must not equal each other.")}

}