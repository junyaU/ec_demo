package popup

import (
	"github.com/demo/layer/domain"
	"github.com/demo/layer/domain/merchandise"
	"github.com/demo/layer/domain/owner"
	"strings"
	"testing"
)

func TestStore_Get(t *testing.T) {
	events := []domain.EventModel{
		{
			ID: POPUP_ID_PREFIX + "hello",
		},
	}

	if _, err := Get(events); err != nil {
		t.Fatal("failed to retrieve popup store")
	}
}

func TestStore_OpenStore(t *testing.T) {
	_, err := OpenStore(
		"2022-11-01 13:00:00 JST",
		"2022-12-31 13:00:00 JST",
		"junya_shop",
		"おしゃれなふくやさん",
		strings.Repeat("a", 300),
		owner.Id{},
	)

	if err != nil {
		t.Fail()
	}
}

func TestStore_ExhibitMerchandise(t *testing.T) {
	events := []domain.EventModel{
		{
			ID: POPUP_ID_PREFIX + "hello",
		},
	}

	store, _ := Get(events)

	var merchandiseIds []merchandise.Id
	for i := 0; i < 100; i++ {
		merchandiseIds = append(merchandiseIds, merchandise.Id{})
	}

	if _, err := store.ExhibitMerchandise(merchandiseIds); err == nil {
		t.Fail()
	}
}
