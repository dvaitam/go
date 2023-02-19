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
func length(m int, cache map[int]int) int {
	if cache[m] > 0 {
		return cache[m]
	}
	if m < 10 {
		return 1
	}
	ans := length(m-10, cache) + length(m-9, cache)
	ans = ans % 1000000007
	cache[m] = ans
	return ans
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	cache := map[int]int{}
	for t := 1; t <= T; t++ {
		var n string
		var m int
		fmt.Fscan(reader, &n, &m)
		ans := 0

		for i := 0; i < len(n); i++ {
			ans += length(m+int(n[i]-'0'), cache)
			ans = ans % 1000000007
		}
		write(f, ans, "\n")

	}
}
