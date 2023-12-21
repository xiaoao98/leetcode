package leetcode75

func longestOnes(nums []int, k int) int {
	queue := make([]int, 0)
	ret := 0
	lengthNow := 0
	zerosFlipped := 0
	for key, value := range nums {
		if value == 1 {
			lengthNow += 1
		} else {
			if k == 0 {
				lengthNow = 0
				continue
			}
			if zerosFlipped < k {
				zerosFlipped += 1
				lengthNow += 1
			} else {
				lengthNow = key - queue[0]
				queue = queue[1:]
			}
			queue = append(queue, key)
		}
		if lengthNow > ret {
			ret = lengthNow
		}
	}
	return ret
}