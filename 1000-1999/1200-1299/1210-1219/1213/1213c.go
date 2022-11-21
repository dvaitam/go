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
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, m int64
		fmt.Fscan(reader, &n, &m)
		ans := int64(0)
		total := n / m
		rep := total / 10
		rem := total % 10
		tens := int64(0)
		for i := 1; i <= 10; i++ {
			if int64(i) <= rem {
				ans += (m * int64(i)) % 10
			}
			tens += (m * int64(i)) % 10
		}
		ans += tens * rep
		write(f, ans, "\n")

	}
}
