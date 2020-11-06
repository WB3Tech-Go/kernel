package z_tests

import (
	"github.com/WB3Tech-Go/kernel/valueObject/username"
	"testing"
)

func TestCreateNewUsername(t *testing.T) {
	un, _ := username.New("SomeCoolName")
	if un.Value() != "SomeCoolName" { t.Error("The username value must be 'SomeCoolName'.")}
}

func TestUsernameMustBeAtLeast3Characters(t *testing.T) {
	_, err := username.New("ab")
	if err == nil { t.Error("Error result is nil, an error must have been thrown with the message 'Username must be at least 3 characters.'")}
}
