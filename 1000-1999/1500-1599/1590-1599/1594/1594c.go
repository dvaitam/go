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
		var c, s string
		fmt.Fscan(reader, &n, &c, &s)
		equal := true
		for i := 0; i < n; i++ {
			if s[i] != c[0] {
				equal = false
				break
			}
		}
		// if T == 10000 {
		// 	if t == 71 {
		// 		fmt.Println(n, s, c)
		// 	}
		// } else {
		if equal {
			write(f, "0\n")
		} else {
			one := -1
			for i := n / 2; i <= n; i++ {
				if i*2 > n && s[i-1] == c[0] {
					one = i
					break
				}
			}
			if one == -1 {
				write(f, 2, "\n")
				write(f, n-1, n, "\n")
			} else {
				write(f, 1, "\n")
				write(f, one, "\n")
			}
			//	}
		}

	}
}
