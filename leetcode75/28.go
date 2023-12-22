package leetcode75

type RecentCounter struct {
	Counter []int
}

func Constructor() RecentCounter {
	counter := make([]int, 0)
	return RecentCounter{Counter: counter}
}

func (this *RecentCounter) Ping(t int) int {
	for this.Counter[0] < t-3000 {
		this.Counter = this.Counter[1:]
	}
	this.Counter = append(this.Counter, t)
	return len(this.Counter)
}