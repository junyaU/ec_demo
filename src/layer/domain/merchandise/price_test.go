package merchandise

import "testing"

func TestPrice_New(t *testing.T) {
	tests := []struct {
		value    int
		itemType ItemType
		wantErr  bool
	}{
		{value: 3000, itemType: T_SHIRT, wantErr: false},
		{value: -300, itemType: PHONE_CASE, wantErr: true},
		{value: 10000, itemType: CAP, wantErr: true},
	}

	for i, test := range tests {
		_, err := newPrice(test.itemType, test.value)

		if test.wantErr && err == nil {
			t.Fatalf("#%d: want error but no error", i)
		}

		if !test.wantErr && err != nil {
			t.Fatalf("#%d: unexpected error returned \n %#v", i, err)
		}
	}
}

func TestPrice_DisplayTotalAmount(t *testing.T) {
	expectedPrice := 4620
	p, _ := newPrice(T_SHIRT, 2000)

	if p.displayTotalAmount() != expectedPrice {
		t.Fatal("not the expected amount")
	}
}
