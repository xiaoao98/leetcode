package leetcode75

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	ret := 1
	right := points[0][1]
	for _, point := range points[1:] {
		if point[0] > right {
			ret += 1
			right = point[1]
		} else {
			if point[1] < right {
				right = point[1]
			}
		}
	}
	return ret
}