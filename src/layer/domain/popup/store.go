package popup

import (
	"errors"
	"github.com/demo/layer/domain"
	"github.com/demo/layer/domain/merchandise"
	"github.com/demo/layer/domain/owner"
)

type Store struct {
	id             Id
	ownerId        owner.Id
	name           name
	promotionText  promotionText
	period         period
	merchandiseIds []merchandise.Id
	version        uint64
}

func Get(events []domain.EventModel) (*Store, error) {
	lastEvent := events[len(events)-1]

	storeId, err := newExistingId(lastEvent.ID)
	if err != nil {
		return nil, err
	}

	s := new(Store)
	s.id = storeId
	s.version = lastEvent.Version

	// イベント読み込み処理がここに来る

	return s, nil
}

func OpenStore(startTime, endingTime, mainName, shoulderName, promotionText string, ownerId owner.Id) (OpenedStoreEvent, error) {
	storeId, err := newId()
	if err != nil {
		return OpenedStoreEvent{}, err
	}

	period, err := newPeriod(startTime, endingTime)
	if err != nil {
		return OpenedStoreEvent{}, err
	}

	storeName, err := newName(mainName, shoulderName)
	if err != nil {
		return OpenedStoreEvent{}, err
	}

	pText, err := newPromotionText(promotionText)
	if err != nil {
		return OpenedStoreEvent{}, err
	}

	return newOpenedStoreEvent(storeId, ownerId, storeName, period, pText), nil
}

func (s Store) IsCorrectOwner(id owner.Id) bool {
	return id.Identifier() == s.ownerId.Identifier()
}

func (s *Store) ExhibitMerchandise(ids []merchandise.Id) (ExhibitedEvent, error) {
	if len(ids) > 50 {
		return ExhibitedEvent{}, errors.New("can only list up to 50 products")
	}

	return newExhibitedEvent(s.id, ids, s.version), nil
}
