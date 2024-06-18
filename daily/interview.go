package daily

import "strconv"

//input- etddgeia, output- vowels:aeiconsonants:d2g1t1
func TransferString(input string) string {
	vowels := make(map[byte]bool)
	vowels['a'] = true
	vowels['e'] = true
	vowels['i'] = true
	vowels['o'] = true
	vowels['u'] = true
	retVowels := make(map[byte]bool)
	consonants := make(map[byte]int)
	for i :=0; i < len(input); i++ {
		c := input[i]
		if vowels[c] {
			retVowels[c] = true
		} else {
			if _, exists := consonants[c]; exists {
				consonants[c] += 1
			} else {
				consonants[c] = 1
			}
		}
	}
	retVowelsChar := make([]byte, 0)
	retConsonantsString := ""
	for key, _ := range retVowels {
		retVowelsChar = append(retVowelsChar, key)
	}
	for key, value := range consonants {
		retConsonantsString += string(key)
		retConsonantsString += strconv.Itoa(value)
	}
	retString := "vowel:" + string(retVowelsChar) + "consonants:" + retConsonantsString
	return retString
}


