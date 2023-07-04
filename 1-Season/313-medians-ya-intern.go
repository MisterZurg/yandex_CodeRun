//package main
//
//import (
//	"fmt"
//)
//
////https://www.cyberforum.ru/cpp-beginners/thread1701784.html
//
//func main() {
//	var n int
//	fmt.Scan(&n)
//
//	L := make([]int, n+1)
//	for i := 1; i < len(L); i++ {
//		fmt.Scan(&L[i])
//	}
//
//	result := 0
//	// result += L[0]
//
//	for i := 2; i < len(L); i++ {
//		var median int
//		sub_slice := L[:i]
//		// sort.Sort(sort.Reverse(sort.IntSlice(sub_slice)))
//		if i%2 == 0 {
//			median = sub_slice[i/2]
//		} else {
//			median = sub_slice[(i+1)/2]
//		}
//		fmt.Println(sub_slice, median)
//		result += median
//	}
//	fmt.Println(result)
//}
package main

import (
	"fmt"
	"sort"
)

func median(sequence []int) float64 {
	n := len(sequence)
	sort.Ints(sequence)
	if n%2 == 0 {
		return float64(sequence[n/2-1]+sequence[n/2]) / 2.0
	} else {
		return float64(sequence[n/2])
	}
}

func main() {
	var n int
	fmt.Print("Введите количество элементов в последовательности: ")
	fmt.Scanln(&n)

	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Введите элемент #%d: ", i+1)
		fmt.Scanln(&sequence[i])

		// Находим медиану для первых i элементов
		medianValue := median(sequence[:i+1])
		fmt.Printf("Медиана для первых %d элементов: %.2f\n", i+1, medianValue)
	}

	// Вычисляем сумму найденных значений
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += median(sequence[:i+1])
	}
	fmt.Printf("Сумма найденных медиан: %.2f\n", sum)
}
