package main

import (
	"fmt"
	"strings"
)

type GoVnoFile struct {
	root string // "/"
	name string
	ext  string
}

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
	// paths := make(map[string][]string)
	goVnoFiles := []GoVnoFile{}
	for i := 0; i < m; i++ {
		var path string
		fmt.Scan(&path)

		inBlackList := false
		for nWord := range blackList {
			if strings.Contains(path, nWord) {
				inBlackList = true
				break
			}
		}

		if !inBlackList {
			continue
		}
		extension := path[strings.LastIndex(path, "."):]
		goVnoFiles = append(goVnoFiles, GoVnoFile{
			ext:  extension,
			name: path,
		})
	}

	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var request string
		fmt.Scan(&request)

		isFound := false
		reqMap := make(map[string]int)

		for _, gvf := range goVnoFiles {
			if len(gvf.name) < len(request) {
				continue
			}

			if strings.Compare(gvf.name[0:len(request)], request) == 0 {
				ext := gvf.ext
				//if _, ok := reqMap[ext]; !ok {
				//	reqMap[ext] = 0
				//}
				reqMap[ext]++
				isFound = true
			}
		}
		var sb strings.Builder
		if !isFound {
			sb.WriteString(fmt.Sprintf("0\n"))
		} else {
			sb.WriteString(fmt.Sprintf("%d\n", len(reqMap)))
			for k, v := range reqMap {
				sb.WriteString(fmt.Sprintf("%s: %d\n", k, v))
			}
		}
		fmt.Print(sb.String())
	}
}
