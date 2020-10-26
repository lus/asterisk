package utils

import (
	"math/rand"
	"regexp"
	"strconv"

	"github.com/Lukaesebrot/asterisk/static"
)

// IntervalRegex defines the compiled interval RegEx
var IntervalRegex = regexp.MustCompile(static.IntervalRegexString)

// GenerateFromInterval generates a random number using the given interval
func GenerateFromInterval(raw string) (bool, int) {
	// Check if the raw string is a valid interval
	if !IntervalRegex.MatchString(raw) {
		return false, 0
	}

	// Define the submatches
	submatches := IntervalRegex.FindStringSubmatch(raw)
	leftChar := submatches[1]
	rightChar := submatches[4]
	leftNumber, _ := strconv.Atoi(submatches[2])
	if leftChar == "(" {
		leftNumber = leftNumber + 1
	}
	rightNumber, _ := strconv.Atoi(submatches[3])
	if rightChar == ")" {
		rightNumber = rightNumber - 1
	}

	// Validate the given numbers
	if rightNumber < leftNumber {
		return false, 0
	}
	if leftNumber == rightNumber {
		return true, leftNumber
	}

	// Return the random-generated number
	return true, rand.Intn(rightNumber-leftNumber+1) + leftNumber
}
