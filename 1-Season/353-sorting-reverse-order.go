package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

// 353. Сортировка положительных чисел в обратном порядке (разминка)
func main() {
	// Строка — URL сервера.
	// Целое число — порт сервера.
	// Целое число — число a.
	// Целое число — число b.
	var url string
	var port, a, b int

	fmt.Scan(&url, &port, &a, &b)

	get_request := fmt.Sprintf("%s:%s?a=%s&b=%s",
		url, port, a, b)

	response, _ := http.Get(get_request)
	// Потому что на этой платформе Go 1.14.4
	data, _ := ioutil.ReadAll(response.Body)

	var numbers []int
	json.Unmarshal(data, &numbers)

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	for _, number := range numbers {
		if number > 0 {
			fmt.Println(number)
		}
	}
}
