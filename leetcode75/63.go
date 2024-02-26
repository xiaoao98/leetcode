package leetcode75

import "math"

func numTilings(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	MOD := int(math.Pow(10, 9) + 7)
	dp1 := make([]int, n)
	dp2 := make([]int, n)
	dp1[0] = 1
	dp1[1] = 2
	dp1[2] = 5
	dp2[0] = 2
	dp2[1] = 4
	dp2[2] = 8
	for i := 3; i < n; i++ {
		dp1[i] = (dp1[i-1] + dp1[i-2] + dp2[i-2]) % MOD
		dp2[i] = (dp1[i-1]*2 + dp2[i-1]) % MOD
	}
	return dp1[n-1]
}