package main

import (
	"bufio"
	"fmt"
	"os"
)

// №352 Камни и украшения (разминка)
func rocksAndJewels(rocks, jewels string) int {
	cnt := 0
	vvs := make(map[rune]bool)
	for _, j := range jewels {
		vvs[j] = true
	}

	for _, rock := range rocks {
		if _, ok := vvs[rock]; ok {
			cnt++
		}
	}
	return cnt
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	jewels := sc.Text()
	sc.Scan()
	rocks := sc.Text()

	fmt.Println(rocksAndJewels(rocks, jewels))
}
