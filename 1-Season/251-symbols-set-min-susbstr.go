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
