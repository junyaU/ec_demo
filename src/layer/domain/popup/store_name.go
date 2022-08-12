package popup

import "errors"

type storeName struct {
	main     string
	shoulder string
}

func newStoreName(mainName, shoulderName string) (storeName, error) {
	if len(mainName) < 1 || len(mainName) > 20 {
		return storeName{}, errors.New("mainName must be between 1 and 20 characters")
	}

	if len(shoulderName) < 5 || len(shoulderName) > 50 {
		return storeName{}, errors.New("shoulderName must be between 1 and 20 characters")
	}

	return storeName{
		main:     mainName,
		shoulder: shoulderName,
	}, nil
}
