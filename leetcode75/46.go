package leetcode75

func minReorder(n int, connections [][]int) int {
	ret := 0
	var dfs46 func(i int)
	seen := make([]bool, n)
	dic1 := make([][]int, n)
	dic2 := make([][]int, n)
	for _, connection := range connections {
		former := connection[0]
		latter := connection[1]
		dic1[former] = append(dic1[former], latter)
		dic2[latter] = append(dic2[latter], former)
	}
	dfs46 = func(i int) {
		for _, ele := range dic1[i] {
			if !seen[ele] {
				ret += 1
				seen[ele] = true
				dfs46(ele)
			}
		}
		for _, ele := range dic2[i] {
			if !seen[ele] {
				seen[ele] = true
				dfs46(ele)
			}
		}
	}
	seen[0] = true
	dfs46(0)
	return ret
}