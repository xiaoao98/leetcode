package leetcode75

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, k)
}

func quickSelect(nums []int, k int) int {
	size := len(nums)
	rand.Seed(time.Now().UnixNano())
	pivotIndex := rand.Intn(size)
	pivot := nums[pivotIndex]
	left := make([]int, 0)
	middle := make([]int, 0)
	right := make([]int, 0)
	for _, v := range nums {
		if v > pivot {
			left = append(left, v)
		} else if v < pivot {
			right = append(right, v)
		} else {
			middle = append(middle, v)
		}
	}
	if k <= len(left) {
		return quickSelect(left, k)
	}
	if k > len(left) + len(middle) {
		return quickSelect(right, k-len(left)-len(middle))
	}
	return pivot
}