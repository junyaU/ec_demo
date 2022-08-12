package popup

import "errors"

type promotionText struct {
	value string
}

func newPromotionText(text string) (promotionText, error) {
	if len(text) < 100 || len(text) > 500 {
		return promotionText{}, errors.New("promotionText must be between 100 and 500 characters")
	}
	return promotionText{
		value: "",
	}, nil
}

func (t promotionText) String() string {
	return t.value
}
