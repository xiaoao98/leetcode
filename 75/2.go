package leetcode75

import (
	"strings"
)
func gcdOfStrings(str1 string, str2 string) string {
    shorter := ""
	if len(str1) > len(str2) {
		shorter = str2
	} else {
		shorter = str1
	}
	i := len(shorter)
	for i >= 0 {
		replaced := shorter[:i]
		if len(strings.Replace(str1, replaced, "", -1)) == 0 && len(strings.Replace(str2, replaced, "", -1)) == 0 {
			return replaced
		}
		i -= 1
	}
	return ""
}