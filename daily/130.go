package daily

func solve(board [][]byte) {
	row := len(board)
	col := len(board[0])
	var isSurronded func(i, j int) bool
	isSurronded = func(i, j int) bool {
		if i <= 0 || i >= row-1 || j <= 0 || j >= col-1 {
			return false
		}
		board[i][j] = 'W'
		flag := true
		adj := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range adj {
			newi := i + dir[0]
			newj := j + dir[1]
			if board[newi][newj] == 'O' {
				if !isSurronded(newi, newj) {
					flag = false
				}
			}
		}
		return flag
	}
	changeBoard := func(changeFrom, changeTo byte) {
		for i := 1; i < row; i++ {
			for j := 1; j < col; j++ {
				if board[i][j] == changeFrom {
					board[i][j] = changeTo
				}
			}
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if board[i][j] == 'O' {
				if isSurronded(i, j) {
					changeBoard('W', 'X')
				} else {
					changeBoard('W', 'P')
				}
			}
		}
	}
	changeBoard('P', 'O')
}