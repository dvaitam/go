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
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	b := make([]int, m)
	x := 0
	junction := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		x = x ^ a[i]
	}
	junction = junction ^ a[0]
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
		x = x ^ b[i]
		if i != m-1 {
			junction = junction ^ b[i]
		}
	}
	if x == 0 {
		write(f, "YES\n")
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if i == 0 {
					if j != m-1 {
						write(f, b[j], " ")
					} else {
						write(f, junction, "\n")
					}
				} else {
					if j != m-1 {
						write(f, "0 ")
					} else {
						write(f, a[i], "\n")
					}
				}
			}
		}
	} else {
		write(f, "NO\n")
	}

}
