package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var mapChar map[rune]int
var mapInteger map[int]rune

// Chad CBD

func main() {
	fillMap()
	reader := bufio.NewReader(os.Stdin)
	textArray, _ := reader.ReadString('\n')
	textArray = strings.TrimSpace(textArray)
	textArraySlice := strings.Split(textArray, " ")
	nStr, _ := reader.ReadString('\n')
	n := int64(0)
	fmt.Sscanf(nStr, "%d", &n)

	message := make([]string, 0)
	for i := int64(0); i < n; i++ {
		word, _ := reader.ReadString('\n')
		message = append(message, strings.TrimSpace(word))
	}
	//reader.Close()

	decode(textArraySlice, message)
}

func decode(textArray []string, message []string) {
	match := make(map[string]string)
	for _, word := range textArray {
		match[convertWord(word)] = word
	}

	sb := strings.Builder{}
	for _, word := range message {
		tmp := convertWord(word)
		sb.WriteString(match[tmp])
		sb.WriteString("\n")
	}

	fmt.Println(sb.String())
}

func convertWord(word string) string {
	charArray := []rune(word)
	k := 1 - mapChar[charArray[0]]
	if k == 0 {
		return word
	}

	sb := strings.Builder{}
	for _, ch := range charArray {
		n := mapChar[ch] + k
		if n <= 0 {
			n += 26
		}
		sb.WriteRune(mapInteger[n])
	}

	return sb.String()
}

func fillMap() {
	mapChar = make(map[rune]int)
	mapInteger = make(map[int]rune)

	i := 1
	for c := 'a'; c <= 'z'; c++ {
		mapChar[c] = i
		mapInteger[i] = c
		i++
	}
}
