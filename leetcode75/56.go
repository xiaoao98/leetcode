package leetcode75

func findPeakElement(nums []int) int {
	size := len(nums)
	low := 0
	high := size - 1
	for low <= high {
		mid := low + (high-low)/2
		if (mid == 0 || nums[mid] > nums[mid-1]) && (mid == size-1 || nums[mid] > nums[mid+1]) {
			return mid
		} else if nums[mid] > nums[mid+1] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}