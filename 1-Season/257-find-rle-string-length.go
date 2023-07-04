// https://www.cyberforum.ru/python-beginners/thread2487445.html
// https://github.com/kissChriss/tasks/blob/master/yandex/readme.md

// https://github.com/naemnamenmea/competitive-programming/blob/2665c6535bba047a88b6658db3b79d2a58e187f5/yandex-contest/yandex-contest%202019%20internship/code/RLE-compress.cpp
package main

import (
	"fmt"
	"strconv"
)

type Request struct {
	l int
	r int
}

type ff struct {
	lenBefore int64
	elHere    int64
	lenHere   int64
}

func InitializeMap(m map[int64]ff, s string) {
	flag := false
	num := ""
	cnt := int64(0)
	len := int64(0)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isSymbol(c) {
			lenPrev := len
			cntPrev := cnt
			if flag {
				flag = false
				numInt, _ := strconv.Atoi(num)
				cnt += int64(numInt)
			} else {
				cnt += 1
			}
			re := symbolsToRLE(cnt - cntPrev)
			len = lenPrev + re
			m[cntPrev] = ff{lenPrev, cnt - cntPrev, re}
			num = ""
		} else {
			flag = true
			num += string(c)
		}
	}
}

func main() {
	var s string // строка, состоящая из строчных букв латинского алфавита и цифр
	fmt.Scan(&s)

	var q int
	fmt.Scan(&q)
	requsts := make([]Request, q)
	for i := range requsts {
		fmt.Scan(&requsts[i].l, &requsts[i].r)
	}

	// InitMap()
}
