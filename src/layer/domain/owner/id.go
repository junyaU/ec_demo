package owner

import (
	"github.com/google/uuid"
)

type Id struct {
	prefix string
	value  string
}

func NewId() (Id, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return Id{}, err
	}

	return Id{
		prefix: "OWNER_",
		value:  u.String(),
	}, err
}

func (i Id) Identifier() string {
	return i.prefix + i.value
}
