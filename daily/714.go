package daily

func maxProfit(prices []int, fee int) int {
	held := -prices[0]
	empty := 0
	for i := 1; i < len(prices); i++ {
		preHeld := held
		held = MaxInt(held, empty-prices[i])
		empty = MaxInt(empty, preHeld+prices[i]-fee)
	}
	return empty
}

func MaxInt(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}