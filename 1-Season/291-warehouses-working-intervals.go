package main

import (
	"bufio"
	"os"
)

type Note struct {
	WarhouseID string
}

func Solution(in []string) {
	replaceLater := make(map[string]string)

}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	in := []string{}
	for sc.Scan() {
		in = append(in, sc.Text())
	}

	Solution(in)
}
