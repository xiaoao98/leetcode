package leetcode75

func nearestExit(maze [][]byte, entrance []int) int {
	rowNum := len(maze)
	colNum := len(maze[0])
	queue := make([][]int, 0)
	queue = append(queue, entrance)
	step := 0
	div := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			locx := queue[i][0]
			locy := queue[i][1]
			if locx == 0 || locx == rowNum-1 || locy == 0 || locy == colNum-1 {
				if locx != entrance[0] || locy != entrance[1] {
					return step
				}
			}
			for _, v := range div {
				newLocx := locx + v[0]
				newLocy := locy + v[1]
				if newLocx >= 0 && newLocx <= rowNum-1 && newLocy >= 0 && newLocy <= colNum-1 && maze[newLocx][newLocy] == '.' {
					maze[newLocx][newLocy] = '+'
					queue = append(queue, []int{newLocx, newLocy})
				}
			}
		}
		queue = queue[size:]
		step += 1
	}
	return -1
}