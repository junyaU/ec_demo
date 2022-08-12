package owner

import "testing"

func TestName_New(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "junyaU", wantErr: false},
		{name: "taro", wantErr: true},
		{name: "たろう", wantErr: true},
		{name: "aaaaaaaaaaaaaaaaaaaaaaaa", wantErr: true},
	}

	for i, test := range tests {
		_, err := newName(test.name)

		if test.wantErr && err == nil {
			t.Fatalf("#%d: want error but no error", i)

		}

		if !test.wantErr && err != nil {
			t.Fatalf("#%d: unexpected error returned \n %#v", i, err)
		}
	}
}
