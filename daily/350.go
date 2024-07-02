package daily

func intersect(nums1 []int, nums2 []int) []int {
	counter := make(map[int]int, 0)
	for _, val := range nums1 {
		if _, exists := counter[val]; !exists {
			counter[val] = 0
		}
		counter[val]++
	}
	intersects := make([]int, 0)
	for _, num := range nums2 {
		if counter[num] > 0 {
			intersects = append(intersects, num)
			counter[num]--
		}
	}
	return intersects
}