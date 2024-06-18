package daily

func NumSteps(s string) int {
	// decimalN := 0
	// for i := 0; i < len(s); i++ {
	// 	decimalN = decimalN*2 + int(s[i]-'0')
	// }
	// ret := 0
	// for decimalN != 1 {
	// 	if decimalN%2 == 1 {
	// 		decimalN += 1
	// 	} else {
	// 		decimalN /= 2
	// 	}
	// 	ret++
	// }
	// return ret
	ret := 0
	for s != "1" {
		if s[len(s)-1] == '0' {
			s = s[:len(s)-1]
		} else {
			s = AddOne(s)
		}
		ret += 1
	}
	return ret
}

func AddOne(s string) string {
	carry := 1
	ret := make([]byte, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' && carry == 1 {
			ret[i] = '1'
			carry = 0
		} else if s[i] == '1' && carry == 1 {
			ret[i] = '0'
		} else {
			ret[i] = s[i]
		}
	}
	if carry == 1 {
		return "1" + string(ret)
	}
	return string(ret)
}