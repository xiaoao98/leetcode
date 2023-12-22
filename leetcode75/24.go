package leetcode75

import "fmt"

func equalPairs(grid [][]int) int {
	size := len(grid)
	rowsMap := make(map[string]int)
	colsArray := make([][]int, size)
	colsMap := make(map[string]int)
	for i := 0; i < size; i++ {
		rowString := fmt.Sprintf("%v", grid[i])
		rowsMap[rowString] += 1
		for j:=0; j < size; j++ {
			colsArray[j] = append(colsArray[j], grid[i][j])
		}
	}
	for i := 0; i < size; i++ {
		colString := fmt.Sprintf("%v", colsArray[i])
		colsMap[colString] += 1
	}
	ret := 0
	for key, val := range(rowsMap) {
		if val2, exists := colsMap[key]; exists {
			ret += val * val2
		}
	}
	return ret
}