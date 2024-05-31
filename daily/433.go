package daily

func MinMutation(startGene string, endGene string, bank []string) int {
	//1. bank to map 2. startgene change 3*len(startGene), =endgene end, else added into queue(first time show)
	ret := 0
	bankMap := make(map[string]bool, 0)
	visited := make(map[string]bool, 0)
	for _, ele := range bank {
		bankMap[ele] = true
	}
	queue := make([]string, 0)
	queue = append(queue, startGene)
	changeMap := make(map[byte][]byte)
	changeMap['A'] = []byte{'T', 'C', 'G'}
	changeMap['T'] = []byte{'A', 'C', 'G'}
	changeMap['C'] = []byte{'T', 'A', 'G'}
	changeMap['G'] = []byte{'T', 'C', 'A'}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			ele := queue[i]
			if ele == endGene {
				return ret
			}
			eleArr := []byte(ele)
			for j := 0; j < len(eleArr); j++ {
				toChange := changeMap[eleArr[j]]
				for _, changed := range toChange {
					newEle := make([]byte, 0)
					newEle = append(newEle, eleArr[:j]...)
					newEle = append(newEle, changed)
					newEle = append(newEle, eleArr[j+1:]...)
					newEleString := string(newEle)
					if bankMap[newEleString] && !visited[newEleString] {
						queue = append(queue, newEleString)
						visited[newEleString] = true
					}
				}
			}
			// fmt.Println(queue)
		}
		queue = queue[size:]
		ret += 1
	}
	return -1
}