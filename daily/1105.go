package daily

func minHeightShelves(books [][]int, shelfWidth int) int {
	//mem[index][widthLeft]
	//recur(index, widthLeft, heightNow) sumHeight
	size := len(books)
	mem := make([][]int, size)
	for i := 0; i < size; i++ {
		mem[i] = make([]int, shelfWidth+1)
	}
	var recursive func(int, int, int) int
	recursive = func(index, widthLeft, heightNow int) int {
		width := books[index][0]
		height := books[index][1]
		if index == size-1 {
			if width <= widthLeft {
				if height < heightNow {
					return heightNow
				} else {
					return height
				}
			} else {
				return heightNow + height
			}
		}
		if mem[index][widthLeft] != 0 {
			return mem[index][widthLeft]
		}
		ret := -1
		if width <= widthLeft {
			// fmt.Println(widthLeft)
			// fmt.Println(width)
			heightNew := heightNow
			if heightNew < height {
				heightNew = height
			}
			ret = recursive(index+1, widthLeft-width, heightNew)
		}
		nextRet := heightNow + recursive(index+1, shelfWidth-width, height)
		if ret == -1 || ret > nextRet {
			ret = nextRet
		}
		mem[index][widthLeft] = ret
		return ret
	}
	return recursive(0, shelfWidth, 0)
}