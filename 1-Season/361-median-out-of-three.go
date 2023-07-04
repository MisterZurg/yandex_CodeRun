package main

import (
	"fmt"
	"sort"
)

// Средний элемент
func medianOutOfThree(nums []int) int {
	sort.Ints(nums)
	return nums[1]
}

func main() {
	nums := make([]int, 3)
	for i := range nums {
		fmt.Scan(&nums[i])
	}

	fmt.Println(medianOutOfThree(nums))
}
