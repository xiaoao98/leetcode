package leetcode75

func minReorder(n int, connections [][]int) int {
	ret := 0
	var dfs46 func(i int)
	dfs46 = func(i int) {
		for _, connection := range connections {
			if connection[0] == i {
				dfs46(connection[1])
			} else if connection[1] == i {
				ret += 1
				dfs46(connection[0])
			}
		}
	}
	dfs46(0)
	return ret
}