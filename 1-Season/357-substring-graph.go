package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	edges := make(map[string]int)
	setNode := make(map[string]bool)

	var number int
	fmt.Scan(&number)

	var line string
	for i := 0; i < number; i++ {
		fmt.Scan(&line)

		for j := 0; j < utf8.RuneCountInString(line)-3; j++ {
			firstSubstring := line[j : j+3]
			secondSubstring := line[j+1 : j+1+3]

			substrs := firstSubstring + secondSubstring

			setNode[firstSubstring] = true
			setNode[secondSubstring] = true
			edges[substrs]++
		}
	}
	fmt.Println(len(setNode))
	fmt.Println(len(edges))

	for k, v := range edges {
		fmt.Println(k[0:3], k[3:6], v)
	}
}
