package leetcode75

func findCircleNum(isConnected [][]int) int {
	size := len(isConnected)
	connected := make([]int, len(isConnected))
	ret := 0
	var dfs45 func(i int)
	dfs45 = func(i int) {
		for key := 0; key < size; key++ {
			ele := isConnected[i][key]
			if ele != 0 && connected[key] == 0 {
				connected[key] = 1
				dfs45(key)
			}
		}
	}

	for i := 0; i < size; i++ {
		if connected[i] == 0 {
			ret += 1
			connected[i] = 1
			dfs45(i)
		}
	}
	return ret
}