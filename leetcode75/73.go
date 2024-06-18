package leetcode75

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ret := 0
	right := intervals[0][1]
	for _, interval := range intervals[1:] {
		if interval[0] < right {
			ret += 1
			if interval[1] < right {
				right = interval[1]
			}
		} else {
			right = interval[1]
		}
	}
	return ret
}
