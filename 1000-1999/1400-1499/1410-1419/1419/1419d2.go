// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
}

type Number interface {
	int64 | float64 | int
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
func abs[K Number](a K) K {
	if a < 0 {
		return -a
	}
	return a
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Ints(a)
	b := make([]int, n)
	i := n - 1
	j := n - 1

	for i >= 0 {
		b[i] = a[j]
		i -= 2
		j--
	}
	if n%2 == 0 {
		b[0] = a[j]
		j--
	}
	i = n - 2
	missed := make([]int, 0)
	ans := 0
	for j >= 0 {
		if i-1 >= 0 {
			if a[j] < b[i-1] && a[j] < b[i+1] {
				b[i] = a[j]
				j--
				i -= 2
				ans++
			} else {
				missed = append(missed, a[j])
				j--
			}
		}
	}
	j = 0
	for i := 0; i < n; i++ {
		if b[i] == 0 {
			b[i] = missed[j]
			j++
		}
	}
	write(f, ans, "\n")
	for i := 0; i < n; i++ {
		write(f, b[i], " ")
	}
	write(f, "\n")
}
