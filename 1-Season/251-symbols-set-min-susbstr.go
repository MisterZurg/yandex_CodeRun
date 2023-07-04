//package main
//
//import (
//	"bufio"
//	"fmt"
//	"math"
//	"os"
//)
//
//func symbolsSetMinSusbstr(s, c string) int {
//	// Default corner case
//	if len(c) > len(s) {
//		return 0
//	}
//
//	// initialize frequency table
//	freq := make(map[rune]int)
//	for _, lt := range c {
//		freq[lt]++
//	}
//
//	// initialize sliding window
//	counter := len(freq)
//	size := math.MaxInt
//	begin := 0
//	end := 0
//	// aaaabbbbbcccdab
//	// |			 | end
//	// | start
//
//	// Traverse until we found window that has c
//	for end < len(s) {
//		// for _, lt := range s {
//		// if current char found in table, decrement count
//		endChar := rune(s[end])
//		if _, ok := freq[endChar]; ok {
//			freq[endChar]--
//			if freq[endChar] == 0 {
//				counter--
//			}
//		}
//		// if counter == 0, means we found an answer,
//		// now try to trim that window by sliding begin to right.
//		end++
//
//		// aaaabbbbbcccdab
//		// |			 | end
//		// | start
//		for counter == 0 {
//			if end-begin < size {
//				size = end - begin
//			}
//
//			startChar := rune(s[begin])
//			// drop char that not in substr
//			if _, ok := freq[startChar]; ok {
//				freq[startChar]++
//				if freq[startChar] > 0 {
//					counter++
//				}
//			}
//			begin++
//		}
//	}
//	if size == math.MaxInt {
//		return 0
//	}
//	return size
//}
//
//func ParseInput() (string, string) {
//	sc := bufio.NewScanner(os.Stdin)
//	sc.Scan()
//	s := sc.Text()
//
//	sc.Scan()
//	c := sc.Text()
//
//	return s, c
//}
//
//func main() {
//	s, c := ParseInput()
//	fmt.Println(symbolsSetMinSusbstr(s, c))
//	//fmt.Println(minWindow(s, c))
//}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 251. Набор символов - by ChatGPT
func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	c, _ := reader.ReadString('\n')

	ch := []rune(strings.TrimSpace(c))
	set := make(map[rune]bool)
	for _, cc := range ch {
		set[cc] = true
	}

	result := int(^uint(0) >> 1)
	left := 0
	right := 0

	for right <= len(s) {
		chars := make(map[rune]bool)
		for i := left; i < right; i++ {
			cur := rune(s[i])
			if !set[cur] {
				left = i + 1
				break
			}
			chars[cur] = true
		}

		if containsAll(set, chars) {
			if right-left < result {
				result = right - left
			}
			left++
		} else {
			right++
		}
	}

	if result == int(^uint(0)>>1) {
		result = 0
	}

	fmt.Println(result)
}

func containsAll(set map[rune]bool, chars map[rune]bool) bool {
	for char := range set {
		if !chars[char] {
			return false
		}
	}
	return true
}

/*
 a |----|
				b |----|



 a |----|
      b |----|


     a |----|
b |----|


a |----|
 b |----|

*/
