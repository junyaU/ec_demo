package popup

import (
	"errors"
	"time"
)

type period struct {
	startTime  time.Time
	endingTime time.Time
}

func newPeriod(start, end string) (period, error) {
	layout := "2006-01-02 15:04:05 MST"

	sTime, err := time.Parse(layout, start)
	if err != nil {
		return period{}, err
	}

	eTime, err := time.Parse(layout, end)
	if err != nil {
		return period{}, err
	}

	if sTime.After(eTime) {
		return period{}, errors.New("startTime must be earlier than endingTime")
	}

	if sTime.Before(time.Now()) {
		return period{}, errors.New("startTime must be set after the current time")
	}

	return period{
		startTime:  sTime,
		endingTime: eTime,
	}, nil
}

func (p period) isNowOpen() bool {
	now := time.Now()
	return now.After(p.startTime) && now.Before(p.endingTime)
}
