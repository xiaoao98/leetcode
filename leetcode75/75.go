package leetcode75

import "container/heap"

func dailyTemperatures(temperatures []int) []int {//do not need to use heap, stack is enough
	ret := make([]int, len(temperatures))
	h := &IntHeap{}
	heap.Init(h)
	dic := make(map[int][]int, 0)
	for i, temperature := range temperatures {
		dic[temperature] = append(dic[temperature], i)
		heap.Push(h, temperature)
		for h.Len() > 0 && (*h)[0] < temperature {
			val := heap.Pop(h).(int)
			for _, ele := range dic[val] {
				ret[ele] = i - ele
			}
			dic[val] = []int{}
		}
	}
	return ret
}
