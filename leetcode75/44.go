package leetcode75

func canVisitAllRooms(rooms [][]int) bool {
	size := len(rooms)
	roomsLocked := make([]int, size)
	roomsLocked[0] = 1
	queue := rooms[0]
	for len(queue) > 0 {
		size2 := len(queue)
		for i := 0; i < size2; i++ {
			if roomsLocked[queue[i]] == 0 {
				queue = append(queue, rooms[queue[i]]...)
				roomsLocked[queue[i]] = 1
			}
		}
		queue = queue[size2:]
	}
	for _, value := range roomsLocked {
		if value == 0 {
			return false
		}
	}
	return true
}