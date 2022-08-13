package popup

import (
	"errors"
	"github.com/demo/layer/domain/owner"
	"github.com/demo/layer/domain/popup"
	"github.com/demo/layer/usecase"
)

type host struct {
	store usecase.EventRepository
}

func NewHost(r usecase.EventRepository) *host {
	return &host{
		store: r,
	}
}

func (h *host) Exec(ownerId, startTime, endingTime, mainName, shoulderName, promotionText string) error {
	events, err := h.store.LoadEventStream(ownerId)
	if err != nil {
		return err
	}

	o, err := owner.Get(events)
	if err != nil {
		return err
	}

	if !o.CanOpenPopup() {
		return errors.New("don`t meet the requirements to host a pop-up store")
	}

	openEvent, err := popup.OpenStore(startTime, endingTime, mainName, shoulderName, promotionText, o.Identify())
	if err != nil {
		return err
	}

	return h.store.AppendToStream(openEvent)
}
