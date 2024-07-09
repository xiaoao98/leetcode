package daily

func nodesBetweenCriticalPoints(head *ListNode) []int {
	if head == nil || head.Next == nil {
		return []int{-1, -1}
	}
	node := head.Next
	lastVal := head.Val
	lastCPos := 0
	firstCPos := 0
	minD := 100000
	maxD := -1
	pos := 1
	for node.Next != nil {
		if (node.Val > lastVal && node.Val > node.Next.Val) || (node.Val < lastVal && node.Val < node.Next.Val) {
			if lastCPos == 0 {
				lastCPos = pos
				firstCPos = pos
			} else {
				maxD = pos - firstCPos
				if pos-lastCPos < minD {
					minD = pos - lastCPos
				}
				lastCPos = pos
			}
		}
		pos++
		lastVal = node.Val
		node = node.Next
	}
	if minD == 100000 {
		minD = -1
	}
	return []int{minD, maxD}
}