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
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var a, b, q int64
		fmt.Fscan(reader, &a, &b, &q)
		lm := lcm(a, b)
		sm := max(a, b)
		for q > 0 {
			var l, r int64
			fmt.Fscan(reader, &l, &r)
			ans := int64(0)
			start := l - (l % lm)
			if l%lm != 0 {
				start += lm
			}
			if start-lm+sm >= l {
				if start-lm+sm <= r {
					ans += start - lm + sm - l
				} else {
					ans += r - l + 1
				}
			}
			end := (r / lm) * lm
			if end >= l {
				ans += ((end - start) / lm) * sm
				if end+sm <= r {
					ans += sm
				} else {
					ans += (r - end + 1)
				}
			}
			write(f, r-l+1-ans, " ")
			q--
		}

		write(f, "\n")
	}
}
