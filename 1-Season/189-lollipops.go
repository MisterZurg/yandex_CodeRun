package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const INF = 1000000000 + 21

var scanner *bufio.Scanner

type Node struct {
	x     int
	value int
	y     int
	l     *Node
	r     *Node
}

func newNode(x, value int) *Node {
	return &Node{
		x:     x,
		value: value,
		y:     rand.Intn(1000000000) + 1,
	}
}

func merge(t1, t2 *Node) *Node {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	if t1.y < t2.y {
		t1.r = merge(t1.r, t2)
		return t1
	} else {
		t2.l = merge(t1, t2.l)
		return t2
	}
}

func split(t *Node, x int) (*Node, *Node) {
	if t == nil {
		return nil, nil
	}

	if t.x < x {
		t1, t2 := split(t.r, x)
		t.r = t1
		return t, t2
	} else {
		t1, t2 := split(t.l, x)
		t.l = t2
		return t1, t
	}
}

func add(t, nd *Node) *Node {
	if t == nil {
		return nd
	}

	if nd.y < t.y {
		nd.l, nd.r = split(t, nd.x)
		return nd
	}

	root := t

	var p *Node
	lastL := false
	for t != nil && t.y < nd.y {
		p = t
		if t.x < nd.x {
			lastL = false
			t = t.r
		} else {
			lastL = true
			t = t.l
		}
	}

	if lastL {
		p.l = nd
	} else {
		p.r = nd
	}

	nd.l, nd.r = split(t, nd.x)

	return root
}

func remove(t *Node, x int) *Node {
	if t.x == x {
		return merge(t.l, t.r)
	}

	root := t

	var p *Node
	lastL := false
	for t != nil && t.x != x {
		p = t
		if t.x < x {
			lastL = false
			t = t.r
		} else {
			lastL = true
			t = t.l
		}
	}

	if lastL {
		p.l = merge(t.l, t.r)
	} else {
		p.r = merge(t.l, t.r)
	}

	return root
}

func getUp(t *Node, x int) *Node {
	var cur *Node
	for t != nil {
		if t.x >= x && (cur == nil || cur.x > t.x) {
			cur = t
		}

		if t.x >= x {
			t = t.l
		} else {
			t = t.r
		}
	}

	return cur
}

func getDown(t *Node, x int) *Node {
	var cur *Node
	for t != nil {
		if t.x <= x && (cur == nil || cur.x < t.x) {
			cur = t
		}

		if t.x >= x {
			t = t.l
		} else {
			t = t.r
		}
	}

	return cur
}

// Converted from Official Yandex Blog
func main() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	n := nextInt()
	_ = nextInt()
	q := nextInt()

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}

	tree := newNode(0, INF)
	tree = add(tree, newNode(n+1, INF))

	for i := 0; i < n; i++ {
		if i == n-1 || arr[i] != arr[i+1] {
			tree = add(tree, newNode(i+1, arr[i]))
		}
	}

	ans := make([]int, q)
	for i := 0; i < q; i++ {
		a := nextInt()
		b := nextInt()
		l := nextInt()
		r := nextInt()

		ptr1 := getUp(tree, l)
		ptr2 := getUp(tree, r)

		if ptr1 == ptr2 && ptr1.value == a {
			pr := getDown(tree, l-1)
			if pr.x+1 != l {
				tree = add(tree, newNode(l-1, a))
			}

			if pr.x+1 == l && pr.value == b {
				tree = remove(tree, pr.x)
			}

			needAdd := true
			if r == ptr1.x {
				nx := getUp(tree, r+1)
				if nx.value == b {
					needAdd = false
				}
				tree = remove(tree, r)
			}

			if needAdd {
				tree = add(tree, newNode(r, b))
			}

			// ans[i] = "1"
			ans[i] = 1
		}
	}
	// strings.ReplaceAll(strings.Join(ans, "\n"), "", "0")
	for _, v := range ans {
		fmt.Println(v)
	}

}

func nextInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}
