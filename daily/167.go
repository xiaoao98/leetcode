package daily

func twoSum(numbers []int, target int) []int {
	size := len(numbers)
	low := 0
	high := size - 1
	for low < high {
		if numbers[low]+numbers[high] == target {
			return []int{low + 1, high + 1}
		} else if numbers[low]+numbers[high] < target {
			low += 1
		} else {
			high -= 1
		}
	}
	return []int{-1, -1}
}