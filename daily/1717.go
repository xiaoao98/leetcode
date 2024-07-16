package daily

func maximumGain(s string, x int, y int) int { //be careful that len(stack) > 0
	//remove higher score pair first
	var higherScoreChar string
	var lowerScoreChar string
	var higherScore int
	var lowerScore int
	if x >= y {
		higherScoreChar = "ab"
		lowerScoreChar = "ba"
		higherScore = x
		lowerScore = y
	} else {
		higherScoreChar = "ba"
		lowerScoreChar = "ab"
		higherScore = y
		lowerScore = x
	}
	score := 0
	removeSubstring := func(s string, sub string, addScore int) string {
		charArr := []byte(s)
		stack := []byte{}
		for i := 0; i < len(s); i++ {
			if charArr[i] == sub[1] && len(stack) > 0 && stack[len(stack)-1] == sub[0] {
				score += addScore
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, charArr[i])
			}
		}
		return string(stack)
	}
	stringNew := removeSubstring(s, higherScoreChar, higherScore)
	_ = removeSubstring(stringNew, lowerScoreChar, lowerScore)
	return score
}