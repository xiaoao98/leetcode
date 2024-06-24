package daily

func minKBitFlips(nums []int, k int) int {
	size := len(nums)
	flipped := make([]bool, size)
	currentFlips := 0
	totalFlips := 0
	for i := 0; i < size; i++ {
		if i >= k && flipped[i-k] == true {
			currentFlips--
		}
		if (nums[i] == 0 && currentFlips%2 == 0) || (nums[i] == 1 && currentFlips%2 == 1) {
			if i+k > size {
				return -1
			}
			flipped[i] = true
			currentFlips++
			totalFlips++
		}
	}
	return totalFlips
}