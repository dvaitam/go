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
func upper_half(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 1 + n/2
	}
}
func get_worst(n int) int {
	if n == 0 {
		return 0
	}
	half1, half2 := n/2, n/2
	if n%2 == 1 {
		half2++
	}
	return max(upper_half(half1)+upper_half(half2), upper_half(half1+1)+upper_half(half2-1))

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
		b := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &b[i])
		}
		known, unknown := 0, 0
		ans := 0
		for i := 0; i < n; i++ {
			if b[i] == 1 {
				unknown++
			} else {
				known += unknown
				unknown = 0
			}
			ans = max(ans, get_worst(known)+unknown)
			//write(f, known, unknown, ans, "\n")
		}
		ans = max(ans, get_worst(known)+unknown)

		write(f, ans, "\n")
	}
}
