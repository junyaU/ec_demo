package popup

import (
	"github.com/google/uuid"
)

type Id struct {
	prefix string
	value  string
}

func newId() (Id, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return Id{}, err
	}

	return Id{
		prefix: "POPUP_",
		value:  u.String(),
	}, err
}
