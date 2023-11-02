package helpers

import "time"

func DefaultNullTime() (time.Time, error) {
	timeLayout := "2006-01-02T15:04:05.999Z"

	nullTime, err := time.Parse(timeLayout, timeLayout)
	if err != nil {
		return time.Time{}, err
	}

	return nullTime, nil
}
