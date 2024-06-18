package daily

import "math"

func scoreOfString(s string) int {
	ret := 0
	for i := 1; i < len(s); i++ {
		ret += int(math.Abs(float64(s[i]) - float64(s[i-1])))
	}
	return ret
}