package leetcode75

func removeStars(s string) string {
	var retArray []rune
	for _, value := range s {
		if string(value) != "*" {
			retArray = append(retArray, value)
		} else {
			retArray = retArray[:len(retArray)-1]
		}
	}
	return string(retArray)
}