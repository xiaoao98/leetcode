package leetcode75

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxValue := 0
	for _, value := range candies {
		if value > maxValue {
			maxValue = value
		}
	}
	size := len(candies)
	output := make([]bool, size)
	for key, value := range candies {
		if value+extraCandies >= maxValue {
			output[key] = true
		}
	}
	return output
}