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
		a := make([]int, n)
		m := map[int]int{}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])

			m[a[i]]++
		}
		ans := make([]byte, n)
		for i := 0; i < n; i++ {
			ans[i] = '0'
		}
		if len(m) == n {
			ans[0] = '1'
		}
		if m[1] > 0 {
			ans[n-1] = '1'
		}
		i, j := 0, n-1
		curr := 1
		index := n - 2
		for index > 0 {
			if a[i] == curr && m[curr] == 1 && m[curr+1] > 0 {
				i++
				ans[index] = '1'
				index--
				curr++
			} else if a[j] == curr && m[curr] == 1 && m[curr+1] > 0 {
				j--
				ans[index] = '1'
				index--
				curr++
			} else {
				break
			}
		}
		write(f, string(ans), "\n")

	}
}
