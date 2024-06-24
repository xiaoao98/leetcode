package daily

func insert(intervals [][]int, newInterval []int) [][]int {
	size := len(intervals)
	start := newInterval[0]
	end := newInterval[1]
	var binarySearchIntervals func(int) int
	binarySearchIntervals = func(target int) int {
		low := 0
		high := size - 1
		for low <= high {
			mid := low + (high-low)/2
			if intervals[mid][0] <= target && (mid == size-1 || intervals[mid+1][0] > target) {
				return mid
			} else if intervals[mid][0] > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return -1
	}
	startPlace := binarySearchIntervals(start)
	endPlace := binarySearchIntervals(end)
	ret := make([][]int, 0)
	for i := 0; i < startPlace; i++ {
		ret = append(ret, intervals[i])
	}
	if startPlace >= 0 && start > intervals[startPlace][1] {
		ret = append(ret, intervals[startPlace])
	} else {
		start = intervals[startPlace][0]
	}
	if endPlace >= 0 && end < intervals[endPlace][1] {
		end = intervals[endPlace][1]
	}
	ret = append(ret, []int{start, end})
	for i := endPlace + 1; i < size; i++ {
		ret = append(ret, intervals[i])
	}
	return ret
}