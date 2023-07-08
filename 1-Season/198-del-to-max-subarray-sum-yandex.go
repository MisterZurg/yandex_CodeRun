package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF   = 1000000000 + 21
	LLINF = (1 << 60) + 5
	MOD   = 1000000000 + 7
	MAX_N = 100000 + 227
	MAX_K = 105
)

var (
	n, k   int
	arr    [MAX_N]int
	dp_l   [MAX_N][MAX_K]int64
	dp_r   [MAX_N][MAX_K]int64
	reader *bufio.Reader
	writer *bufio.Writer
)

func delToMaxSubarraySumYandex() {
	fmt.Fscan(reader, &n, &k)

	ans := int64(-LLINF)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
		ans = max(ans, int64(arr[i]))
	}

	if ans <= 0 {
		fmt.Fprintln(writer, ans)
		return
	}

	for i := 0; i <= k; i++ {
		dp_l[n+1][i] = 0
		dp_r[n+1][i] = 0
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			dp_l[i][j] = max64(0, dp_l[i-1][j]+int64(arr[i-1]))
			if j > 0 {
				dp_l[i][j] = max64(dp_l[i][j], dp_l[i-1][j-1])
			}

			ans = max(ans, dp_l[i][j])
		}
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			dp_l[i][j] = dp_l[i-1][j] + int64(arr[i-1])
			if j > 0 {
				dp_l[i][j] = max64(dp_l[i][j], dp_l[i-1][j-1])
			}
		}
	}

	for i := n; i >= 1; i-- {
		for j := 0; j <= k; j++ {
			dp_r[i][j] = dp_r[i+1][j] + int64(arr[i-1])
			if j > 0 {
				dp_r[i][j] = max64(dp_r[i][j], dp_r[i+1][j-1])
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			if i != 1 {
				dp_l[i][j] = max64(dp_l[i][j], dp_l[i-1][j])
			}
			if j > 0 {
				dp_l[i][j] = max64(dp_l[i][j], dp_l[i][j-1])
			}
		}
	}

	for i := n; i >= 1; i-- {
		for j := 0; j <= k; j++ {
			if i != n {
				dp_r[i][j] = max64(dp_r[i][j], dp_r[i+1][j])
			}
			if j > 0 {
				dp_r[i][j] = max64(dp_r[i][j], dp_r[i][j-1])
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= k; j++ {
			ans = max(ans, dp_l[i][j]+dp_r[i+1][k-j])
		}
	}

	fmt.Fprintln(writer, ans)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)
	for t > 0 {
		delToMaxSubarraySumYandex()
		t--
	}
}
