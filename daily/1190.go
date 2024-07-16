package daily

func reverseParentheses(s string) string { //can be solved by using wormhole
	size := len(s)
	output := make([]byte, 0)
	for i := 0; i < size; i++ {
		if s[i] != ')' {
			output = append(output, s[i])
		} else {
			toAdd := make([]byte, 0)
			for output[len(output)-1] != '(' {
				toAdd = append(toAdd, output[len(output)-1])
				output = output[:len(output)-1]
			}
			output = output[:len(output)-1]
			output = append(output, toAdd...)
		}
	}
	return string(output)
}