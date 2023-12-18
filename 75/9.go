package leetcode75

import (
	"math"
)

func increasingTriplet(nums []int) bool {
    firstNum, secondNum := math.MaxInt32, math.MaxInt32
	for _, value := range(nums) {
		if value < firstNum && value < secondNum {
			firstNum = value
		} else if value > firstNum && value < secondNum {
			secondNum = value
		} else if value > firstNum && value >secondNum {
			return true
		}
	}
	return false
}