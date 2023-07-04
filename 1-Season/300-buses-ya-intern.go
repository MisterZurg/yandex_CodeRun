package main

import (
	"fmt"
	"math"
)

//http://mathhelpplanet.com/viewtopic.php?f=36&t=49238

// F_ksi(x)
func cumulativeDistributionFunction(x int) int {
	return int(1 - math.Pow(float64(1-(x/2)), float64(2)))
}

func mathExpectation(x) int {
	// integral (1 - x/2)dx
	return 0
}

func main() {
	var n int
	fmt.Scan(&n)

	routes := make([]int, n)
	for i := range routes {
		fmt.Scan(&routes[i])
	}

}
