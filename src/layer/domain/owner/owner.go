package owner

import (
	"github.com/demo/layer/domain"
)

type Owner struct {
	id             Id
	name           name
	icon           icon
	selfIntro      selfIntro
	email          email
	isHoldingPopup bool
	version        uint64
}

func Get(events []domain.EventModel) (*Owner, error) {
	lastEvent := events[len(events)-1]

	ownerId, err := newExistingId(lastEvent.ID)
	if err != nil {
		return nil, err
	}

	o := new(Owner)
	o.id = ownerId
	o.version = lastEvent.Version

	// イベント読み込み処理がここに来る

	return o, nil
}

func (o Owner) CanExhibitMerchandise(ownerId Id) bool {
	return o.id.Identifier() == ownerId.Identifier()
}

func (o Owner) CanOpenPopup() bool {
	return o.isHoldingPopup && o.email.show() != ""
}

func (o Owner) Identify() Id {
	return o.id
}
