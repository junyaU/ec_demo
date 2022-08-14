package popup

import (
	"strings"
	"testing"
)

func TestPromotionText_New(t *testing.T) {
	tests := []struct {
		text    string
		wantErr bool
	}{
		{text: strings.Repeat("a", 150), wantErr: false},
		{text: "hello", wantErr: true},
		{text: strings.Repeat("a", 520), wantErr: true},
	}

	for i, test := range tests {
		_, err := newPromotionText(test.text)

		if test.wantErr && err == nil {
			t.Fatalf("#%d: want error but no error", i)

		}

		if !test.wantErr && err != nil {
			t.Fatalf("#%d: unexpected error returned \n %#v", i, err)
		}
	}
}
