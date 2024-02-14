package leetcode75

import "container/heap"

func totalCost(costs []int, k int, candidates int) int64 {
	ret := 0
	size := len(costs)
	h := &IntArrayHeap{nums: costs}
	heap.Init(h)
    left := candidates
	right := size - candidates
	for i := 0; i < candidates; i++ {
		heap.Push(h, i)
	}
	for i := size-1; i >= max(left, right); i-- {
		heap.Push(h, i)
	}
	for i := 0; i < k; i++ {
		index := heap.Pop(h).(int)
		ret += costs[index]
		if left >= right {
			continue
		} else if index < left {
			heap.Push(h, left)
			left ++
		} else if index > left {
			heap.Push(h, right-1)
			right --
		}
	}
	return int64(ret)
}

type IntArrayHeap struct {
	indices []int
	nums    []int
}

func (h IntArrayHeap) Len() int { return len(h.indices) }
func (h IntArrayHeap) Less(i, j int) bool {
	if h.nums[h.indices[i]] != h.nums[h.indices[j]] {
		return h.nums[h.indices[i]] < h.nums[h.indices[j]]
	} else {
		return h.indices[i] < h.indices[j]
	}
}
func (h IntArrayHeap) Swap(i, j int) { h.indices[i], h.indices[j] = h.indices[j], h.indices[i] }
func (h *IntArrayHeap) Push(x any)   { h.indices = append(h.indices, x.(int)) }
func (h *IntArrayHeap) Pop() any {
	old := h.indices
	n := len(old)
	x := old[n-1]
	h.indices = old[0 : n-1]
	return x
}