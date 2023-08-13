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
	var k int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &k)
	a := make([]int64, 10)
	for i := 0; i < 10; i++ {
		a[i] = 1
	}
	i := 0
	curr := int64(1)
	for curr < k {
		a[i]++
		i++
		i = i % 10
		tmp := int64(1)
		for j := 0; j < 10; j++ {
			tmp = tmp * a[j]
		}
		curr = tmp
	}
	s := "codeforces"
	for i := 0; i < 10; i++ {
		for j := int64(0); j < a[i]; j++ {
			write(f, s[i:i+1])
		}
	}
	write(f, "\n")
}
