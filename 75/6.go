package leetcode75

import (
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	size := len(words)
	for i := 0; i < size/2; i++ {
		words[i], words[size-i-1] = words[size-i-1], words[i]
	}
	return strings.Join(words, " ")
}