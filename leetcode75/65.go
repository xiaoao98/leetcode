package leetcode75

func longestCommonSubsequence(text1 string, text2 string) int {
	size1 := len(text1)
	size2 := len(text2)
	dp := make([][]int, size1+1)
	for i := range dp {
		dp[i] = make([]int, size2+1)
	}
	for i, char1 := range text1 {
		for j, char2 := range text2 {
			if char1 != char2 {
				dp[i+1][j+1] = findMax(dp[i+1][j], dp[i][j+1])
			} else {
				dp[i+1][j+1] = dp[i][j] + 1
			}
		}
	}
	return dp[size1][size2]
}

func findMax(nums ...int) int {
	if len(nums) == 0 {
		// 如果没有提供任何参数，则返回0（或者根据实际情况返回一个错误）
		return 0
	}

	// 初始化max为第一个参数的值
	max := nums[0]

	// 遍历参数，找到最大的一个
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}