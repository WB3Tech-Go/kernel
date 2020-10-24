package z_tests

import (
	"github.com/WB3Tech-Go/kernel/valueObject/person"
	"testing"
)

func CreateNewPersonObject(t *testing.T) {

	per := person.New("Bill", "Bensing")

	if per.FirstName() != "Bill" { t.Error("The first name must equal 'Bill'.")}
	if per.LastName() != "Bensing" { t.Error("The last name must equal 'Bensing'.")}
	if per.FullName() != "Bill Bensing" { t.Error("The full name must equal 'Bill Bensing'.")}

}

func TwoPersonWithSameNameMustEqual(t *testing.T)  {

	per1 := person.New("Bill", "Bensing")
	per2 := person.New("Bill", "Bensing")

	if !per1.Equals(*per2) { t.Error("Two person instances with the name 'Bill Bensing' must equal each other.")}

}

func TwoPersonWithDifferentNameMustNotEqual(t *testing.T)  {

	per1 := person.New("Bill", "Bensing")
	per2 := person.New("Bill-2", "Bensing-2")

	if !per1.Equals(*per2) { t.Error("Two person instances with different names must not equal each other.")}

}