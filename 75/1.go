package leetcode75

func mergeAlternately(word1 string, word2 string) string {
	merged := ""
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		merged += string(word1[i])
		merged += string(word2[j])
		i += 1
		j += 1
	}
	for i < len(word1) {
		merged += string(word1[i])
		i += 1
	}
	for j < len(word2) {
		merged += string(word2[j])
		j += 1
	}
	return merged
}