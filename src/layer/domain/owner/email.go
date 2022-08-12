package owner

import (
	"errors"
	"strings"
)

type email struct {
	account string
	domain  string
}

func newEmail(v string) (email, error) {
	emailSplit := strings.Split(v, "@")
	if len(emailSplit) < 2 {
		return email{}, errors.New("not in the form of an email address")
	}

	if emailSplit[0] == "" || emailSplit[1] == "" {
		return email{}, errors.New("invalid email address")
	}

	return email{
		account: emailSplit[0],
		domain:  emailSplit[1],
	}, nil
}

func (e email) show() string {
	return e.account + "@" + e.domain
}
