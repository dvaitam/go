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
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		ans := int64(0)
		sm := map[int64]int64{}
		//	last := map[int64]int{}
		for i := 0; i < n; i++ {
			if sm[a[i]] > 0 {
				ans += sm[a[i]] * int64(n-i)
				sm[a[i]] += int64(i + 1)
			} else {
				sm[a[i]] = int64(i + 1)
			}
		}
		write(f, ans, "\n")
	}
}
