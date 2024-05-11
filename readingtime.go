/*
 go-readingtime - Estimate how long it takes to read a text

 https://github.com/edoardottt/go-readingtime

*/

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
	Version         = "0.0.4"
	wpm             = 238
	billion         = 1e9
	secondsInMinute = 60
	exceedMinute    = 30
)

// RawEstimate returns a float64 (seconds) representing an
// estimation of how long it would take to read the input text.
func RawEstimate(input string) float64 {
	words := strings.Fields(input)
	minutes := len(words) / wpm
	seconds := len(words) % wpm
	f := float64(minutes*secondsInMinute + durationLessThanAMinute(seconds))

	return f
}

// Estimate returns a time.Duration object representing an
// estimation of how long it would take to read the input text.
func Estimate(input string) time.Duration {
	f := RawEstimate(input)
	return duration(f)
}

// HumanEstimate returns a string representing a humanly readable
// estimation of how long it would take to read the input text.
func HumanEstimate(input string) string {
	var (
		rawMinutes = Estimate(input).Minutes()
		fmtResult  = "%s minute"
	)

	intResult := int(rawMinutes + 0.5)

	if intResult > 1 {
		fmtResult = "%s minutes"
	}

	return fmt.Sprintf(fmtResult, strconv.Itoa(intResult))
}

func duration(f float64) time.Duration {
	return time.Duration(f * billion)
}

func durationLessThanAMinute(words int) int {
	f := float64(words) * secondsInMinute / wpm
	return int(math.Ceil(f))
}
