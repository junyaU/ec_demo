package merchandise

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

const MERCHANDISE_ID_PREFIX = "MERCHANDISE_"

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
		prefix: MERCHANDISE_ID_PREFIX,
		value:  u.String(),
	}, err
}

func newExistingId(value string) (Id, error) {
	if !strings.Contains(value, MERCHANDISE_ID_PREFIX) {
		return Id{}, errors.New("it is not an merchandise id")
	}

	return Id{
		prefix: MERCHANDISE_ID_PREFIX,
		value:  value[len(MERCHANDISE_ID_PREFIX):],
	}, nil
}

func (i Id) Identifier() string {
	return i.prefix + i.value
}
