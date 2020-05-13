package utils

import (
	"strings"
)

// StringArrayContains checks whether or not an array of strings contains the given element
func StringArrayContains(array []string, element string) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}

// StringHasPrefix checks whether or not the given string starts with at least one element of the given array
func StringHasPrefix(str string, prefixes []string, replace bool) (string, bool) {
	for _, prefix := range prefixes {
		if strings.HasPrefix(str, prefix) {
			if replace {
				str = strings.TrimSpace(strings.Replace(str, prefix, "", 1))
			}
			return str, true
		}
	}
	return str, false
}

// StringHasSuffix checks whether or not the given string ends with at least one element of the given array
func StringHasSuffix(str string, suffixes []string, replace bool) (string, bool) {
	for _, suffix := range suffixes {
		if strings.HasSuffix(str, suffix) {
			if replace {
				str = strings.TrimSpace(strings.Replace(str, suffix, "", 1))
			}
			return str, true
		}
	}
	return str, false
}
