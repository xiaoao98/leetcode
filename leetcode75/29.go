package leetcode75

func predictPartyVictory(senate string) string {
	banned := make([]bool, len(senate))
	rigD, rigR := 0, 0
	for true {
		for key, value := range senate {
			if !banned[key] && value == 'D' {
				if rigR > 0 {
					rigR--
					banned[key] = true
				} else {
					rigD++
				}
			} else if !banned[key] && value == 'R' {
				if rigD > 0 {
					rigD--
					banned[key] = true
				} else {
					rigR++
				}
			}
		}
		if rigD > 0 {
			for key, value := range senate {
				if !banned[key] && value == 'R' {
					rigD--
					banned[key] = true
				}
				if rigD == 0 {
					break
				}
			}
			if rigD > 0 {
				return "Dire"
			}
		}
		if rigR > 0 {
			for key, value := range senate {
				if !banned[key] && value == 'D' {
					rigR--
					banned[key] = true
				}
				if rigR == 0 {
					break
				}
			}
			if rigR > 0 {
				return "Radiant"
			}
		}
	}
	return "Radiant"
}