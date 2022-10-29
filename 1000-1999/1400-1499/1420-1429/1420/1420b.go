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
func count_bits(a int) int {
	ans := 1
	start := 1
	for start*2 <= a {
		start = start << 1
		ans++
	}
	return ans
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
		a := make([]int, n)
		m := map[int]int{}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			m[count_bits(a[i])]++
		}
		ans := int64(0)
		for _, v := range m {
			ans += int64(v) * int64(v-1) / 2
		}
		write(f, ans, "\n")

	}
}
