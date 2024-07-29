package daily

func numTeams(rating []int) int {
	size := len(rating)
	increasingTeam := make([][3]int, size)
	decreasingTeam := make([][3]int, size)
	var increaseRecur func(int, int) int
	increaseRecur = func(index int, loc int) int {
		if index == size {
			return 0
		}
		if loc == 2 {
			return 1
		}
		tmp := 0
		if increasingTeam[index][loc] != 0 {
			return increasingTeam[index][loc]
		}
		for i := index + 1; i < size; i++ {
			if rating[i] > rating[index] {
				tmp += increaseRecur(i, loc+1)
			}
		}
		increasingTeam[index][loc] = tmp
		return tmp
	}
	var decreaseRecur func(int, int) int
	decreaseRecur = func(index int, loc int) int {
		if index == size {
			return 0
		}
		if loc == 2 {
			return 1
		}
		tmp := 0
		if decreasingTeam[index][loc] != 0 {
			return decreasingTeam[index][loc]
		}
		for i := index + 1; i < size; i++ {
			if rating[i] < rating[index] {
				tmp += decreaseRecur(i, loc+1)
			}
		}
		decreasingTeam[index][loc] = tmp
		return tmp
	}
	ret := 0
	for i := 0; i < size; i++ {
		ret += increaseRecur(i, 0)
		ret += decreaseRecur(i, 0)
	}
	return ret
}