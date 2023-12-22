package leetcode75

func asteroidCollision(asteroids []int) []int {
	var stack []int
	i := 0
	for i < len(asteroids) {
		value := asteroids[i]
		if value > 0 || len(stack) == 0 || stack[len(stack)-1] < 0 {
			stack = append(stack, value)
			i++
		} else if stack[len(stack)-1] < -value {
			stack = stack[:len(stack)-1]
		} else if stack[len(stack)-1] > -value {
			i++
		} else if stack[len(stack)-1] == -value {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return stack
}