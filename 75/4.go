package leetcode75

func canPlaceFlowers(flowerbed []int, n int) bool {
	for key, value := range flowerbed {
		if value == 0 {
			emptyLeft := key == 0 || flowerbed[key-1] == 0
			emptyRight := key == len(flowerbed)-1 || flowerbed[key+1] == 0
			if (emptyLeft) && (emptyRight) {
				n -= 1
				flowerbed[key] = 1
				if n <= 0 {
					return true
				}
			}
		}
	}
	return n <= 0
}