package owner

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

const OWNER_ID_PREFIX = "OWNER_"

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
		prefix: OWNER_ID_PREFIX,
		value:  u.String(),
	}, err
}

func (i Id) Identifier() string {
	return i.prefix + i.value
}

func NewExistingId(value string) (Id, error) {
	if !strings.Contains(value, OWNER_ID_PREFIX) {
		return Id{}, errors.New("it is not an owner id")
	}

	return Id{
		prefix: OWNER_ID_PREFIX,
		value:  value[len(OWNER_ID_PREFIX):],
	}, nil
}
