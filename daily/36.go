package daily

func isValidSudoku(board [][]byte) bool {
	rowN := len(board)
	colN := len(board[0])
	var rowsCheck [9][9]bool
	var colsCheck [9][9]bool
	var boxesCheck [9][9]bool
	for i := 0; i < rowN; i++ {
		for j := 0; j < colN; j++ {
			if board[i][j] == '.' {
				continue
			}
			ele := int(board[i][j]-'0') - 1
			if rowsCheck[i][ele] {
				return false
			} else {
				rowsCheck[i][ele] = true
			}
			if colsCheck[j][ele] {
				return false
			} else {
				colsCheck[j][ele] = true
			}
			boxesIndex := i/3*3 + j/3
			if boxesCheck[boxesIndex][ele] {
				return false
			} else {
				boxesCheck[boxesIndex][ele] = true
			}
		}
	}
	return true
}