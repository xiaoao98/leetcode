package leetcode75

func maxProfit(prices []int, fee int) int {
	size := len(prices)
	hold := make([]int, size)
	free := make([]int, size)
	hold[0] = -prices[0]
	free[0] = 0
	for i := 1; i < size; i++ {
		hold[i] = MaxInt(hold[i-1], free[i-1]-prices[i])
		free[i] = MaxInt(free[i-1], hold[i-1]+prices[i]-fee)
	}
	return free[size-1]
}