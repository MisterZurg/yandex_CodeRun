package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// 306. День недели
func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		raw_date := sc.Text()
		t, _ := time.Parse("2 January 2006", raw_date)
		fmt.Println(t.Weekday())
	}
}
