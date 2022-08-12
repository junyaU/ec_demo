package domain

type Eventer interface {
	Identifier() string
	Version() uint64
	Type() string
}

type Event struct {
	id        string
	version   uint64
	eventType string
}

func NewEvent(id string, version uint64, eventType string) Event {
	return Event{
		id:        id,
		version:   version,
		eventType: eventType,
	}
}

func (e Event) Identifier() string {
	return e.id
}

func (e Event) Version() uint64 {
	return e.version
}

func (e Event) Type() string {
	return e.eventType
}
