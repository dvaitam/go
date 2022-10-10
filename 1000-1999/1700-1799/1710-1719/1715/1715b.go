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
		var n, k, b, s int64
		fmt.Fscan(reader, &n, &k, &b, &s)

		if s < b*k || s > b*k+(k-1)*n {
			write(f, "-1\n")
		} else {
			rem := s - b*k
			start := b * k
			if rem > (n-1)*(k-1) {
				start += rem - (n-1)*(k-1)
				rem = (n - 1) * (k - 1)
			}
			write(f, start)

			for i := 1; i < int(n); i++ {
				if rem >= k-1 {
					write(f, " ", k-1)
					rem -= k - 1
				} else if rem > 0 {
					write(f, " ", rem)
					rem = 0
				} else {
					write(f, " ", 0)
				}

			}
			write(f, "\n")
		}

	}
}
