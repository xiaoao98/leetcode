package daily

import "strconv"

func countSeniors(details []string) int {
	ret := 0
	for _, detail := range details {
		age, _ := strconv.Atoi(detail[11:13])
		if age > 60 {
			ret ++
		}
	}
	return ret
}