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
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		if n <= 3 {
			write(f, "YES\n")
		} else {
			ok := true
			inc := true

			if a[0] != 1 && a[0] != n {
				if a[1] < a[0] {
					inc = false
				}
			} else {
				if a[0] == 1 && a[1] == n || a[0] == n && a[1] != 1 {
					inc = false
				}
			}
			curr := a[0]
			for i := 0; i < n; i++ {
				if a[i] != curr {
					ok = false
					break
				}
				if inc {
					curr++
					if curr == n+1 {
						curr = 1
					}
				} else {
					curr--
					if curr == 0 {
						curr = n
					}
				}
			}
			if ok {
				write(f, "YES\n")
			} else {
				write(f, "NO\n")
			}
		}
	}
}
