package readingtime

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// https://www.sciencedirect.com/science/article/abs/pii/S0749596X19300786.
const (
	wpm             = 238
	billion         = 1e9
	secondsInMinute = 60
	exceedMinute    = 30
)

// Estimate returns a string represting a humanly readable
// estimation of how long it would take to read the input text.
func Estimate(input string) string {
	var (
		rawMinutes = RawEstimate(input).Minutes()
		fmtResult  = "%s minute"
	)

	intResult := int(rawMinutes + 0.5)

	if intResult > 1 {
		fmtResult = "%s minutes"
	}

	return fmt.Sprintf(fmtResult, strconv.Itoa(intResult))
}

// RawEstimate returns a time.Time object represting an
// estimation of how long it would take to read the input text.
func RawEstimate(input string) time.Duration {
	words := strings.Fields(input)
	minutes := len(words) / wpm
	seconds := len(words) % wpm
	f := float64(minutes*secondsInMinute + durationLessThanAMinute(seconds))

	return duration(f)
}

func duration(f float64) time.Duration {
	return time.Duration(f * billion)
}

func durationLessThanAMinute(words int) int {
	f := float64(words) * secondsInMinute / wpm
	return int(math.Ceil(f))
}
