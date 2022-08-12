package popup

import (
	"github.com/demo/layer/domain/merchandise"
	"github.com/demo/layer/domain/owner"
)

type Store struct {
	id             Id
	ownerId        owner.Id
	name           string
	promotionText  promotionText
	period         period
	merchandiseIds []merchandise.Id
}

func OpenStore(startTime, endingTime, name, promotionText string, ownerId owner.Id, merchandiseIds []merchandise.Id) (openedStoreEvent, error) {
	storeId, err := newId()
	if err != nil {
		return openedStoreEvent{}, err
	}

	period, err := newPeriod(startTime, endingTime)
	if err != nil {
		return openedStoreEvent{}, err
	}

	pText, err := newPromotionText(promotionText)
	if err != nil {
		return openedStoreEvent{}, err
	}

	return newOpenedStoreEvent(storeId, ownerId, merchandiseIds, name, period, pText), nil
}
