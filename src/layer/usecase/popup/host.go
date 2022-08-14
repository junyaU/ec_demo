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

func (h *host) Exec(command HostCommand) error {
	events, err := h.store.LoadEventStream(command.ownerId)
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

	openEvent, err := popup.OpenStore(command.startTime, command.endingTime, command.mainName, command.shoulderName, command.promotionText, o.Identify())
	if err != nil {
		return err
	}

	return h.store.AppendToStream(openEvent)
}

type HostCommand struct {
	ownerId       string
	startTime     string
	endingTime    string
	mainName      string
	shoulderName  string
	promotionText string
}

func NewHostCommand(ownerId, startTime, endingTime, mainName, shoulderName, promotionText string) HostCommand {
	return HostCommand{
		ownerId:       ownerId,
		startTime:     startTime,
		endingTime:    endingTime,
		mainName:      mainName,
		shoulderName:  shoulderName,
		promotionText: promotionText,
	}
}
