package owner

import "testing"

func TestIcon_New(t *testing.T) {
	tests := []struct {
		path    string
		wantErr bool
	}{
		{path: "helloworld.jpg", wantErr: false},
		{path: "hogehoge", wantErr: true},
		{path: "hello.svg", wantErr: true},
		{path: ".png", wantErr: true},
	}

	for i, test := range tests {
		_, err := newIcon(test.path)

		if test.wantErr && err == nil {
			t.Fatalf("#%d: want error but no error", i)
		}

		if !test.wantErr && err != nil {
			t.Fatalf("#%d: unexpected error returned \n %#v", i, err)
		}
	}

}
