package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	productsToCategories := make(map[int]int)
	for i := 0; i < n; i++ {
		var prod, cat int
		fmt.Scan(&prod, &cat)
		productsToCategories[prod] = cat
	}

	products := make([]int, n)
	for i := range products {
		fmt.Scan(&products[i])
	}

	i, j := 0, 1
	currentUnique := make(map[int]bool)
	currentUnique[products[i]] = true

	diversity := n
	for j < n {
		if i < n && j < n &&
			!currentUnique[productsToCategories[products[j]]] {
			currentUnique[productsToCategories[products[j]]] = true
			j++
		} else {
			diversity = min(diversity, j-i)
			for i < j {
				currentUnique[productsToCategories[products[i]]] = false
				i++
			}
		}
		// fmt.Println(currentUnique)
	}
	fmt.Println(diversity)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//5
//1 1
//2 1
//3 1
//4 2
//5 2
//1 4 2 5 3
