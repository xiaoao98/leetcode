package daily

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	//2, 4, 6, 8: 0, 2, 1, 5
	grumpyIndexs := make([]int, 0)
	ret := 0
	for i, gru := range grumpy {
		if gru == 1 {
			grumpyIndexs = append(grumpyIndexs, i)
		} else {
			ret += customers[i]
		}
	}
	toAdd := 0
	for i, index := range grumpyIndexs {
		end := index + minutes
		endIndex := binarySearchIndex(grumpyIndexs, i, end)
		canAdd := 0
		for j := i; j <= endIndex; j++ {
			canAdd += customers[grumpyIndexs[j]]
		}
		if canAdd > toAdd {
			toAdd = canAdd
		}
	}
	return ret + toAdd
}

func binarySearchIndex(arr []int, low int, target int) int {
	size := len(arr)
	high := size - 1
	ret := 0
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] <= target {
			ret = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ret
}