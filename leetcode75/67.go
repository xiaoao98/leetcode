package leetcode75

func minDistance(word1 string, word2 string) int {
	size1 := len(word1)
	size2 := len(word2)
	dp := make([][]int, size1+1)
	for i := range dp {
		dp[i] = make([]int, size2+1)
		dp[i][0] = i
	}
	for i := 0; i < size2+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < size1+1; i++ {
		for j := 1; j < size2+1; j++ {
			candicadate := dp[i-1][j-1] + 1
			if word1[i-1] == word2[j-1] {
				candicadate = dp[i-1][j-1]
			}
			dp[i][j] = findMin(candicadate, dp[i][j-1]+1, dp[i-1][j]+1)
		}
	}
	return dp[size1][size2]
}

func findMin(nums ...int) int {
	if len(nums) == 0 {
		// 如果没有提供任何参数，则返回0（或者根据实际情况返回一个错误）
		return 0
	}

	// 初始化max为第一个参数的值
	min := nums[0]

	// 遍历参数，找到最大的一个
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}