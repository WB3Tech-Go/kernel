package z_tests

import (
	"github.com/WB3Tech-Go/kernel/valueObject/username"
	"testing"
)

func CreateNewUsername(t *testing.T) {
	un, _ := username.New("SomeCoolName")
	if un.Value() != "SomeCoolName" { t.Error("The username value must be 'SomeCoolName'.")}
}

