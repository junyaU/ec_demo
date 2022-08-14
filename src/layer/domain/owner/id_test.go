package owner

import "testing"

func TestId_NewExistingId(t *testing.T) {
	tests := []struct {
		id      string
		wantErr bool
	}{
		{id: "OWNER_chsdhdwsihufwhifihe", wantErr: false},
		{id: "hudsifhsiudkgh", wantErr: true},
		{id: "POPUP_hudsifhsiudkgh", wantErr: true},
	}

	for i, test := range tests {
		_, err := newExistingId(test.id)

		if test.wantErr && err == nil {
			t.Fatalf("#%d: want error but no error", i)
		}

		if !test.wantErr && err != nil {
			t.Fatalf("#%d: unexpected error returned \n %#v", i, err)
		}
	}
}
