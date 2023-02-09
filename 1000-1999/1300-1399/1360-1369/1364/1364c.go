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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	b := make([]int, n)
	curr := 0
	assigned := map[int]bool{}
	unassigned := make([]int, 0)
	zero := false
	mx := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		mx = max(mx, a[i])
	}
	if mx > n {
		write(f, "-1\n")
	} else {
		for i := 0; i < n; i++ {
			if assigned[a[i]] {
				unassigned = append(unassigned, i)
			} else {
				if a[i] == 1 && !zero {
					b[i] = 0
					curr = 1
				} else if a[i] == 0 {
					zero = true
					unassigned = append(unassigned, i)
				} else {
					if zero {
						b[i] = 0
						curr = 1
						zero = false
					} else {
						b[i] = curr
						curr++
					}
					for curr < a[i] && len(unassigned) > 0 {
						len := len(unassigned)
						last := unassigned[len-1]
						b[last] = curr
						curr++
						unassigned = unassigned[:len-1]
					}

				}
				assigned[a[i]] = true
			}
		}
		curr++
		for i := 0; i < len(unassigned); i++ {
			b[unassigned[i]] = curr
		}
		for i := 0; i < n; i++ {
			write(f, b[i])
			if i == n-1 {
				write(f, "\n")
			} else {
				write(f, " ")
			}
		}

	}

}
