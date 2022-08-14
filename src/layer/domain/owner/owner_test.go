package owner

import (
	"github.com/demo/layer/domain"
	"testing"
)

func TestOwner_Get(t *testing.T) {
	events := []domain.EventModel{
		{
			ID:      OWNER_ID_PREFIX + "hello",
			Version: 5,
		},
	}

	if _, err := Get(events); err != nil {
		t.Fatal("failed to retrieve owner")
	}
}
