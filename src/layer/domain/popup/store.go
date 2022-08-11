package popup

import "github.com/demo/layer/domain/owner"

type Store struct {
	id            Id
	ownerId       owner.Id
	name          string
	promotionText promotionText
	period        period
	merchandiseId []int
}

func NewStore(startTime, endingTime, promotionText string) (*Store, error) {
	storeId, err := newId()
	if err != nil {
		return nil, err
	}

	period, err := newPeriod(startTime, endingTime)
	if err != nil {
		return nil, err
	}

	pText, err := newPromotionText(promotionText)
	if err != nil {
		return nil, err
	}

	return &Store{
		id:            storeId,
		period:        period,
		promotionText: pText,
	}, nil
}

func (s Store) Open() {

}
