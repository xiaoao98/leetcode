package leetcode75

func productExceptSelf(nums []int) []int {
	leftSum := make([]int, len(nums))
	rightSum := make([]int, len(nums))
	leftTmp, rightTmp := 1, 1
	for key, value := range nums {
		leftSum[key] = leftTmp * value
		leftTmp = leftSum[key]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		rightSum[i] = rightTmp * nums[i]
		rightTmp = rightSum[i]
	}
	answer := make([]int, len(nums))
	answer[0] = rightSum[1]
	answer[len(nums)-1] = leftSum[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		answer[i] = leftSum[i-1] * rightSum[i+1]
	}
	return answer
}