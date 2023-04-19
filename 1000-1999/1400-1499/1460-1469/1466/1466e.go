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
	mod := int64(1000000007)
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		x := make([]int64, n)
		mx := int64(0)
		total := int64(0)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &x[i])
			mx = max(mx, x[i])
			total += x[i]
			total %= mod
		}
		curr := int64(1)
		sm := make([]int64, 64)
		j := 0
		for j < 63 {
			for i := 0; i < n; i++ {
				if x[i]&curr == curr {
					sm[j] += (curr % mod)
					sm[j] %= mod
				}
			}
			curr = curr << 1
			j++
		}
		//write(f, sm, "\n")
		ans := int64(0)

		for i := 0; i < n; i++ {
			or_sum := int64(0)
			and_sum := int64(0)
			curr := int64(1)
			j := 0
			for j < 63 {
				if x[i]&curr == curr {
					or_sum += (curr % mod) * int64(n)
					or_sum %= mod
					and_sum += sm[j]
					and_sum %= mod
				} else {
					or_sum += sm[j]
					or_sum %= mod
				}
				curr = curr << 1
				j++
			}
			//write(f, or_sum, and_sum, "\n")
			ans += (or_sum * and_sum) % mod
			ans %= mod
		}

		write(f, ans, "\n")

	}
}
