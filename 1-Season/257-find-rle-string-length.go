package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func binarySearch(list [][]int64, num int64) int {
	if num <= list[0][0] {
		return 0
	}
	length := len(list)
	if num == list[length-1][0] {
		return length - 1
	}
	begin := 0
	end := length - 1
	for begin+1 < end {
		mid := (begin + end) / 2
		if list[mid][0] == num {
			return mid
		} else if list[mid][0] > num {
			end = mid
		} else {
			begin = mid
		}
	}
	return begin + 1
}

// Chad CBD
func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))
	list := make([][]int64, 0)
	ss := []rune(strings.TrimSpace(s))
	i := 0
	var sum int64 = 0
	var sumRle int64 = 0
	for i < len(ss) {
		beginI := i
		ch := ss[i]
		for !((ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')) {
			i++
			ch = ss[i]
		}
		newSumRle := int64(i - beginI + 1)
		var newSum int64 = 0
		if i-beginI <= 0 {
			newSum = 1
		} else {
			newSum, _ = strconv.ParseInt(string(ss[beginI:i]), 10, 64)
		}
		sum += newSum
		sumRle += newSumRle
		list = append(list, []int64{sum, sumRle})
		i++
	}

	sb := strings.Builder{}
	for j := 0; j < n; j++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Split(strings.TrimSpace(line), " ")
		left, _ := strconv.ParseInt(parts[0], 10, 64)
		right, _ := strconv.ParseInt(parts[1], 10, 64)
		if right-left <= 1 {
			sb.WriteString(fmt.Sprintf("%d\n", right-left+1))
			continue
		}
		indexLeft := binarySearch(list, left)
		indexRight := binarySearch(list, right)
		var res int64 = 0
		if indexLeft == indexRight {
			if right-left <= 1 {
				res = right - left + 1
			} else {
				d := right - left + 1
				res = int64(math.Ceil(math.Log10(float64(d)+0.1))) + 1
			}
		} else {
			d := list[indexLeft][0] - left + 1
			if d == 1 {
				res = 1
			} else {
				res = int64(math.Ceil(math.Log10(float64(d)+0.1))) + 1
			}
			d = right - list[indexRight-1][0]
			if d == 1 {
				res++
			} else {
				res += int64(math.Ceil(math.Log10(float64(d)+0.1))) + 1
			}
			if indexRight-indexLeft > 1 {
				res += list[indexRight-1][1] - list[indexLeft][1]
			}
		}
		sb.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(sb.String())
}
