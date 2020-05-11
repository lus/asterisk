package utils

// StringArrayContains checks whether or not an array of strings contains the given element
func StringArrayContains(array []string, element string) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}
