package main

import "fmt"

type Cheker struct {
	n int
	m int
}

func main() {
	var n, m, w, b int // размеры доски, количество белых шашек, blak, на поле
	fmt.Scan(&n, &m, &w)

	whites := make(map[Cheker]bool)
	for i := 0; i < w; i++ {
		var chek Cheker
		fmt.Scan(&chek.n, &chek.m)
		whites[chek] = true
	}

	fmt.Scan(&b)

	blacks := make(map[Cheker]bool)

	for i := 0; i < b; i++ {
		var chek Cheker
		fmt.Scan(&chek.n, &chek.m)
		blacks[chek] = true
	}

	var first string
	fmt.Scan(&first)

	switch first {
	case "white":
		fmt.Println(canEat(whites, blacks, n, m))
	case "black":
		fmt.Println(canEat(blacks, whites, n, m))
	}
}

// Brute?
func canEat(first, second map[Cheker]bool, n, m int) string {
	for i := range first {
		for j := range second {
			// We have to check if:
			// - still on GAYming board
			// - no other figures
			// top right
			if i.n+2 <= n && i.m+2 <= m {
				if i.n+1 == j.n && i.m+1 == j.m && !first[Cheker{n: i.n + 2, m: i.m + 2}] && !second[Cheker{n: i.n + 2, m: i.m + 2}] {
					return "Yes"
				}
			}
			// top left
			if i.n+2 <= n && i.m-2 > 0 {
				if i.n+1 == j.n && i.m-1 == j.m && !first[Cheker{n: i.n + 2, m: i.m - 2}] && !second[Cheker{n: i.n + 2, m: i.m - 2}] {
					return "Yes"
				}
			}
			// bot right
			if i.n-2 > 0 && i.m+2 <= m {
				if i.n-1 == j.n && i.m+1 == j.m && !first[Cheker{n: i.n - 2, m: i.m + 2}] && !second[Cheker{n: i.n - 2, m: i.m + 2}] {
					return "Yes"
				}
			}

			// bot left
			if i.n-2 > 0 && i.m-2 > 0 {
				if i.n-1 == j.n && i.m-1 == j.m && !first[Cheker{n: i.n - 2, m: i.m - 2}] && !second[Cheker{n: i.n - 2, m: i.m - 2}] {
					return "Yes"
				}
			}
		}
	}
	return "No"
}
