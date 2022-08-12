package popup

import (
	"github.com/demo/layer/domain"
	"github.com/demo/layer/domain/merchandise"
	"github.com/demo/layer/domain/owner"
)

type openedStoreEvent struct {
	domain.Event
	mainName       string
	shoulderName   string
	promotionText  string
	period         string
	ownerId        string
	merchandiseIds []string
}

func newOpenedStoreEvent(storeId Id, ownerId owner.Id, merchandiseIds []merchandise.Id, storeName storeName, period period, text promotionText) openedStoreEvent {
	var ids []string
	for i, id := range merchandiseIds {
		ids[i] = id.Identifier()
	}

	return openedStoreEvent{
		Event:          domain.NewEvent(storeId.Identifier(), 1, "OPENED_POPUP_STORE"),
		mainName:       storeName.main,
		shoulderName:   storeName.shoulder,
		promotionText:  text.String(),
		period:         period.string(),
		ownerId:        ownerId.Identifier(),
		merchandiseIds: ids,
	}
}
