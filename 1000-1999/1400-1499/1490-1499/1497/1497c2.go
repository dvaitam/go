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
func solve3(n int) []int {
	if n%3 == 0 {
		return []int{n / 3, n / 3, n / 3}
	} else {
		if n%2 == 1 {
			return []int{(n - 1) / 2, (n - 1) / 2, 1}
		} else {
			if n%4 == 0 {
				return []int{n / 2, n / 4, n / 4}

			} else {
				return []int{2, (n - 2) / 2, (n - 2) / 2}
			}
		}
	}
}
func solve(n int, k int, factor int) []int {
	if n/k == 1 {
		ans := make([]int, k)
		curr := n % k
		for i := 0; i < k; i++ {
			if curr > 0 {
				ans[i] = 2 * factor
				curr--
			} else {
				ans[i] = 1 * factor
			}
		}
		return ans
	} else if k == 3 {
		ans := solve3(n)
		for i := 0; i < k; i++ {
			ans[i] = ans[i] * factor
		}
		return ans
	} else if n%k == 0 {
		ans := make([]int, k)
		for i := 0; i < k; i++ {
			ans[i] = (n / k) * factor
		}
		return ans
	} else {
		if n%2 == 1 {
			part := solve(n-1, k-1, factor)
			part = append(part, factor)
			return part
		} else {
			return solve(n/2, k, factor*2)
		}
	}
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
		ans := solve(n, k, 1)
		for i := 0; i < k; i++ {
			write(f, ans[i], " ")
		}
		write(f, "\n")
		// total := n
		// for i := k; i > 0; i-- {
		// 	curr := ff
		// 	for (i-1)*curr < total {
		// 		curr = curr / 2
		// 	}
		// 	total -= curr
		// 	write(f, curr, " ")
		// }
	}
}
