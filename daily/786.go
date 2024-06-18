package daily

import "container/heap"

type pair struct {
	fraction []int
	value    float64
}

type priorityQueue []pair

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].value < pq[j].value }
func (pq priorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x any)        { *pq = append(*pq, x.(pair)) }
func (pq *priorityQueue) Pop() any {
	x := (*pq)[len(*pq)-1]
	(*pq) = (*pq)[:len(*pq)-1]

	return x
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	h := &priorityQueue{}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			heap.Push(h, pair{
				[]int{arr[i], arr[j]},
				float64(arr[i]) / float64(arr[j]),
			})
		}
	}

	for k > 1 {
		heap.Pop(h)
		k--
	}

	return heap.Pop(h).(pair).fraction
}