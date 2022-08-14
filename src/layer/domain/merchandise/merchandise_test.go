package merchandise

import (
	"github.com/demo/layer/domain"
	"testing"
)

func TestMerchandise_Get(t *testing.T) {
	events := []domain.EventModel{
		{
			ID:      MERCHANDISE_ID_PREFIX + "hello",
			Version: 2,
		},
	}

	if _, err := Get(events); err != nil {
		t.Fatal("failed to retrieve merchandise")
	}
}
