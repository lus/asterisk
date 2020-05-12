package utils

import (
	"math/rand"
	"strconv"
	"strings"
)

// FormatInterval generates a random number using the given interval
func FormatInterval(raw string) (bool, int) {
	// Format the raw string
	raw = strings.ReplaceAll(raw, " ", "")
	split := strings.Split(raw, ",")

	// Validate the raw string
	if len(split) != 2 {
		return false, 0
	}
	_, prefixValid := StringHasPrefix(raw, []string{"(", "["}, false)
	_, suffixValid := StringHasSuffix(raw, []string{")", "]"}, false)
	if !prefixValid || !suffixValid {
		return false, 0
	}

	// Define the pre- and suffix
	prefix := raw[0]
	suffix := raw[len(raw)-1]

	// Define the two numbers
	rawNum1 := strings.Replace(strings.Replace(split[0], "(", "", 1), "[", "", 1)
	rawNum2 := strings.Replace(strings.Replace(split[1], ")", "", 1), "]", "", 1)
	num1 := 0
	num2 := 0
	if rawNum1 != "" {
		parsedNum1, err := strconv.Atoi(rawNum1)
		if err != nil || parsedNum1 < 0 {
			return false, 0
		}
		num1 = parsedNum1
	}
	if rawNum2 != "" {
		parsedNum2, err := strconv.Atoi(rawNum2)
		if err != nil || parsedNum2 < 0 {
			return false, 0
		}
		num2 = parsedNum2
	}

	// Validate the two numbers
	if (num1 > 0) && (num1 == num2) {
		return false, 0
	}
	if (num1 > 0) && num1 > num2 {
		return false, 0
	}
	if (prefix == '(' && suffix == ')') && num1+1 == num2 {
		return false, 0
	}

	// Generate the random number
	if num1+num2 == 0 {
		return true, rand.Int()
	}
	min := num1
	if prefix == '(' {
		min = min + 1
	}
	max := num2
	if suffix == ')' {
		max = max - 1
	}
	return true, rand.Intn(max-min+1) + min
}
