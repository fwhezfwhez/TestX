package main

import (
	"fmt"
	"time"
)

func main() {
fmt.Println([]byte{})
	fmt.Println([]byte(""))

	fmt.Println(int(int32(50)))
}

func GetMondayDate(today time.Time) time.Time {
	if today.IsZero() {
		today = time.Now()
	}
	switch today.Weekday() {
	case time.Monday:
		return today
	case time.Tuesday:
		return today.Add(-24 * time.Hour)
	case time.Wednesday:
		return today.Add(-2 * 24 * time.Hour)
	case time.Thursday:
		return today.Add(-3 * 24 * time.Hour)
	case time.Friday:
		return today.Add(-4 * 24 * time.Hour)
	case time.Saturday:
		return today.Add(-5 * 24 * time.Hour)
	case time.Sunday:
		return today.Add(-6 * 24 * time.Hour)
	}
	return time.Now()
}
