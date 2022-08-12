package owner

import "errors"

type selfIntro struct {
	text string
}

func newSelfIntro(value string) (selfIntro, error) {
	if len(value) < 20 || len(value) > 500 {
		return selfIntro{}, errors.New("selfIntro must be between 20 and 500 characters")
	}

	return selfIntro{
		text: value,
	}, nil
}
