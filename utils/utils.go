package utils

import "strings"

// TrimLeftRightSpace trim left and rgiht space
func TrimLeftRightSpace(str string) string {
	str = strings.TrimRight(str, "\t \n")
	str = strings.TrimLeft(str, "\t \n")
	return str
}
