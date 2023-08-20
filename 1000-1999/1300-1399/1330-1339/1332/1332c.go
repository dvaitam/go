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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, k int
		var s string
		fmt.Fscan(reader, &n, &k, &s)
		ans := 0
		for i := 0; i < k/2; i++ {
			c := make([]int, 26)
			for j := i; j < n; j += k {
				c[s[j]-'a']++
			}
			for j := k - 1 - i; j < n; j += k {
				c[s[j]-'a']++
			}
			mx := 0
			sm := 0
			for j := 0; j < 26; j++ {
				mx = max(mx, c[j])
				sm += c[j]
			}
			ans += sm - mx
		}
		if k%2 == 1 {
			c := make([]int, 26)
			for j := k / 2; j < n; j += k {
				c[s[j]-'a']++
			}
			mx := 0
			sm := 0
			for j := 0; j < 26; j++ {
				mx = max(mx, c[j])
				sm += c[j]
			}
			ans += sm - mx
		}
		write(f, ans, "\n")
	}
}
