package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m, q int
	fmt.Scan(&n)

	blackList := make(map[string]bool)
	for i := 0; i < n; i++ {
		var bl string
		fmt.Scan(&bl)
		blackList[bl] = true
	}

	fmt.Scan(&m)
	paths := make(map[string][]string)
	for i := 0; i < m; i++ {
		var path string
		fmt.Scan(&path)
		slashIdx := strings.LastIndex(path, "/")

		pToFile := path[:slashIdx+1]
		file := path[slashIdx+1:]
		paths[pToFile] = append(paths[pToFile], file)
	}

	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var request string
		fmt.Scan(&request)

		reqMap := make(map[string]int)
		for k, v := range paths {
			if k[:len(request)] == request {
				for _, fl := range v {
					fileTypeIdx := strings.LastIndex(fl, ".")
					reqMap[fl[:fileTypeIdx]]++
				}
			}
		}

		fmt.Println(len(reqMap))
		for k, v := range reqMap {
			fmt.Printf("%s: %d", k, v)
		}
	}
}
