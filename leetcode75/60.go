package leetcode75

func tribonacci(n int) int {
	t0 := 0
	t1 := 1
	t2 := 1
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	for i := 3; i <= n; i++ {
		t2, t1, t0 = t0+t1+t2, t2, t1
	}
	return t2
}