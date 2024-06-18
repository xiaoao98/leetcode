package daily

import (
	"fmt"
	"sort"
)

func findRelativeRanks(score []int) []string {
	scoreCopy := make([]int, len(score))
	copy(scoreCopy, score)
	sort.Ints(scoreCopy)
	scoreMap := make(map[int]int, 0)
	for i, sc := range scoreCopy {
		scoreMap[sc] = len(scoreCopy) - i
	}
	ret := make([]string, 0)
	for _, s := range score {
		switch scoreMap[s] {
		case 1:
			ret = append(ret, "Gold medal")
		case 2:
			ret = append(ret, "Silver medal")
		case 3:
			ret = append(ret, "Bronze medal")
		default:
			ret = append(ret, fmt.Sprint(scoreMap[s]))
		}
	}
	return ret
}