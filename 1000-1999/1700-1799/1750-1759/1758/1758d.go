// 00
package main

import (
	"bufio"
	"fmt"
	"math"
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
		var n int64
		fmt.Fscan(reader, &n)
		for i := int64(1); i < 100; i++ {
			first_n_sum := n*i + ((n)*(n-1))/2
			sq := math.Sqrt(float64(first_n_sum))
			mx := int64(sq) + 1
			if i+n-1 <= mx {
				ans := make([]int64, 0)
				for j := int64(0); j < n-1; j++ {
					ans = append(ans, i+j)
				}
				ans = append(ans, i+mx)
				excess := mx*mx - first_n_sum
				start := n - 2
				for excess > 0 && start >= 0 {
					ans[start] += 1
					excess--
				}
				for j := int64(0); j < n-1; j++ {
					write(f, ans[j], " ")
				}
				write(f, ans[n-1], "\n")
				break

			}

		}

	}
}
