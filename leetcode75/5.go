package leetcode75

func reverseVowels(s string) string {
	vowelSet := make(map[string]bool)
	vowelArray := [10]string{"a", "i", "e", "o", "u", "A", "O", "E", "I", "U"}
	for _, value := range vowelArray {
		vowelSet[value] = true
	}
	left, right := 0, len(s)-1
	output := []rune(s)
	for left < right {
		if _, leftIsVowel := vowelSet[(string(s[left]))]; !leftIsVowel {
			left++
			continue
		}
		if _, rightIsVowel := vowelSet[(string(s[right]))]; !rightIsVowel {
			right--
			continue
		}
		output[left], output[right] = output[right], output[left]
		left++
		right--
	}
	return string(output)
}