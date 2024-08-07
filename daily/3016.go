package daily

import "sort"

func minimumPushes(word string) int {
	counter := make(map[int]int, 0)
	for i := 0; i < len(word); i++ {
		c := word[i]
		cNum := int(c) - int('a')
		if _, exists := counter[cNum]; !exists {
			counter[cNum] = 0
		}
		counter[cNum]++
	}
	counterVal := make([]int, 0)
	for _, val := range counter {
		counterVal = append(counterVal, val)
	}
	sort.Ints(counterVal)
	ret := 0
	num := 8
	pushNum := 1
	for i := len(counterVal)-1; i >= 0; i -- {
		ret += counterVal[i] * pushNum
		num -- 
		if num == 0 {
			pushNum ++
			num = 8
		}
	}
	return ret
}