package leetcode75

func singleNumber(nums []int) int {
	output := 0
	for _, num := range nums {
		output ^= num
	}
	return output
}