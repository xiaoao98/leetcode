package leetcode75

func rob(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}
	dp := make([]int, size)
	dp[0] = nums[0]
	dp[1] = MaxInt(nums[0], nums[1])
	for i := 2; i < size; i++ {
		dp[i] = MaxInt(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[size-1]
}

func MaxInt(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}