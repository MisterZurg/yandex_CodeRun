// // https://www.cyberforum.ru/cpp-beginners/thread1701784.html
//
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DoubleHeap []float64

// Functions required for heap interface
// Len returns the length of the heap
func (h DoubleHeap) Len() int {
	return len(h)
}

// Less reports whether the element with index i should sort before the element with index j
func (h DoubleHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the elements with indexes i and j
func (h DoubleHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds an element to the heap
func (h *DoubleHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

// Pop removes and returns the smallest element from the heap
func (h *DoubleHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func solveHeap(arr []int, n int) int64 {
	g := make(DoubleHeap, 0)
	s := make(DoubleHeap, 0)

	heap.Init(&g)
	heap.Init(&s)

	var result int64

	for i := 0; i < n; i++ {
		heap.Push(&s, -1.0*float64(arr[i]))
		heap.Push(&g, -1.0*s[0])
		heap.Pop(&s)

		if g.Len() > s.Len() {
			heap.Push(&s, -1.0*g[0])
			heap.Push(&g, -1.0*g[0])
			heap.Pop(&g)
		}

		result += int64(-1.0 * s[0])
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	partsStr, _ := reader.ReadString('\n')
	parts := strings.Fields(partsStr)
	arr := make([]int, n)
	for i, part := range parts {
		arr[i], _ = strconv.Atoi(part)
	}

	result := solveHeap(arr, n)
	fmt.Println(result)
}

/*
import java.io.*;
import java.time.Duration;
import java.time.LocalDateTime;
import java.util.*;

import java.io.*;
import java.util.*;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));

        int n = Integer.parseInt(reader.readLine());
        String[] parts = reader.readLine().split(" ");
        int[] arr = new int[n];
        for (int i = 0; i < n; i++) {
            arr[i] = Integer.parseInt(parts[i]);
        }
        long result = solveHeap(arr, n);
        System.out.println(result);
    }

    public static long solveHeap(int arr[], int n) {
        PriorityQueue<Double> g = new PriorityQueue<>();
        PriorityQueue<Double> s = new PriorityQueue<>();
        long result = 0;
        for (int i = 0; i < n; i++) {
            s.add(-1.0 * arr[i]);
            g.add(-1.0 * s.poll());
            if (g.size() > s.size()) {
                s.add(-1.0 * g.poll());
            }
            result += (-1.0 * s.peek());
        }
        return result;
    }
}
*/
