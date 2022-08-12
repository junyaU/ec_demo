package popup

import "github.com/demo/layer/usecase"

type host struct {
	store usecase.EventRepository
}

func NewHost(r usecase.EventRepository) *host {
	return &host{
		store: r,
	}
}

func (h *host) Exec() error {
	return nil
}
