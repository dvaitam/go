// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		has_one := false
		has_zero := false
		other := false
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if a[i] == 0 {
				has_zero = true
			}
			if a[i] == 1 {
				has_one = true
			} else {
				other = true
			}
		}
		if has_one && !other {
			write(f, "YES\n")
		} else {
			if has_one && has_zero {
				write(f, "NO\n")
			} else {
				if !has_one {
					write(f, "YES\n")
				} else {
					sort.Ints(a)
					ok := true
					for i := 1; i < n; i++ {
						if a[i] != a[i-1] && a[i] == a[i-1]+1 {
							ok = false
							break
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
	}
}
