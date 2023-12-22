package leetcode75

func uniqueOccurrences(arr []int) bool {
	occurrenceMap := make(map[int]int, 0)
	occurrenceSet := make(map[int]bool, 0)
	for _, value := range arr {
		if _, exists := occurrenceMap[value]; exists {
			occurrenceMap[value] += 1
		} else {
			occurrenceMap[value] = 1
		}
	}
	for _, item := range occurrenceMap {
		if _, exists := occurrenceSet[item]; exists {
			return false
		} else {
			occurrenceSet[item] = true
		}
	}
	return true
}