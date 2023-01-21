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
func main() {
	var n int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	if n == 1 {
		write(f, "1\n")
	} else if n == 2 {
		write(f, "2\n")
	} else {
		start := n * (n - 1)
		ans := int64(1)
		for i := n; i > 0; i-- {
			ans = max(ans, start*i/gcd(start, i))
		}
		for i := n; i > 2; i-- {
			ans = max(ans, (i*(i-1)*(i-2))/gcd(i*(i-1), i-2))
			//write(f, i, (i*(i-1)*(i-2))/gcd(i*(i-1), i-2), "\n")
		}
		//write(f, gcd(923, gcd(922, 921)), "\n")

		write(f, ans, "\n")
	}

}
