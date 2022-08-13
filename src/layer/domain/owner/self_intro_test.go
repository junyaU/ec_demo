package owner

import (
	"strings"
	"testing"
)

func TestSelfIntro_New(t *testing.T) {
	tests := []struct {
		text    string
		wantErr bool
	}{
		{text: "よろしくお願いしますよろしくお願いしますよろしくお願いします", wantErr: false},
		{text: "taro", wantErr: true},
		{text: strings.Repeat("a", 550), wantErr: true},
	}

	for i, test := range tests {
		_, err := newSelfIntro(test.text)

		if test.wantErr && err == nil {
			t.Fatalf("#%d: want error but no error", i)

		}

		if !test.wantErr && err != nil {
			t.Fatalf("#%d: unexpected error returned \n %#v", i, err)
		}
	}
}
