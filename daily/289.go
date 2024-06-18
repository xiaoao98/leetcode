package daily

func gameOfLife(board [][]int) {
	rowN := len(board)
	colN := len(board[0])
	for i := 0; i < rowN; i++ {
		for j := 0; j < colN; j++ {
			numberOne := getNumberOfOnes(i, j, rowN, colN, board)
			if board[i][j] == 0 {
				if numberOne == 3 {
					board[i][j] = 2
				}
			} else {
				if numberOne < 2 || numberOne > 3 {
					board[i][j] = -1
				}
			}
		}
	}
	for i := 0; i < rowN; i++ {
		for j := 0; j < colN; j++ {
			if board[i][j] == -1 {
				board[i][j] = 0
			} else if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
}

func getNumberOfOnes(i, j, row, col int, board [][]int) int {
	ret := 0
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, dirs := range directions {
		newi := i + dirs[0]
		newj := j + dirs[1]
		if newi >= 0 && newi < row && newj >= 0 && newj < col {
			if board[newi][newj] == 1 || board[newi][newj] == -1 {
				ret += 1
			}
		}
	}
	return ret
}