package main

import (
	"fmt"
)

// Закрытый ключ
// для каждого сервиса генерируется закрытый ключ p q — натуральные числа
//  генерируется открытый ключ (НОД(p,q), НОК(p,q))
func gcdAndLcm(gcd, lcm int) int {
	if lcm%gcd != 0 {
		return 0
	}
	mul := lcm / gcd
	keys := 0

	for i := 1; i*i <= mul; i++ {
		//fmt.Println(a)
		if mul%i == 0 {
			if GCDRemainder(i, mul/i) == 1 {
				if i*i != mul {
					keys++
				}
				keys++
			}
		}
	}
	return keys
}

// GCDRemainder calculates GCD iteratively using remainder.
func GCDRemainder(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(gcdAndLcm(x, y))
}
