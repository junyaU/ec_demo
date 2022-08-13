package popup

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

const POPUP_ID_PREFIX = "POPUP_"

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
		prefix: POPUP_ID_PREFIX,
		value:  u.String(),
	}, err
}

func (i Id) Identifier() string {
	return i.prefix + i.value
}

func NewExistingId(value string) (Id, error) {
	if !strings.Contains(value, POPUP_ID_PREFIX) {
		return Id{}, errors.New("it is not an owner id")
	}

	return Id{
		prefix: POPUP_ID_PREFIX,
		value:  value[len(POPUP_ID_PREFIX):],
	}, nil
}
