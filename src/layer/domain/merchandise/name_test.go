package merchandise

import (
	"strings"
	"testing"
)

func TestName_New(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "すごいシャツ", wantErr: false},
		{name: "", wantErr: true},
		{name: strings.Repeat("a", 51), wantErr: true},
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
