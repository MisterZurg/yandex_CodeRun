////package main
////
////import (
////	"bufio"
////	"fmt"
////	"os"
////	"strconv"
////	"strings"
////)
////
////func main() {
////	//var n int
////	//fmt.Scan(&n)
////	sc := bufio.NewScanner(os.Stdin)
////	sc.Scan()
////
////	number := sc.Text()
////	n, _ := strconv.Atoi(number)
////
////	sc.Scan()
////	words := sc.Text()
////	nums := strings.Split(words, " ")
////
////	freq := make(map[uint]uint)
////	var topFreq uint = 0
////
////	// var curr uint
////	//for i := 0; i < n; i++ {
////	for _, num := range nums {
////		curr, _ := strconv.Atoi(num)
////		// fmt.Scan(&curr)
////		freq[uint(curr)]++
////
////		if freq[uint(curr)] == freq[topFreq] && uint(curr) > topFreq || freq[uint(curr)] > freq[topFreq] {
////			topFreq = uint(curr)
////		}
////	}
////	fmt.Println(topFreq)
////}
//
//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func main() {
//	// s := []int{-1, -1, -2, -3, -3, -4}
//
//	sc := bufio.NewScanner(os.Stdin)
//	sc.Scan()
//
//	_ = sc.Text()
//	sc.Scan()
//
//	words := sc.Text()
//	nums := strings.Split(words, " ")
//
//	maxCount := 0
//	maxElement := 0
//	processed := make(map[int]bool)
//	for _, num := range nums {
//		count := 0
//		a, _ := strconv.Atoi(num)
//		if processed[a] {
//			continue
//		}
//		for _, nm := range nums {
//			b, _ := strconv.Atoi(nm)
//			if a == b {
//				count++
//			}
//		}
//		processed[a] = true
//		if count > maxCount {
//			maxCount = count
//			maxElement = a
//		}
//	}
//
//	fmt.Println(maxElement)
//}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//func main() {
//	sc := bufio.NewScanner(os.Stdin)
//	sc.Scan()
//
//	_ = sc.Text()
//	sc.Scan()
//
//	words := sc.Text()
//	strNums := strings.Split(words, " ")
//	numbers := make([]int, len(strNums))
//	for i := range strNums {
//		numbers[i], _ = strconv.Atoi(strNums[i])
//	}
//
//	sort.Ints(numbers)
//
//	var topFreq int = 0
//	freq := make(map[int]int)
//
//	for _, num := range numbers {
//		freq[num]++
//
//		if freq[num] >= freq[topFreq] {
//			topFreq = num
//		}
//	}
//	fmt.Println(topFreq)
//}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	//var n int
	//fmt.Scan(&n)
	//
	//freq := make(map[uint]uint)
	//var topFreq uint = 0
	//
	//var curr uint
	//for i := 0; i < n; i++ {
	//	fmt.Scan(&curr)
	//	freq[curr]++
	//
	//	if freq[curr] == freq[topFreq] {
	//		if curr > topFreq {
	//			topFreq = curr
	//		}
	//	} else if freq[curr] > freq[topFreq] {
	//		topFreq = curr
	//	}
	//}
	//fmt.Println(topFreq)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	_ = sc.Text()
	sc.Scan()
	toBytes := "[" + strings.ReplaceAll(sc.Text(), " ", ",") + "]"
	//fmt.Println(txt)
	var fuckingUINT []int
	if err := json.Unmarshal([]byte(toBytes), &fuckingUINT); err != nil {
		panic(err)
	}

	sort.Slice(fuckingUINT, func(i, j int) bool {
		return fuckingUINT[i] < fuckingUINT[j]
	})

	var topFreq int = 0
	freq := make(map[int]uint)
	for _, num := range fuckingUINT {
		freq[num]++
		if freq[num] >= freq[topFreq] {
			topFreq = num
		}
	}
	fmt.Println(topFreq)
}
