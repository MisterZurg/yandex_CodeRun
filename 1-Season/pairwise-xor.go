package main

import (
	"fmt"
	"math"
)

// Программист на пляже
//func pairwiseXor(a, b int) int {
//	similarity := a ^ b
//	return similarity
//}

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
		fmt.Println(pairwiseXor(seats))
		//var seat1, seat2, sim int
		//fmt.Scan(&seat1, &seat2)
		//sim = pairwiseXor(seat1, seat2)
		//for j := 2; j < places; j++ {
		//	seat1 = seat2
		//	fmt.Scan(&seat2)
		//	sim = min(sim, pairwiseXor(seat1, seat2))
		//}
		//fmt.Println(sim)
	}
}
