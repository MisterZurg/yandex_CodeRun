package main

import "fmt"

// 227. Библиотека
func library(k, m, d int) int {
	cnt := 1
	firstWeekPass := true

	for d != 1 {
		if d < 6 {
			m += k
		}

		if m < cnt {
			firstWeekPass = false
			break
		}

		m -= cnt
		cnt++

		d = d%7 + 1
	}

	if firstWeekPass {
		for m+5*k >= 7*cnt+21 {
			m = m + 5*k - (7*cnt + 21)
			cnt += 7
		}
		for {
			if d < 6 {
				m += k
			}
			if m < cnt {
				break
			}

			m -= cnt
			cnt++

			d = d%7 + 1
		}
	}
	return cnt - 1
}

func main() {
	var k, m, d int
	fmt.Scan(&k, &m, &d)

	fmt.Println(library(k, m, d))
}

/* C++
#include <iostream>

void library() {
    long k, m, d;
    std::cin >> k >> m >> d;
    long count = 1;
    bool first_week_pass = true;
    while (d != 1) {
        if (d < 6) m += k;
        if (m < count) {
            first_week_pass = false;
            break;
        }
        m -= count++;
        d = d % 7 + 1;
    }
    if (first_week_pass) {
        while (m + 5 * k >= 7 * count + 21) {
            m = m + 5 * k - (7 * count + 21);
            count += 7;
        }
        while (true) {
            if (d < 6) m += k;
            if (m < count) {
                break;
            }
            m -= count++;
            d = d % 7 + 1;
        }
    }
    std::cout << count - 1;
}

int main() {
    library();
    return 0;
}
*/
