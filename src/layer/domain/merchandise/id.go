package merchandise

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
		prefix: "MERCHANDISE_",
		value:  u.String(),
	}, err
}

func (i Id) Identifier() string {
	return i.prefix + i.value
}
