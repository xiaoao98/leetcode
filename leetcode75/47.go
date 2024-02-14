package leetcode75

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	equationGraph := make(map[string]map[string]float64, 0)
	equationSet := make(map[string]bool, 0)
	for i, equation := range equations {
		former := equation[0]
		latter := equation[1]
		if _, ok := equationGraph[former]; !ok {
			equationGraph[former] = make(map[string]float64)
		}
		equationGraph[former][latter] = values[i]
		if _, ok := equationGraph[latter]; !ok {
			equationGraph[latter] = make(map[string]float64)
		}
		equationGraph[latter][former] = 1 / values[i]
		equationSet[former] = true
		equationSet[latter] = true
	}
	rets := make([]float64, 0)
	for _, query := range queries {
		former := query[0]
		latter := query[1]
		if !equationSet[former] || !equationSet[latter] {
			rets = append(rets, -1)
			continue
		}
		if former == latter {
			rets = append(rets, 1)
			continue
		}
		visited := make(map[string]bool, 0)
		ret := dfs47(former, latter, 1, visited, equationGraph)
		rets = append(rets, ret)
	}
	return rets
}

func dfs47(former string, latter string, accumulated float64, visited map[string]bool, equationGraph map[string]map[string]float64) float64 {
	visited[former] = true
	var ret float64
	ret = -1
	if ans, exists := equationGraph[former][latter]; exists {
		ret = accumulated * ans
		return ret
	}
	for key, value := range equationGraph[former] {
		if !visited[key] {
			ret = dfs47(key, latter, accumulated*value, visited, equationGraph)
			if ret != -1 {
				return ret
			}
		}
	}
	visited[former] = false
	return ret
}
