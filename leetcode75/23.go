package leetcode75

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	// check if the elements are same
	word1Freq := getFreq(word1)
	word2Freq := getFreq(word2)
	if len(word1Freq) != len(word2Freq) {
		return false
	}
	for key1 := range word1Freq {
		if _, ok := word2Freq[key1]; !ok {
			return false
		}
	}
	//check if freq are same
	word1FreqMap := getFreqMap(word1Freq)
	word2FreqMap := getFreqMap(word2Freq)
	for key1, val1 := range word1FreqMap {
		if val2, ok := word2FreqMap[key1]; !ok || val2 != val1 {
			return false
		}
	}
	return true
}

func getFreq(s string) map[byte]int {
	charFreq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		char := s[i]
		if _, ok := charFreq[char]; ok {
			charFreq[char]++
		} else {
			charFreq[char] = 1
		}
	}
	return charFreq
}
func getFreqMap(charFreq map[byte]int) map[int]int {
	freqMap := make(map[int]int)
	for _, item := range charFreq {
		if _, ok := freqMap[item]; ok {
			freqMap[item]++
		} else {
			freqMap[item] = 1
		}
	}
	return freqMap
}