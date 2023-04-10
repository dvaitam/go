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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var a, b int64
		fmt.Fscan(reader, &a, &b)
		if gcd(a, b) == 1 {
			write(f, 1, "\n")
			write(f, a, b, "\n")
		} else {
			write(f, 2, "\n")

			if gcd(a+1, b) == 1 && a != 1000000000 {
				write(f, a+1, b, "\n")
				write(f, a, b, "\n")
			} else if gcd(a, b+1) == 1 && b != 1000000000 {
				write(f, a, b+1, "\n")
				write(f, a, b, "\n")
			} else if gcd(a-1, b) == 1 {
				write(f, a-1, b, "\n")
				write(f, a, b, "\n")
			} else {
				write(f, a, b-1, "\n")
				write(f, a, b, "\n")
			}

		}
	}
}
