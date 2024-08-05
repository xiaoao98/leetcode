package daily

func kthDistinct(arr []string, k int) string {
	dic := make(map[string]int, 0)
	for _, val := range arr {
		if _, exists := dic[val]; !exists {
			dic[val] = 0
		}
		dic[val]++
	}
	for _, val := range arr {
		if dic[val] <= 1 {
			k--
			if k == 0 {
				return val
			}
		}
	}
	return ""
}