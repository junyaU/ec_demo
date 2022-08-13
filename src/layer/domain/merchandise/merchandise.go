package merchandise

import (
	"github.com/demo/layer/domain"
	"github.com/demo/layer/domain/owner"
)

type Merchandise struct {
	id         Id
	ownerId    owner.Id
	name       name
	itemType   ItemType
	printImage printImage
	price      price
	printArea  PrintArea
	version    uint64
}

func Get(events []domain.EventModel) (*Merchandise, error) {
	lastEvent := events[len(events)-1]

	ownerId, err := newExistingId(lastEvent.ID)
	if err != nil {
		return nil, err
	}

	m := new(Merchandise)
	m.id = ownerId
	m.version = lastEvent.Version

	// イベント読み込み処理
	for _, event := range events {
		switch event {
		}
	}

	return m, nil
}

func (m Merchandise) OwnerIdentify() owner.Id {
	return m.ownerId
}

func (m Merchandise) Identify() Id {
	return m.id
}
