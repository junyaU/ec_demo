package popup

import (
	"github.com/demo/layer/domain/merchandise"
	"github.com/demo/layer/domain/owner"
)

type openedStoreEvent struct {
	Id             string
	name           string
	promotionText  string
	period         string
	ownerId        string
	merchandiseIds []string
	version        uint64
	eventType      string
}

func newOpenedStoreEvent(storeId Id, ownerId owner.Id, merchandiseIds []merchandise.Id, name string, period period, text promotionText) openedStoreEvent {
	var ids []string
	for i, id := range merchandiseIds {
		ids[i] = id.Identifier()
	}

	return openedStoreEvent{
		Id:             storeId.Identifier(),
		name:           name,
		promotionText:  text.String(),
		period:         period.string(),
		ownerId:        ownerId.Identifier(),
		merchandiseIds: ids,
		version:        1,
		eventType:      "OPENED_POPUP_STORE",
	}
}
