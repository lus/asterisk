package utils

// PrettifyBool prettifies a boolean
func PrettifyBool(value bool) string {
	if value {
		return "enabled"
	}
	return "disabled"
}
