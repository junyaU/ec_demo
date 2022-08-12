package owner

import "testing"

func TestEmail_New(t *testing.T) {
	tests := []struct {
		email   string
		wantErr bool
	}{
		{email: "hello@gmail.com", wantErr: false},
		{email: "hogehoge", wantErr: true},
		{email: "@gmail.com", wantErr: true},
		{email: "hello@", wantErr: true},
	}

	for i, test := range tests {
		_, err := newEmail(test.email)

		if test.wantErr && err == nil {
			t.Fatalf("#%d: want error but no error", i)

		}

		if !test.wantErr && err != nil {
			t.Fatalf("#%d: unexpected error returned \n %#v", i, err)
		}
	}
}
