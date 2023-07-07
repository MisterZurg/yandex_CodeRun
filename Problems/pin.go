package main

import (
	"fmt"
	"sort"
)

// Гвоздики
func pin(n int, carnations []int) int {
	sort.Ints(carnations)

	dp := make([]int, n+1)
	dp[2] = carnations[2] - carnations[1]

	if n > 2 {
		dp[3] = carnations[3] - carnations[1]

		for i := 4; i <= n; i++ {
			dp[i] = min(dp[i-2], dp[i-1]) + carnations[i] - carnations[i-1]
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)
	carnations := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&carnations[i])
	}

	fmt.Println(pin(n, carnations))
}
