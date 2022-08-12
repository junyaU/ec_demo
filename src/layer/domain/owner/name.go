package owner

import (
	"errors"
	"regexp"
)

type name struct {
	value string
}

func newName(ownerName string) (name, error) {
	rep := regexp.MustCompile(`^[0-9a-zA-Z]*$`)
	if !rep.MatchString(ownerName) {
		return name{}, errors.New("only one-byte alphanumeric characters can be registered")
	}

	if len(ownerName) < 5 || len(ownerName) > 20 {
		return name{}, errors.New("name must be between 5 and 20 characters")
	}

	return name{
		value: ownerName,
	}, nil
}

func (n name) show() string {
	return n.value
}
