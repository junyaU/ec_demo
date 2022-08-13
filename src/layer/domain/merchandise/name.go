package merchandise

import "errors"

type name struct {
	value string
}

func newName(text string) (name, error) {
	if len(text) < 1 || len(text) > 50 {
		return name{}, errors.New("name must be between 1 and 50 characters")
	}

	return name{value: text}, nil
}
