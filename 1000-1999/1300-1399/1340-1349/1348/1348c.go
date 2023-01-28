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
	int64 | float64 | int | string
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
		var n, k int
		var s string
		fmt.Fscan(reader, &n, &k, &s)
		sb := []byte(s)
		sort.Slice(sb, func(i, j int) bool { return sb[i] < sb[j] })
		if sb[k-1] != sb[0] || k == n {
			write(f, string(sb[k-1:k]), "\n")
		} else {
			if sb[k] != sb[n-1] {
				write(f, string(sb[k-1:]), "\n")
			} else {
				ans := make([]byte, 0)
				ans = append(ans, sb[k-1])
				l := (n - k) / k
				if (n-k)%k != 0 {
					l++
				}
				for i := 0; i < l; i++ {
					ans = append(ans, sb[k])
				}
				write(f, string(ans), "\n")
			}
		}
	}
}
