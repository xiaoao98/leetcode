package leetcode75

func minCostClimbingStairs(cost []int) int {
	t0 := cost[0]
	t1 := cost[1]
	// t2 := MinInts(t0, t1) + cost[2]
	for i := 2; i < len(cost); i++ {
		t0, t1 = t1, MinInts(t0, t1)+cost[i]
	}
	return MinInts(t0, t1)
}