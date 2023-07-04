package main

import (
	"fmt"
	"math"
	"sort"
)

// Программист на пляже
// TODO : 8TC - Time limit exceeded
func pairwiseXor(seats []int) int {
	similarity := math.MaxInt
	for i := 1; i < len(seats); i++ {
		similarity = min(similarity, seats[i]^seats[i-1])
	}
	return similarity
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)

	var places int
	for i := 0; i < n; i++ {
		fmt.Scan(&places)
		seats := make([]int, places)
		for st := range seats {
			fmt.Scan(&seats[st])
		}
		sort.Ints(seats)

		fmt.Println(pairwiseXor(seats))
	}
}

/* C++
#include <algorithm>
#include <cstdint>
#include <cstdio>

using namespace std;

int main() {
    int t;
    scanf("%d", &t);
    while (t--) {
        int n;
        scanf("%d", &n);
        int* v = (int*) malloc(sizeof(int) * n);
        for (int i = 0; i < n; ++i) {
            scanf("%d", &v[i]);
        }
        sort(v, v + n);
        int ans = INT32_MAX;
        for (int i = 1; i < n; ++i) {
            ans = min(ans, v[i] ^ v[i - 1]);
        }
        printf("%d\n", ans);
    }
    return 0;
}
*/
