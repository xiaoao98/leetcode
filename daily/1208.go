package daily

import "math"

func equalSubstring(s string, t string, maxCost int) int {
	ret := 0
	size := len(s)
	queue := make([]int, 0)
	costNow := 0
	for i := 0; i < size; i++ {
		cost := int(math.Abs(float64(int(s[i]) - int(t[i]))))
		costNow += cost
        // fmt.Println(costNow)
		queue = append(queue, cost)
		if costNow <= maxCost {
			if len(queue) > ret {
				ret = len(queue)
			}
		} else {
			for len(queue) > 0 && costNow > maxCost {
				costNow -= queue[0]
				queue = queue[1:]
			}
		}
        // fmt.Println(queue)
	}
	return ret
}
