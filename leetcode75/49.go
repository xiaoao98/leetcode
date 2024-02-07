package leetcode75

func orangesRotting(grid [][]int) int {
	goodNum := 0
	row := len(grid)
	col := len(grid[0])
	queue := make([][2]int, 0)
	ret := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				goodNum++
			} else if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	for len(queue) > 0 {
		if goodNum == 0 {
			break
		}
		for _, v := range queue {
			x := v[0]
			y := v[1]
			for _, dir := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				xNew := x + dir[0]
				yNew := y + dir[1]
				if xNew >= 0 && xNew <= row-1 && yNew >= 0 && yNew <= col-1 && grid[xNew][yNew] == 1 {
					goodNum -= 1
					grid[xNew][yNew] = 2
					queue = append(queue, [2]int{xNew, yNew})
				}
			}
			queue = queue[1:]
		}
		ret += 1
	}
	if goodNum == 0 {
		return ret
	} else {
		return -1
	}
}