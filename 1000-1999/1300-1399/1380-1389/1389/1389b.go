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

func get_max(a []int, i int, k int, z int, last_move int, cache [][][]int) int {
	if k == 0 {
		return 0
	}
	if cache[k][z][last_move] > 0 {
		return cache[k][z][last_move]
	}
	right := a[i+1] + get_max(a, i+1, k-1, z, 0, cache)
	if last_move == 1 || z == 0 || i == 0 {
		cache[k][z][last_move] = right
		return right
	}
	left := a[i-1] + get_max(a, i-1, k-1, z-1, 1, cache)
	cache[k][z][last_move] = max(left, right)
	return cache[k][z][last_move]
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, k, z int
		fmt.Fscan(reader, &n, &k, &z)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		cache := make([][][]int, n+1)

		for j := 0; j <= k; j++ {
			cache[j] = make([][]int, z+1)
			for l := 0; l <= z; l++ {
				cache[j][l] = make([]int, 2)
			}
		}

		write(f, a[0]+get_max(a, 0, k, z, 0, cache), "\n")
	}
}
