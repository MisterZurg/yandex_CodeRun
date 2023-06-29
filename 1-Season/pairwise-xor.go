package main

import (
	"fmt"
	"math"
	"sort"
)

// Программист на пляже
// TODO : 8TC - Time limit exceeded
func pairwiseXor(seats []int) int {
	similarity := math.MaxInt
	for i := 1; i < len(seats); i++ {
		similarity = min(similarity, seats[i]^seats[i-1])
	}
	return similarity
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

	var places int
	for i := 0; i < n; i++ {
		fmt.Scan(&places)
		seats := make([]int, places)
		for st := range seats {
			fmt.Scan(&seats[st])
		}
		sort.Ints(seats)

		fmt.Println(pairwiseXor(seats))
	}
}
