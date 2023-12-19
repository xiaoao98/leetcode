package leetcode75

import (
	"strconv"
)

func Compress(chars []byte) int {
	last := ""
	length := 0
	pos := 0
	for _, char := range chars {
		if string(char) != last {
			if length > 1 {
				lengthStr := strconv.Itoa(length)
				for i := 0; i < len(lengthStr); i++ {
					chars[pos] = lengthStr[i]
					pos++
				}
			}
			chars[pos] = char
			length = 1
			pos++
		} else {
			length++
		}
		last = string(char)
	}
	if length > 1 {
		lengthStr := strconv.Itoa(length)
		for i := 0; i < len(lengthStr); i++ {
			chars[pos] = lengthStr[i]
			pos++
		}
	}
	return pos
}

