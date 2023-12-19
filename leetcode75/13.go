package leetcode75

import "fmt"

func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	ret := 0
	for left < right {
		if height[left] > height[right] {
			if height[right]*(right-left) > ret {
				ret = height[right] * (right - left)
			}
			right--
		} else {
			if height[left]*(right-left) > ret {
				ret = height[left] * (right - left)
			}
			left++
		}
		fmt.Println(left, right, ret)
	}
	return ret
}