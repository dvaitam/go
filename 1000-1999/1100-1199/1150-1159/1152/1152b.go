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
	var x int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &x)
	ans := make([]int, 0)
	curr := x
	ops := 0
	for true {
		if curr&(curr+1) == 0 {
			break
		}
		j := 0
		for i := 0; i <= 29; i++ {
			mask := 1 << i
			if mask > curr {
				break
			}
			if mask&curr == 0 {
				j = i
			}
		}
		ops++
		ans = append(ans, j+1)
		curr = curr ^ ((1 << (j + 1)) - 1)

		if curr&(curr+1) == 0 {
			break
		}
		ops++
		curr++
	}
	write(f, ops, "\n")
	for i := 0; i < len(ans); i++ {
		write(f, ans[i], " ")
	}
	write(f, "\n")
}
