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

var table [][]bool
var upd [][]int
var updl []int

func setv(a int, b int) {
	if !table[a][b] {
		table[a][b] = true
		if a > 0 {
			setv(a-1, b)
		}
		upd[a] = append(upd[a], b)
	}
}
func roll(a int) {
	for len(upd[a]) > 0 {
		ll := len(upd[a])
		setv(a-1, upd[a][ll-1]^a)
		ll = len(upd[a])
		upd[a] = upd[a][:ll-1]
	}
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
	table = make([][]bool, 5001)
	upd = make([][]int, 5001)
	updl = make([]int, 8192)

	for i := 0; i < 5001; i++ {
		table[i] = make([]bool, 8192)
		upd[i] = make([]int, 0, 8192)
	}
	for i := 0; i <= 5000; i++ {
		setv(i, 0)
	}
	for i := n - 1; i >= 0; i-- {
		if a[i] > 0 {
			roll(a[i])
		}
	}
	ans := make([]int, 0)
	for i := 0; i < 8192; i++ {
		if table[0][i] {
			ans = append(ans, i)
		}
	}
	write(f, len(ans), "\n")
	for i := 0; i < len(ans); i++ {
		write(f, ans[i], " ")
	}
	write(f, "\n")

}
