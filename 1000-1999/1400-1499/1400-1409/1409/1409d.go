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
func get_sum(n int64) int64 {
	ans := int64(0)
	for n > 0 {
		ans += n % 10
		n = n / 10
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
		var n, s int64
		fmt.Fscan(reader, &n, &s)
		ans := int64(0)
		curr := int64(10)
		for i := 0; i < 18; i++ {
			ns := get_sum(n)
			if ns <= s {
				break
			}
			ans += curr - n%curr
			n += curr - n%curr
			curr = curr * 10
		}
		write(f, ans, "\n")
	}
}
