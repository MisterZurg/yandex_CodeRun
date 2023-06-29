package main

import (
	"fmt"
	"sort"
)

func main() {
	var seller, buyer int
	fmt.Scan(&seller, &buyer)

	sellPrices := make([]int, seller)
	for i := 0; i < seller; i++ {
		fmt.Scan(&sellPrices[i])
	}

	buyPrices := make([]int, buyer)
	for i := 0; i < buyer; i++ {
		fmt.Scan(&buyPrices[i])
	}

	// O(n log n)
	sort.Ints(sellPrices)
	sort.Sort(sort.Reverse(sort.IntSlice(buyPrices)))

	result := 0
	minRange := min(seller, buyer)

	for i := 0; i < minRange; i++ {
		profit := buyPrices[i] - sellPrices[i]
		if profit > 0 {
			result += profit
		}
	}
	fmt.Println(result)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
