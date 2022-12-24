// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
}

type Number interface {
	int64 | float64 | int
}
type Pair struct {
	i, j int
}

func min[K Number](a K, b K) K {
	if a < b {
		return a
	}
	return b
}
func max[K Number](a K, b K) K {
	if a > b {
		return a
	}
	return b
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		b := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &b[i])
		}
		operations := make([]Pair, 0)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if a[i] > a[j] || b[i] > b[j] {
					a[i], a[j] = a[j], a[i]
					b[i], b[j] = b[j], b[i]
					operations = append(operations, Pair{i: i + 1, j: j + 1})
				}
			}
		}
		ok := true
		for i := 1; i < n; i++ {
			if b[i] < b[i-1] {
				ok = false
				break
			}
		}
		for i := 1; i < n; i++ {
			if a[i] < a[i-1] {
				ok = false
				break
			}
		}
		if ok {
			write(f, len(operations), "\n")
			for i := 0; i < len(operations); i++ {
				write(f, operations[i].i, operations[i].j, "\n")
			}
		} else {
			write(f, "-1\n")
		}
	}
}
