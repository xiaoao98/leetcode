package daily

import "sort"

type DiffAndPro struct {
	difficulty int
	profit     int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	size := len(difficulty)
	diffAndPros := make([]DiffAndPro, 0)
	for i := 0; i < size; i++ {
		diffAndPros = append(diffAndPros, DiffAndPro{difficulty: difficulty[i], profit: profit[i]})
	}
	sort.Slice(diffAndPros, func(i, j int) bool{
		return diffAndPros[i].difficulty < diffAndPros[j].difficulty
	})
	filteredDiffAndPros := make([]DiffAndPro, 0)
	filteredDiffAndPros = append(filteredDiffAndPros, diffAndPros[0])
	lastProfit := diffAndPros[0].profit
	for i := 1; i < size; i++ {
		if diffAndPros[i].profit > lastProfit {
			filteredDiffAndPros = append(filteredDiffAndPros, diffAndPros[i])
			lastProfit = diffAndPros[i].profit
		}
	}
	ret := 0
	for _, val := range worker {
		ret += binarySearchProfits(filteredDiffAndPros, val)
	}
	return ret
}

func binarySearchProfits(filteredDiffAndPros []DiffAndPro, target int) int{
	size := len(filteredDiffAndPros)
	low := 0
	high := size-1
	for low <= high {
		mid := low + (high - low)/2
		if filteredDiffAndPros[mid].difficulty <= target && (mid >= size -1 || filteredDiffAndPros[mid+1].difficulty > target) {
			return filteredDiffAndPros[mid].profit
		} else if filteredDiffAndPros[mid].difficulty > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0
}