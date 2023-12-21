package leetcode75

func pivotIndex(nums []int) int {
	leftSum := make([]int, 0)
	rightSum := make([]int, 0)
	leftSum = append(leftSum, 0)
	rightSum = append(rightSum, 0)
	leftNow := 0
	rightNow := 0
	for i := 0; i < len(nums); i++ {
		leftNow += nums[i]
		leftSum = append(leftSum, leftNow)
	}
	for i := len(nums) - 1; i >= 0; i-- {
		rightNow += nums[i]
		rightSum = append(rightSum, rightNow)
	}
	for i := 0; i < len(nums); i++ {
		if leftSum[i] == rightSum[len(nums)-i-1] {
			return i
		}
	}
	return -1
}