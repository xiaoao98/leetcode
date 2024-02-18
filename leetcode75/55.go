package leetcode75

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	ret := make([]int, 0, len(spells))
	for _, spell := range spells {
		successNum := binSearch(potions, spell, success)
		ret = append(ret, successNum)
	}
	return ret
}

func binSearch(potions []int, spell int, success int64) int {
	low := 0
	size := len(potions)
	high := size - 1
	for  low <= high {
		mid := low + (high - low) / 2
		bool1 := int64(potions[mid] * spell) >= success
		bool2 := mid == 0 || int64(potions[mid-1] * spell) < success
		if bool1 && bool2 {
			return size - mid
		} else if int64(potions[mid] * spell) < success {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 0
}