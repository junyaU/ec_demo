package popup

import (
	"github.com/demo/layer/domain"
	"github.com/demo/layer/domain/owner"
)

const OPENED_POPUP_STORE_EVENT = "OPENED_POPUP_STORE"

type OpenedStoreEvent struct {
	domain.Event
	mainName      string
	shoulderName  string
	promotionText string
	period        string
	ownerId       string
}

func newOpenedStoreEvent(storeId Id, ownerId owner.Id, storeName name, period period, text promotionText) OpenedStoreEvent {
	return OpenedStoreEvent{
		Event:         domain.NewEvent(storeId.Identifier(), 1, OPENED_POPUP_STORE_EVENT),
		mainName:      storeName.main,
		shoulderName:  storeName.shoulder,
		promotionText: text.String(),
		period:        period.string(),
		ownerId:       ownerId.Identifier(),
	}
}
