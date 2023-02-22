package climbingstairs

/*
	70. Climbing Stairs

	You are climbing a staircase. It takes n steps to reach the top.
	Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/

// Basically, it's a n+1 Fibonacci's number
func climbStairs(n int) int {
	a, b := 0, 1

	for i := 1; i < n; i++ {
		a, b = b, a+b
	}

	return a + b
}
