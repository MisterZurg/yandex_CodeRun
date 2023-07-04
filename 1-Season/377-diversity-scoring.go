package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 377. Оценка разнообразия - by ChatGPT
func main() {
	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))
	allProduct := make(map[int64]int64, n)
	for i := 0; i < n; i++ {
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		productAndcategory := strings.Split(s, " ")
		productStr, categoryStr := productAndcategory[0], productAndcategory[1]
		product, _ := strconv.ParseInt(productStr, 10, 64)
		category, _ := strconv.ParseInt(categoryStr, 10, 64)
		allProduct[product] = category
	}

	mapCategories := make(map[int64][]int, n)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	products := strings.Split(s, " ")
	for i, str := range products {
		product, _ := strconv.ParseInt(str, 10, 64)
		category := allProduct[product]
		arr := mapCategories[category]
		arr = append(arr, i+1)
		mapCategories[category] = arr
	}

	result := int(9223372036854775807)
	count := int64(0)

	for _, arr := range mapCategories {
		if len(arr) == 1 {
			count++
		} else {
			minimal := arr[1] - arr[0]
			for i := 2; i < len(arr); i++ {
				div := arr[i] - arr[i-1]
				if div < minimal {
					minimal = div
				}
			}
			if minimal < result {
				result = minimal
			}
		}
	}

	if count == int64(len(mapCategories)) {
		result = int(count)
	}

	fmt.Println(result)
}
