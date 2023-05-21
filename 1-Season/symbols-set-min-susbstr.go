package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func symbolsSetMinSusbstr(s, c string) int {
	// Default corner case
	if len(c) > len(s) {
		return 0
	}

	// initialize frequency table
	freq := make(map[rune]int)
	for _, lt := range c {
		freq[lt]++
	}

	// initialize sliding window
	counter := len(freq)
	size := math.MaxInt
	begin := 0
	end := 0
	// aaaabbbbbcccdab
	// |			 | end
	// | start

	// Traverse until we found window that has c
	for end < len(s) {
		// for _, lt := range s {
		// if current char found in table, decrement count
		endChar := rune(s[end])
		if _, ok := freq[endChar]; ok {
			freq[endChar]--
			if freq[endChar] == 0 {
				counter--
			}
		}
		// if counter == 0, means we found an answer,
		// now try to trim that window by sliding begin to right.
		end++

		// aaaabbbbbcccdab
		// |			 | end
		// | start
		for counter == 0 {
			if end-begin < size {
				size = end - begin
			}

			startChar := rune(s[begin])
			// drop char that not in substr
			if _, ok := freq[startChar]; ok {
				freq[startChar]++
				if freq[startChar] > 0 {
					counter++
				}
			}
			begin++
		}
	}
	if size == math.MaxInt {
		return 0
	}
	return size
}

func ParseInput() (string, string) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()

	sc.Scan()
	c := sc.Text()

	return s, c
}

func main() {
	s, c := ParseInput()
	fmt.Println(symbolsSetMinSusbstr(s, c))
	//fmt.Println(minWindow(s, c))
}
