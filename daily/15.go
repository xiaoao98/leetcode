package daily

import "sort"

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	var twoSums func(int, int) [][]int
	twoSums = func(start int, sumAdded int) [][]int {
		left := start
		right := len(nums) - 1
		ret := make([][]int, 0)
		for left < right {
			if nums[left] + nums[right] == sumAdded {
				ret = append(ret, []int{nums[left], nums[right]})
				left ++
                for left < right && nums[left] == nums[left-1] {
                    left ++
                }
			} else if nums[left] + nums[right] < sumAdded {
				left ++
			} else {
				right --
			}
		}
		return ret
	}
	for i := 0; i < len(nums); i ++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		twoRet := twoSums(i+1, -nums[i])
		for _, ele := range twoRet {
			newEle := make([]int, 0)
			newEle = append(newEle, nums[i])
			newEle = append(newEle, ele...)
			ret = append(ret, newEle)
		}
	}
	return ret
}