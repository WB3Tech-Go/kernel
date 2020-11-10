package person

import (
	"testing"
)

func TestCreateNewPersonObject(t *testing.T) {

	per := New("Bill", "Bensing")

	if per.FirstName() != "Bill" { t.Error("The first name must equal 'Bill'.")}
	if per.LastName() != "Bensing" { t.Error("The last name must equal 'Bensing'.")}
	if per.FullName() != "Bill Bensing" { t.Error("The full name must equal 'Bill Bensing'.")}

}

func TestTwoPersonWithSameNameMustEqual(t *testing.T)  {

	per1 := New("Bill", "Bensing")
	per2 := New("Bill", "Bensing")

	if !per1.Equals(*per2) { t.Error("Two person instances with the name 'Bill Bensing' must equal each other.")}

}

func TestTwoPersonWithDifferentNameMustNotEqual(t *testing.T)  {

	per1 := New("Bill", "Bensing")
	per2 := New("Bill-2", "Bensing-2")

	if per1.Equals(*per2) { t.Error("Two person instances with different names must not equal each other.")}

}