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
		a := make([]int64, n)
		b := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			b[i] = a[i]
		}
		// if t == 1252 {
		// 	fmt.Println(a)
		// } else if T > 1000 {
		// 	continue
		// }
		for i := 1; i < n; i++ {
			a[i-1], a[i] = 0, a[i]-a[i-1]
			//	write(f, a, "\n")
		}
		if a[n-1] >= 0 {
			write(f, "YES\n")
		} else {
			for i := n - 1; i > 0; i-- {
				a[i], a[i-1] = 0, a[i-1]-a[i]
			}
			if a[0] <= 0 {
				write(f, "YES\n")
			} else {
				write(f, "NO\n")
			}
		}
	}
}
