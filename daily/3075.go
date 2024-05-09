package daily

import (
	"math/rand"
	"sort"
)

func maximumHappinessSum(happiness []int, k int) int64 {
	rightPart := make([]int, 0)
	var quickselect func([]int, int) []int
	quickselect = func(arr []int, k int) []int {
		size := len(arr)
		if size == 1 {
			return arr
		}
		selectedLoc := rand.Intn(size)
		selected := arr[selectedLoc]
		leftArr := make([]int, 0)
		midArr := make([]int, 0)
		rightArr := make([]int, 0)
		for i := 0; i < size; i ++ {
			if arr[i] < selected {
				leftArr = append(leftArr, arr[i])
			} else if arr[i] == selected {
				midArr = append(midArr, arr[i])
			} else {
				rightArr = append(rightArr, arr[i])
			}
		}
		if len(rightArr) > k {
			return quickselect(rightArr, k)
		} else if len(leftArr) + k > size {
			rightPart = append(rightPart, midArr...)
			rightPart = append(rightPart, rightArr...)
			return quickselect(leftArr, len(leftArr)-size+k)
		} else {
			ret := make([]int, 0)
			for i:=0; i < k - len(rightArr); i++ {
				ret = append(ret, selected)
			}
			ret = append(ret, rightArr...)
			return ret
		}
	}
	tmp := quickselect(happiness, k)
	tmp = append(tmp, rightPart...)
	sort.Ints(tmp)
	sum := 0
	for i, ele := range tmp {
		if ele >= k-i {
			sum += ele-k+i+1
		}
	}
	return int64(sum)
}

