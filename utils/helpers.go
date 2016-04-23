package utils

import (
	"fmt"
	"strings"
)

const (
	Nanosecond  int64 = 1
	Microsecond       = 1000 * Nanosecond
	Milisecond        = 1000 * Microsecond
	Second            = 1000 * Milisecond
	Minute            = 60 * Second
	Hour              = 60 * Minute
	Day               = 24 * Hour
	Year              = 365 * Day
)

func HumanReadable(tm int64) string {
	var build []string
	durationsValues := []int64{Milisecond, Second, Minute, Hour, Day, Year}
	durationsNames := []string{"milisecond", "second", "minute", "hour", "day", "year"}

	for i := len(durationsValues) - 1; i >= 0; i-- {
		numOfThisUnit := tm / durationsValues[i]
		tm %= durationsValues[i]
		if numOfThisUnit != 0 {
			suffix := "s"
			if numOfThisUnit == 1 {
				suffix = ""
			}
			build = append(build, fmt.Sprintf("%d %s%s", numOfThisUnit, durationsNames[i], suffix))
		}
	}

	return strings.Join(build, " ")
}
