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
		fmt.Fscan(reader, &n, &k)
		a := make([]int, n)
		rems := make([]int, 0)
		ans := int64(0)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			ans += int64(a[i] / k)
			rem := a[i] % k
			rems = append(rems, rem)
		}
		sort.Ints(rems)
		i, j := 0, len(rems)-1
		for i < j {
			for i < len(rems) && rems[i]+rems[j] < k {
				i++
			}
			if i < len(rems) && i < j {
				ans++
				i++
				j--
			}
		}

		write(f, ans, "\n")

	}
}
