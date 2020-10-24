package username

import (
	"errors"
	"github.com/WB3Tech-Go/kernel/strings"
)

type Username struct {
	value string
}

func New(username string) (*Username, error) {
	if len(username) < 3 { return &Username{}, errors.New("hello") }
	return &Username{
		value: strings.CleanString(username),
	}, nil
}

func (u *Username) Value() string {
	return u.value
}
