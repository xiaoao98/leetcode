package daily

func threeConsecutiveOdds(arr []int) bool {
	counter := 0
	for _, val := range arr {
		if val%2 == 1 {
			counter += 1
			if counter == 3 {
				return true
			}
		} else {
			counter = 0
		}
	}
	return false
}