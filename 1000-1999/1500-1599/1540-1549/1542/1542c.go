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
func gcd(a int64, b int64) int64 {
	if b > a {
		return gcd(b, a)
	}
	if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
}

func lcm(a int64, b int64) int64 {
	return a * b / gcd(a, b)
}
func clcm(n int64, cache map[int64]int64) int64 {
	if cache[n] > 0 {
		return cache[n]
	}
	if n == 1 {
		return 1
	} else {
		ans := lcm(n, clcm(n-1, cache))
		cache[n] = ans
		return ans
	}
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	mod := int64(1000000007)
	cache := map[int64]int64{}
	for t := 1; t <= T; t++ {
		var n int64
		fmt.Fscan(reader, &n)
		if n == 1 {
			write(f, "2\n")
			continue
		} else if n == 2 {
			write(f, "5\n")
			continue
		}
		ans := int64(0)
		curr := int64(2)
		for true {
			prev := clcm(curr-1, cache)
			total := n / prev
			if total == 0 {
				break
			}

			nxt := clcm(curr, cache)
			ans += (total - (n / nxt)) * curr
			ans = ans % mod
			curr++
		}
		write(f, ans, "\n")

	}
}
