package daily

//easier way is to see a common node in the first and second node
func findCenter(edges [][]int) int {
    counter := make(map[int]int, 0)
	for _, edge := range edges {
		if _, exists := counter[edge[0]]; !exists {
			counter[edge[0]] = 0
		}
		if _, exists := counter[edge[1]]; !exists {
			counter[edge[1]] = 0
		}
		counter[edge[0]] ++
		counter[edge[1]] ++
	}
	size := len(counter)
	for key, val := range counter {
		if val == size-1 {
			return key
		}
	}
	return -1
}