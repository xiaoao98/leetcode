package daily

func heightChecker(heights []int) int {
	heightsCopy := make([]int, len(heights))
	for i, ele := range heights {
		heightsCopy[i] = ele
	}
	heightsCopy = countSort(heightsCopy)
	ret := 0
	for i, ele := range heights {
		if heightsCopy[i] != ele {
			ret += 1
		}
	}
	return ret
}

func countSort(arr []int) []int {
	minV := 101
	maxV := 0
	countMap := make(map[int]int, 0)
	for _, ele := range arr {
		if ele < minV {
			minV = ele
		}
		if ele > maxV {
			maxV = ele
		}
		if _, exists := countMap[ele]; exists {
			countMap[ele] += 1
		} else {
			countMap[ele] = 1
		}
	}
	index := 0
	for i := minV; i <= maxV; i++ {
		val, exists := countMap[i]
		for exists && val > 0 {
			arr[index] = i
			index++
			countMap[i]--
			val, exists = countMap[i]
		}
	}
	return arr
}