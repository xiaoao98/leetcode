package leetcode75

func maxOperations(nums []int, k int) int {
	dic := make(map[int]int)
	ret := 0
	for _, num := range nums {
		if val, ok := dic[k-num]; ok && val != 0 {
			ret += 1
			dic[k-num]--
		} else {
			dic[num]++
		}
	}
	return ret
}