package testkit

import "time"

func MustTime(str string) time.Time {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		panic(err)
	}
	return t
}
