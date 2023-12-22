package leetcode75

func findDifference(nums1 []int, nums2 []int) [][]int {
	ret := make([][]int, 2)
	nums1Set := make(map[int]bool, 0)
	nums2Set := make(map[int]bool, 0)
	for i := 0; i < len(nums1); i++ {
		nums1Set[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		nums2Set[nums2[i]] = true
	}
	for item := range nums1Set {
		if !nums2Set[item] {
			ret[0] = append(ret[0], item)
		}
	}
	for item := range nums2Set {
		if !nums1Set[item] {
			ret[1] = append(ret[1], item)
		}
	}
	return ret
}