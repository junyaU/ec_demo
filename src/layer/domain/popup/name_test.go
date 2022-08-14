package popup

import (
	"strings"
	"testing"
)

func TestName_New(t *testing.T) {
	tests := []struct {
		main     string
		shoulder string
		wantErr  bool
	}{
		{main: "junya_shop", shoulder: "おしゃれなおみせ", wantErr: false},
		{main: "", shoulder: "hello hello", wantErr: true},
		{main: strings.Repeat("a", 21), shoulder: "hello_hello", wantErr: true},
		{main: "junya_shop", shoulder: "nice", wantErr: true},
		{main: "junya_shop", shoulder: strings.Repeat("a", 51), wantErr: true},
	}

	for i, test := range tests {
		_, err := newName(test.main, test.shoulder)

		if test.wantErr && err == nil {
			t.Fatalf("#%d: want error but no error", i)

		}

		if !test.wantErr && err != nil {
			t.Fatalf("#%d: unexpected error returned \n %#v", i, err)
		}
	}
}
