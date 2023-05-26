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
func pow(x int64, n int) int64 {
	if n == 0 {
		return 1
	} else {
		return (x * pow(x, n-1)) % 1000000007
	}
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i+1])
	}
	ans := int64(1)
	visited := map[int]bool{}
	all_counted := 0
	for i := 1; i <= n; i++ {
		if visited[i] {
			continue
		}
		sub_visited := map[int]bool{}
		sub_visited[i] = true
		curr := a[i]
		for !sub_visited[curr] && !visited[curr] {
			sub_visited[curr] = true
			curr = a[curr]
		}
		if visited[curr] {
			for g := range sub_visited {
				visited[g] = true
			}
			continue
		}
		nxt_curr := a[curr]
		count := 1
		for nxt_curr != curr {
			nxt_curr = a[nxt_curr]
			count++
		}
		all_counted += count
		ans = ans * (pow(2, count) - 2 + 1000000007)

		ans = ans % 1000000007
		for g := range sub_visited {
			visited[g] = true
		}
		sub_visited = nil

	}
	ans = ans * pow(2, n-all_counted)
	ans = ans % 1000000007

	write(f, ans, "\n")

}
