package daily

func averageWaitingTime(customers [][]int) float64 {
	size := len(customers)
	totalWaitTime := 0
	lastTime := 0
	for _, customer := range customers {
		if customer[0] > lastTime {
			lastTime = customer[0]
		}
		totalWaitTime = lastTime + customer[1] - customer[0]
		lastTime += customer[1]

	}
	return float64(totalWaitTime) / float64(size)
}