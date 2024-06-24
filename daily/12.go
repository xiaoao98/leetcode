package daily

type ValueSymbolPair struct {
	value   int
	symbol1 byte
	symbol2 byte
}

func intToRoman(num int) string {
	valueSymbols := make([]ValueSymbolPair, 0)
	valueSymbols = append(valueSymbols, ValueSymbolPair{value: 1000, symbol1: 'M'})
	valueSymbols = append(valueSymbols, ValueSymbolPair{value: 100, symbol1: 'C', symbol2: 'D'})
	valueSymbols = append(valueSymbols, ValueSymbolPair{value: 10, symbol1: 'X', symbol2: 'L'})
	valueSymbols = append(valueSymbols, ValueSymbolPair{value: 1, symbol1: 'I', symbol2: 'V'})
	ret := make([]byte, 0)
	for i, pair := range valueSymbols {
		numS := num / pair.value
		if pair.symbol1 != 'M' && numS == 9 {
			ret = append(ret, pair.symbol1)
			ret = append(ret, valueSymbols[i-1].symbol1)
		} else if pair.symbol1 != 'M' && numS == 4 {
			ret = append(ret, pair.symbol1)
			ret = append(ret, pair.symbol2)
		} else {
			if numS >= 5 {
				ret = append(ret, pair.symbol2)
				numS -= 5
			}
			for i := 0; i < numS; i++ {
				ret = append(ret, pair.symbol1)
			}
		}
		num = num % pair.value
	}
	return string(ret)
}