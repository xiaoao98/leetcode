package leetcode75

import (
	"container/heap"
	"sort"
)

func maxScore(nums1 []int, nums2 []int, k int) int64 {
    arr := [][]int{}
    for i := 0; i < len(nums1); i++ {
        arr = append(arr, []int{nums1[i], nums2[i]})
    }

    // O(NlogN)
    // Sort in decreasing order of nums2
    sort.Slice(arr, func (i, j int) bool {
        return arr[i][1] > arr[j][1]
    })

    minHeap := &MinHeap52{}
    sum := 0
    // O(KlogK)
    for i := 0; i < k; i++ {
        // O(logK)
        heap.Push(minHeap, arr[i])
        sum += arr[i][0]
    }

    res := sum * arr[k-1][1]

    // O(NlogK)
    for i := k; i < len(arr); i++ {
        // Remove element with smallest nums1
        // O(logK)
        smallNum := heap.Pop(minHeap).([]int)
        sum += arr[i][0] - smallNum[0]

        // O(logK)
        heap.Push(minHeap, arr[i])

        res = max(res, sum * arr[i][1])
    }

    return int64(res)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

type MinHeap52 [][]int

func (h MinHeap52) Len() int           { return len(h) }
func (h MinHeap52) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap52) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap52) Push(x interface{}) {
    *h = append(*h, x.([]int))
}

func (h *MinHeap52) Pop() interface{} {
    x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
// type IntArrayHeap struct{
// 	indices []int
// 	nums []int
// }

// func (h IntArrayHeap) Len() int           { return len(h.indices) }
// func (h IntArrayHeap) Less(i, j int) bool { return h.nums[i] < h.nums[j] }
// func (h IntArrayHeap) Swap(i, j int)      { h.indices[i], h.indices[j] = h.indices[j], h.indices[i] }
// func (h *IntArrayHeap) Push(x any)        { h.indices = append(h.indices, x.(int)) }
// func (h *IntArrayHeap) Pop() any {
// 	old := h.indices
// 	n := len(old)
// 	x := old[n-1]
// 	h.indices = old[0 : n-1]
// 	return x
// }
