package popup

import "testing"

func TestPeriod_New(t *testing.T) {
	tests := []struct {
		startTime  string
		endingTime string
		wantErr    bool
	}{
		{
			startTime:  "2022-11-01 13:00:00 JST",
			endingTime: "2022-12-31 13:00:00 JST",
			wantErr:    false,
		},
		{
			startTime:  "2022-12-01 13:00:00 JST",
			endingTime: "2022-08-31 13:00:00 JST",
			wantErr:    true,
		},
		{
			startTime:  "2022-12-01 13:00:00 JST",
			endingTime: "2022-12-01 08:00:00 JST",
			wantErr:    true,
		},
		{
			startTime:  "2022-06-01 13:00:00 JST",
			endingTime: "2022-12-31 13:00:00 JST",
			wantErr:    true,
		},
	}

	for i, test := range tests {
		_, err := newPeriod(test.startTime, test.endingTime)

		if test.wantErr && err == nil {
			t.Fatalf("#%d: want error but no error", i)
		}

		if !test.wantErr && err != nil {
			t.Fatalf("#%d: unexpected error returned \n %#v", i, err)
		}
	}
}

func TestPeriod_isNowOpen(t *testing.T) {
	p, _ := newPeriod("2022-04-01 13:00:00 JST", "2022-04-30 13:00:00 JST")

	if p.isNowOpen() {
		t.Fail()
	}
}
