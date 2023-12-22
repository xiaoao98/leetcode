package leetcode75

import (
	"unicode"
)

func DecodeString(s string) string {
	var stack []rune
	for _, char := range s {
		if string(char) != "]" {
			stack = append(stack, char)
		} else {
			var repeated []rune
			for string(stack[len(stack)-1]) != "[" {
				repeated = append(repeated, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			repeatedTimes := 0
			digit := 1
			for len(stack) > 0 && unicode.IsDigit(stack[len(stack)-1]) {
				repeatedTimes = int(stack[len(stack)-1]-'0') * digit + repeatedTimes
				stack = stack[:len(stack)-1]
				digit *= 10
			}
			// fmt.Print(repeatedTimes)
			// fmt.Print(repeated)
			for i := 0; i < repeatedTimes; i++ {
				for j := len(repeated) - 1; j >= 0; j-- {
					stack = append(stack, repeated[j])
				}
			}
			// fmt.Print(string(stack))
		}
	}
	return string(stack)
}