package main

import (
	"fmt"
)

func exactlyOneOccur(n int) int {
	set := make(map[int]int)

	var num int
	res := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if _, ok := set[num]; !ok {
			res++
		}
		set[num]++

		if val := set[num]; 1 < val && val < 3 {
			res--
		}
	}
	return res
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(exactlyOneOccur(n))
}
