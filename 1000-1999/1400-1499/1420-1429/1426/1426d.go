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
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sm := int64(0)
	index := map[int64]int{}
	ans := 0
	index[0] = -1
	rnd := int64(4398493894343)
	for i := 0; i < n; i++ {
		sm += a[i]
		if index[sm] != 0 {
			ans++
			sm += rnd
			index[sm-a[i]] = i + 1
		}
		index[sm] = i + 1
		//write(f, i, ans, " y\n")
	}
	if sm == 0 && ans == 0 {
		ans++
	}
	write(f, ans, "\n")
}
