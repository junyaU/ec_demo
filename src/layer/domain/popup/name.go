package popup

import "errors"

type name struct {
	main     string
	shoulder string
}

func newName(mainName, shoulderName string) (name, error) {
	if len(mainName) < 1 || len(mainName) > 20 {
		return name{}, errors.New("mainName must be between 1 and 20 characters")
	}

	if len(shoulderName) < 5 || len(shoulderName) > 50 {
		return name{}, errors.New("shoulderName must be between 1 and 20 characters")
	}

	return name{
		main:     mainName,
		shoulder: shoulderName,
	}, nil
}
