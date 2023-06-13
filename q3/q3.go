package q3

func ClimbStairs(n int) int {

	if n <= 2 {
		return n
	}

	prev1, prev2 := 1, 2
	ways := 0

	for i := 3; i <= n; i++ {
		ways = prev1 + prev2
		prev1, prev2 = prev2, ways

	}

	return ways

}
